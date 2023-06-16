package domain

import (
	"time"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/id64"
	"github.com/kokizzu/lexid"
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
	GuestDebugIn struct {
		RequestCommon
	}
	GuestDebugOut struct {
		ResponseCommon
		Request RequestCommon `json:"request" form:"request" query:"request" long:"request" msg:"request"`
	}
)

const (
	GuestDebugAction = `guest/debug`
)

func (d *Domain) GuestDebug(in *GuestDebugIn) (out GuestDebugOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.Request = in.RequestCommon
	return
}

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
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
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
	user.SecretCode = id64.SID() + S.RandomCB63(1)
	user.SetGenUniqueUsernameByEmail(in.Email, in.UnixNow())
	if !user.DoInsert() {
		out.SetError(500, ErrGuestRegisterUserCreationFailed)
		return
	}
	out.actor = user.Id

	// send verification link
	hash := S.EncodeCB63(user.Id, 8)
	out.verifyEmailUrl = in.Host + `/` + GuestVerifyEmailAction + `?secretCode=` + user.SecretCode + `&hash=` + hash

	user.CensorFields()
	out.User = user.Users

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
		Ok    bool   `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
		Email string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
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
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
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
	out.actor = userId

	out.Email = user.Email

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
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments M.SB `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
	}
)

const (
	GuestLoginAction = `guest/login`

	ErrGuestLoginEmailInvalid             = `email must be valid`
	ErrGuestLoginUserDeactivated          = `user deactivated`
	ErrGuestLoginEmailOrPasswordIncorrect = `incorrect email or password`
	ErrGuestLoginPasswordOrEmailIncorrect = `incorrect password or email`
	ErrGuestLoginFailedStoringSession     = `failed storing session for login`
)

func (d *Domain) GuestLogin(in *GuestLoginIn) (out GuestLoginOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
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
	out.actor = user.Id

	if user.DeletedAt > 0 {
		out.SetError(400, ErrGuestLoginUserDeactivated)
		return
	}

	if err := user.CheckPassword(in.Password); err != nil {
		out.SetError(400, ErrGuestLoginPasswordOrEmailIncorrect)
		return
	}
	user.CensorFields()
	out.User = user
	session, sess := d.createSession(user.Id, user.Email, in.UserAgent)

	// TODO: set list of roles in the session
	if !session.DoInsert() {
		out.SetError(500, ErrGuestLoginFailedStoringSession)
		return
	}
	out.SessionToken = session.SessionToken
	out.Segments = d.segmentsFromSession(sess)
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

	ErrGuestForgotPasswordEmailNotFound          = `forgot password email not found`
	ErrGuestForgotPasswordTriggeredTooFrequently = `forgot password triggered to frequently`
	ErrGuestForgotPasswordModificationFailed     = `forgot password modification failed`
)

var guestForgotPasswordLock = nsync.NewNamedMutex()

