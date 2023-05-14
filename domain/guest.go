package domain

import (
	"time"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/id64"
	"github.com/vburenin/nsync"

	"street/conf"
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file guest.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type guest.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type guest.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type guest.go
//go:generate farify doublequote --file guest.go

type (
	GuestRegisterIn struct {
		RequestCommon
		Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		Password string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}
	GuestRegisterOut struct {
		ResponseCommon
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		verifyEmailUrl string
	}
)

const (
	GuestRegisterAction = `guest/register`

	ErrGuestRegisterEmailInvalid       = `email must be valid`
	ErrGuestRegisterPasswordTooShort   = `password must be at least 12 characters`
	ErrGuestRegisterEmailUsed          = `email already used`
	ErrGuestRegisterUserCreationFailed = `user creation failed`

	minPassLength = 12
)

func (d *Domain) GuestRegister(in *GuestRegisterIn) (out GuestRegisterOut) {
	in.Email = S.Trim(S.ValidateEmail(in.Email))
	if in.Email == `` {
		out.SetError(400, ErrGuestRegisterEmailInvalid)
		return
	}
	if len(in.Password) < minPassLength {
		out.SetErrorf(400, ErrGuestRegisterPasswordTooShort)
		return
	}
	exists := rqAuth.NewUsers(d.AuthOltp)
	exists.Email = in.Email
	if exists.FindByEmail() {
		out.SetError(400, ErrGuestRegisterEmailUsed)
		return
	}
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = in.Email
	user.SetEncryptedPassword(in.Password, in.UnixNow())
	user.CreatedAt = in.UnixNow()
	user.SecretCode = S.RandomCB63(1)
	if !user.DoInsert() {
		out.SetError(500, ErrGuestRegisterUserCreationFailed)
		return
	}
	user.CensorFields()
	out.User = user.Users

	// send verification link
	hash := S.EncodeCB63(user.Id, 8)
	out.verifyEmailUrl = in.Host + `/` + GuestVerifyEmailAction + `?code=` + user.SecretCode + `&hash=` + hash

	d.runSubtask(func() {
		err := d.Mailer.SendRegistrationEmail(user.Email, out.verifyEmailUrl)
		L.IsError(err, `SendRegistrationEmail`)
		// TODO: insert failed event to clickhouse
	})
	return
}

type (
	GuestVerifyEmailIn struct {
		RequestCommon
		SecretCode string `json:"secretCode" form:"secretCode" query:"secretCode" long:"secretCode" msg:"secretCode"`
		Hash       string `json:"hash" form:"hash" query:"hash" long:"hash" msg:"hash"`
	}
	GuestVerifyEmailOut struct {
		ResponseCommon
		Ok bool `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
	}
)

const (
	GuestVerifyEmailAction = `guest/verifyEmail`

	ErrGuestVerifyEmailInvalidHash        = `invalid hash`
	ErrGuestVerifyEmailUserNotFound       = `user not found`
	ErrGuestVerifyEmailSecretCodeMismatch = `secret code mismatch`
	ErrGuestVerifyEmailModificationFailed = `failed modifying user`
)

func (d *Domain) GuestVerifyEmail(in *GuestVerifyEmailIn) (out GuestVerifyEmailOut) {
	userId, ok := S.DecodeCB63[uint64](in.Hash)
	if !ok {
		out.SetError(400, ErrGuestVerifyEmailInvalidHash)
		return
	}
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = userId
	if !user.FindById() {
		out.SetError(400, ErrGuestVerifyEmailUserNotFound)
		return
	}
	if user.VerifiedAt != 0 { // already verified
		out.Ok = true
		return
	}
	if user.SecretCode != in.SecretCode {
		out.SetError(400, ErrGuestVerifyEmailSecretCodeMismatch)
		return
	}
	user.SetVerifiedAt(in.UnixNow())
	user.SetSecretCode(``)
	user.SetSecretCodeAt(0)

	if !user.DoUpdateById() {
		out.SetError(500, ErrGuestVerifyEmailModificationFailed)
		return
	}
	out.Ok = true
	return
}

type (
	GuestLoginIn struct {
		RequestCommon
		Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		Password string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}
	GuestLoginOut struct {
		ResponseCommon
		User rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`
	}
)

const (
	GuestLoginAction = `guest/login`

	ErrGuestLoginEmailInvalid             = `email must be valid`
	ErrGuestLoginEmailOrPasswordIncorrect = `incorrect email or password`
	ErrGuestLoginPasswordOrEmailIncorrect = `incorrect password or email`
	ErrFailedStoringSession               = `failed storing session`
)

func (d *Domain) GuestLogin(in *GuestLoginIn) (out GuestLoginOut) {
	in.Email = S.Trim(S.ValidateEmail(in.Email))
	if in.Email == `` {
		out.SetError(400, ErrGuestLoginEmailInvalid)
		return
	}
	user := rqAuth.NewUsers(d.AuthOltp)
	user.Email = in.Email
	if !user.FindByEmail() {
		out.SetError(400, ErrGuestLoginEmailOrPasswordIncorrect)
		return
	}
	if err := S.CheckPassword(user.Password, in.Password); err != nil {
		out.SetError(400, ErrGuestLoginPasswordOrEmailIncorrect)
		return
	}
	user.CensorFields()
	out.User = *user
	session := d.createSession(user.Id, user.Email, in.UserAgent)
	// TODO: set list of roles in the session
	if !session.DoInsert() {
		out.SetError(500, ErrFailedStoringSession)
		return
	}
	out.SessionToken = session.SessionToken
	return
}

type (
	GuestForgotPasswordIn struct {
		RequestCommon
		Email string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	}

	GuestForgotPasswordOut struct {
		ResponseCommon
		Ok bool `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`

		resetPassUrl string
	}
)

const (
	GuestForgotPasswordAction = `guest/forgotPassword`

	ErrGuestForgotPasswordEmailNotFound         = `forgot password email not found`
	ErrGuestForgotPassworTriggeredTooFrequently = `forgot password triggered to frequently`
	ErrGuestForgotPasswordModificationFailed    = `forgot password modification failed`
)

var guestForgotPasswordLock = nsync.NewNamedMutex()

func (d *Domain) GuestForgotPassword(in *GuestForgotPasswordIn) (out GuestForgotPasswordOut) {
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = in.Email

	guestForgotPasswordLock.Lock(in.Email)

	if !user.FindByEmail() {
		guestForgotPasswordLock.Unlock(in.Email)
		out.SetError(400, ErrGuestForgotPasswordEmailNotFound)
		return
	}

	recently := in.TimeNow().Add(-conf.ForgotPasswordThrottleMinute * time.Minute).Unix()
	if user.SecretCodeAt >= recently {
		guestForgotPasswordLock.Unlock(in.Email)
		out.SetError(400, ErrGuestForgotPassworTriggeredTooFrequently)
		return
	}

	secretCode := id64.SID() + S.RandomCB63[int64](1)
	user.SetSecretCode(secretCode)
	user.SetSecretCodeAt(in.UnixNow())
	hash := S.EncodeCB63(user.Id, 8)

	out.resetPassUrl = in.Host + `/` + GuestResetPasswordAction + `?secretCode=` + secretCode + `&hash=` + hash
	d.runSubtask(func() {
		defer guestForgotPasswordLock.Unlock(in.Email)
		err := d.Mailer.SendResetPasswordEmail(user.Email, out.resetPassUrl)
		L.IsError(err, `SendResetPasswordEmail`)
		// TODO: insert failed event to clickhouse
	})

	if !user.DoUpdateById() {
		out.SetError(500, ErrGuestForgotPasswordModificationFailed)
		return
	}

	out.Ok = true
	return
}

type (
	GuestResetPasswordIn struct {
		RequestCommon
		SecretCode string `json:"secretCode" form:"secretCode" query:"secretCode" long:"secretCode" msg:"secretCode"`
		Hash       string `json:"hash" form:"hash" query:"hash" long:"hash" msg:"hash"`
		Password   string `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	}
	GuestResetPasswordOut struct {
		ResponseCommon
		Ok bool `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
	}
)

const (
	GuestResetPasswordAction = `guest/resetPassword`

	ErrGuestResetPasswordInvalidHash        = `invalid hash`
	ErrGuestResetPasswordTooShort           = `password too short`
	ErrGuestResetPasswordUserNotFound       = `user not found`
	ErrGuestResetPasswordWrongSecret        = `wrong secret code`
	ErrGuestResetPasswordExpiredLink        = `expired link`
	ErrGuestResetPasswordModificationFailed = `reset password modification failed`
)

func (d *Domain) GuestResetPassword(in *GuestResetPasswordIn) (out GuestResetPasswordOut) {
	userId, ok := S.DecodeCB63[uint64](in.Hash)
	if !ok {
		out.SetError(400, ErrGuestResetPasswordInvalidHash)
		return
	}
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = userId
	if !user.FindById() {
		out.SetError(400, ErrGuestResetPasswordUserNotFound)
		return
	}
	if len(in.Password) < minPassLength {
		out.SetErrorf(400, ErrGuestResetPasswordTooShort)
		return
	}
	if user.SecretCode == `` ||
		user.SecretCodeAt == 0 ||
		in.UnixNow()-user.SecretCodeAt > conf.ForgotPasswordExpireMinute*60 {
		out.SetError(404, ErrGuestResetPasswordExpiredLink)
		return
	}
	if user.SecretCode != in.SecretCode {
		out.SetError(400, ErrGuestResetPasswordWrongSecret)
		return
	}
	// also verify the user if not verified yet
	if user.VerifiedAt == 0 {
		user.SetVerifiedAt(in.UnixNow())
	}
	user.SetSecretCode(``)
	user.SetSecretCodeAt(0)
	user.SetEncryptedPassword(in.Password, in.UnixNow())

	if !user.DoUpdateById() {
		out.SetError(500, ErrGuestResetPasswordModificationFailed)
		return
	}
	out.Ok = true
	return
}