func (d *Domain) GuestForgotPassword(in *GuestForgotPasswordIn) (out GuestForgotPasswordOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = in.Email

	if !guestForgotPasswordLock.TryLock(in.Email) {
		out.SetError(400, ErrGuestForgotPasswordTriggeredTooFrequently)
		return
	}

	if !user.FindByEmail() {
		guestForgotPasswordLock.Unlock(in.Email)
		out.SetError(400, ErrGuestForgotPasswordEmailNotFound)
		return
	}
	out.actor = user.Id

	recently := in.TimeNow().Add(-conf.ForgotPasswordThrottleMinute * time.Minute).Unix()
	if user.SecretCodeAt >= recently {
		guestForgotPasswordLock.Unlock(in.Email)
		out.SetError(400, ErrGuestForgotPasswordTriggeredTooFrequently)
		return
	}

	secretCode := id64.SID() + S.RandomCB63(1)
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
		guestForgotPasswordLock.Unlock(in.Email)
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
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
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
	out.actor = user.Id
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

type (
	GuestResendVerificationEmailIn struct {
		RequestCommon
		Email string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	}
	GuestResendVerificationEmailOut struct {
		ResponseCommon
		Ok bool `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`

		verifyEmailUrl string
	}
)

const (
	GuestResendVerificationEmailAction = `guest/resendVerificationEmail`

	ErrGuestResendVerificationEmailUserNotFound           = `user not found`
	ErrGuestResendVerificationEmailTriggeredTooFrequently = `resend verification triggered to frequently`
	ErrGuestResendVerificationEmailUserAlreadyVerified    = `user already verified`
	ErrGuestResendVerificationEmailModificationFailed     = `resend verification modification failed`
)

var guestResendVerificationEmailLock = nsync.NewNamedMutex()

func (d *Domain) GuestResendVerificationEmail(in *GuestResendVerificationEmailIn) (out GuestResendVerificationEmailOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = in.Email

	if !guestResendVerificationEmailLock.TryLock(in.Email) {
		out.SetError(400, ErrGuestResendVerificationEmailTriggeredTooFrequently)
		return
	}

	if !user.FindByEmail() {
		guestResendVerificationEmailLock.Unlock(in.Email)
		out.SetError(400, ErrGuestResendVerificationEmailUserNotFound)
		return
	}
	out.actor = user.Id

	if user.VerifiedAt > 0 {
		guestResendVerificationEmailLock.Unlock(in.Email)
		out.SetError(400, ErrGuestResendVerificationEmailUserAlreadyVerified)
		return
	}

	recently := in.TimeNow().Add(-conf.ResendVerificationEmailThrottleMinute * time.Minute).Unix()
	if user.SecretCodeAt >= recently {
		guestResendVerificationEmailLock.Unlock(in.Email)
		out.SetError(400, ErrGuestResendVerificationEmailTriggeredTooFrequently)
		return
	}

	secretCode := id64.SID() + S.RandomCB63(1)
	user.SetSecretCode(secretCode)
	user.SetSecretCodeAt(in.UnixNow())
	hash := S.EncodeCB63(user.Id, 8)

	out.verifyEmailUrl = in.Host + `/` + GuestVerifyEmailAction + `?secretCode=` + secretCode + `&hash=` + hash
	d.runSubtask(func() {
		defer guestResendVerificationEmailLock.Unlock(in.Email)
		err := d.Mailer.SendRegistrationEmail(user.Email, out.verifyEmailUrl)
		L.IsError(err, `SendRegistrationEmail`)
		// TODO: insert failed event to clickhouse
	})

	if !user.DoUpdateById() {
		guestResendVerificationEmailLock.Unlock(in.Email)
		out.SetError(500, ErrGuestResendVerificationEmailModificationFailed)
		return
	}

	out.Ok = true
	return
}

type (
	GuestExternalAuthIn struct {
		RequestCommon
		Provider string `json:"provider" form:"provider" query:"provider" long:"provider" msg:"provider"`
	}
	GuestExternalAuthOut struct {
		ResponseCommon
		Link string `json:"link" form:"link" query:"link" long:"link" msg:"link"`

		// these for manual client-side oauth link generation
		ClientID    string   `json:"clientId" form:"clientId" query:"clientId" long:"clientId" msg:"clientId"`
		RedirectUrl string   `json:"redirectUrl" form:"redirectUrl" query:"redirectUrl" long:"redirectUrl" msg:"redirectUrl"`
		Scopes      []string `json:"scopes" form:"scopes" query:"scopes" long:"scopes" msg:"scopes"`
		CsrfState   string   `json:"csrfState" form:"csrfState" query:"csrfState" long:"csrfState" msg:"csrfState"`
	}
)

const (
	GuestExternalAuthAction = `guest/externalAuth`

	GuestExternalAuthProviderNotSet = `oauth provider not set`
	GuestExternalAuthInvalidUrl     = `oauth provider invalid url`
)

func (d *Domain) GuestExternalAuth(in *GuestExternalAuthIn) (out GuestExternalAuthOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	csrfState := in.Provider + `|`
	if in.SessionToken == `` {
		in.SessionToken = `TEMP__` + lexid.ID()
		out.SessionToken = in.SessionToken
	}
	csrfState += S.Left(in.SessionToken, 20)

	switch in.Provider {
	case OauthGoogle:
		provider := d.Oauth.Google[in.Host]
		if provider == nil {
			out.SetError(400, GuestExternalAuthInvalidUrl)
			return
		}
		out.Link = provider.AuthCodeURL(csrfState)
		out.ClientID = provider.ClientID
		out.RedirectUrl = provider.RedirectURL
		out.Scopes = provider.Scopes
		out.CsrfState = csrfState
	default:
		out.SetError(400, GuestExternalAuthProviderNotSet)
	}
	return
}

type (
	GuestOauthCallbackIn struct {
		RequestCommon
		State       string `json:"state" form:"state" query:"state" long:"state" msg:"state"`
		Code        string `json:"code" form:"code" query:"code" long:"code" msg:"code"`
		AccessToken string `json:"accessToken" form:"accessToken" query:"accessToken" long:"accessToken" msg:"accessToken"`
	}

	GuestOauthCallbackOut struct {
		ResponseCommon
		OauthUser   M.SX         `json:"oauthUser" form:"oauthUser" query:"oauthUser" long:"oauthUser" msg:"oauthUser"`
		Email       string       `json:"email" form:"email" query:"email" long:"email" msg:"email"`
		CurrentUser rqAuth.Users `json:"currentUser" form:"currentUser" query:"currentUser" long:"currentUser" msg:"currentUser"`
		Provider    string       `json:"provider" form:"provider" query:"provider" long:"provider" msg:"provider"`

		Segments M.SB `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
	}
)

const (
	GuestOauthCallbackAction = `guest/oauthCallback`

	ErrGuestOauthCallbackInvalidState           = `invalid csrf state`
	ErrGuestOauthCallbackInvalidCsrf            = `invalid csrf token`
	ErrGuestOauthCallbackInvalidUrl             = `invalid url`
	ErrGuestOauthCallbackFailedExchange         = `failed exchange oauth token`
	ErrGuestOauthCallbackFailedUserCreation     = `failed user creation from oauth`
	ErrGuestOauthCallbackFailedUserModification = `failed user modification from oauth`
	ErrGuestOauthCallbackFailedStoringSession   = `failed storing session from oauth`
)

func (d *Domain) GuestOauthCallback(in *GuestOauthCallbackIn) (out GuestOauthCallbackOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	csrf := S.RightOf(in.State, `|`)
	if csrf == `` {
		out.SetError(400, ErrGuestOauthCallbackInvalidState)
		return
	}

	L.Print(in.SessionToken)
	L.Print(csrf)
	if !S.StartsWith(in.SessionToken, csrf) {
		out.SetError(400, ErrGuestOauthCallbackInvalidCsrf)
		return
	}

	out.Provider = S.LeftOf(in.State, `|`)

	switch out.Provider {
	case OauthGoogle:
		provider := d.Oauth.Google[in.Host]
		if provider == nil {
			L.Print(in.Host)
			out.SetError(400, ErrGuestOauthCallbackInvalidUrl)
			return
		}

		token, err := provider.Exchange(in.TracerContext, in.Code)
		if L.IsError(err, `google.provider.Exchange`) {
			out.SetError(400, ErrGuestOauthCallbackFailedExchange)
			return
		}

		client := provider.Client(in.TracerContext, token)
		if d.googleUserInfoEndpointCache == `` {
			json := fetchJsonMap(client, `https://accounts.google.com/.well-known/openid-configuration`, &out.ResponseCommon)
			d.googleUserInfoEndpointCache = json.GetStr(`userinfo_endpoint`)
			if out.HasError() {
				return
			}
		}
		out.OauthUser = fetchJsonMap(client, d.googleUserInfoEndpointCache, &out.ResponseCommon)
		/* from google:
		{
			"email":			"",
			"email_verified":	true,
			"family_name":		"",
			"gender":			"",
			"given_name":		"",
			"locale":			"en-GB",
			"name":				"",
			"picture":			"http://",
			"profile":			"http://",
			"sub":				"number"
		} */
		if out.HasError() {
			return
		}
		out.Email = out.OauthUser.GetStr(`email`)
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = S.ValidateEmail(out.Email)

	if !user.FindByEmail() {
		// create user if not exists
		user.VerifiedAt = in.UnixNow()
		user.SetGenUniqueUsernameByEmail(user.Email, in.UnixNow())

		if !user.DoInsert() {
			out.SetError(500, ErrGuestOauthCallbackFailedUserCreation)
			return
		}
		out.actor = user.Id
	} else {
		out.actor = user.Id

		// update verifiedAt if not verified
		if user.VerifiedAt == 0 {
			user.SetVerifiedAt(in.UnixNow())
		}

		if !user.DoUpdateById() {
			out.SetError(500, ErrGuestOauthCallbackFailedUserModification)
			return
		}
	}

	d.expireSession(in.SessionToken, &out.ResponseCommon)

	// create new session
	session, sess := d.createSession(user.Id, user.Email, in.UserAgent)
	if !session.DoInsert() {
		out.SetError(500, ErrGuestOauthCallbackFailedStoringSession)
		return
	}
	out.SessionToken = session.SessionToken
	out.Segments = sess.Segments

	return
}
