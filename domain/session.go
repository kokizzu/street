package domain

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/lexid"
	"github.com/kpango/fastime"
	"github.com/mojura/enkodo"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/zeebo/xxh3"

	"street/conf"
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
)

type Session struct {
	UserId    uint64
	ExpiredAt int64 // in seconds
	Email     string

	// not saved but retrieved from SUPERADMIN_EMAILS env
	IsSuperAdmin bool

	Segments M.SB
}

// list of first segment of url path, if empty then only /guest segment
const (
	GuestSegment   = `guest`
	UserSegment    = `user`
	BuyerSegment   = `buyer`
	AdminSegment   = `admin`
	RealtorSegment = `realtor`
)

func (s *Session) MarshalEnkodo(enc *enkodo.Encoder) (err error) {
	_ = enc.Uint64(s.UserId)
	_ = enc.Int64(s.ExpiredAt)
	_ = enc.String(s.Email)
	return
}

func (s *Session) UnmarshalEnkodo(dec *enkodo.Decoder) (err error) {
	if s.UserId, err = dec.Uint64(); err != nil {
		return
	}
	if s.ExpiredAt, err = dec.Int64(); err != nil {
		return
	}
	if s.Email, err = dec.String(); err != nil {
		return
	}
	return
}

func createHash(key1, key2 string) string {
	res := xxh3.HashString128(key1 + conf.PROJECT_NAME + key2) // PROJECT_NAME = salt, if you change this, all token will be invalidated
	const x = 256
	return string([]byte{
		byte(res.Hi >> (64 - 8) % x),
		byte(res.Hi >> (64 - 16) % x),
		byte(res.Hi >> (64 - 24) % x),
		byte(res.Hi >> (64 - 32) % x),
		byte(res.Hi >> (64 - 40) % x),
		byte(res.Hi >> (64 - 48) % x),
		byte(res.Hi >> (64 - 56) % x),
		byte(res.Hi >> (64 - 64) % x), // nolint: staticcheck
		byte(res.Lo >> (64 - 8) % x),
		byte(res.Lo >> (64 - 16) % x),
		byte(res.Lo >> (64 - 24) % x),
		byte(res.Lo >> (64 - 32) % x),
		byte(res.Lo >> (64 - 40) % x),
		byte(res.Lo >> (64 - 48) % x),
		byte(res.Lo >> (64 - 56) % x),
		byte(res.Lo >> (64 - 64) % x), // nolint: staticcheck
	})
}

const TokenSeparator = `|`

func (s *Session) Encrypt(userAgent string) string {
	key1 := lexid.NanoID()
	key2 := S.EncodeCB63(fnv1a.HashString64(userAgent), 1)
	block, err := aes.NewCipher([]byte(createHash(key1, key2)))
	L.PanicIf(err, `aes.NewCipher`)
	gcm, err := cipher.NewGCM(block)
	L.PanicIf(err, `cipher.NewGCM`)
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	L.PanicIf(err, `io.ReadFull`)
	buffer := bytes.NewBuffer(nil)
	w := enkodo.NewWriter(buffer)
	err = w.Encode(s)
	L.PanicIf(err, `w.Encode`)
	cipherText := gcm.Seal(nonce, nonce, buffer.Bytes(), nil)
	return key1 + TokenSeparator + hex.EncodeToString(cipherText) + TokenSeparator + key2
}

func (s *Session) Decrypt(sessionToken, userAgent string) bool {
	strs := strings.Split(sessionToken, TokenSeparator)
	tokenLen := len(strs)
	if tokenLen != 3 {
		L.Print(`sessionToken length mismatch: ` + I.ToStr(tokenLen) + ` value: ` + sessionToken)
		return false
	}
	uaHash := S.EncodeCB63(fnv1a.HashString64(userAgent), 1)
	if strs[2] != uaHash {
		L.Print(`userAgent mismatch: ` + strs[2])
		return false
	}
	data, err := hex.DecodeString(strs[1])
	if L.IsError(err, `hex.DecodeString`) {
		return false
	}
	key := []byte(createHash(strs[0], strs[2]))
	block, err := aes.NewCipher(key)
	if L.IsError(err, `aes.NewCipher`) {
		return false
	}
	gcm, err := cipher.NewGCM(block)
	if L.IsError(err, `cipher.NewGCM`) {
		return false
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		L.Print(`len(data) < nonceSize`)
		return false
	}
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if L.IsError(err, `gcm.Open`) {
		return false
	}
	err = enkodo.Unmarshal(plainText, s)
	return !L.IsError(err, `enkodo.Unmarshal`)
}

func (d *Domain) createSession(userId uint64, email string, userAgent string) (*wcAuth.SessionsMutator, *Session) {
	session := wcAuth.NewSessionsMutator(d.AuthOltp)
	session.UserId = userId
	sess := &Session{
		UserId:    userId,
		ExpiredAt: fastime.Now().AddDate(0, 0, conf.CookieDays).Unix(),
		Email:     email,
	}
	session.SessionToken = sess.Encrypt(userAgent)
	session.ExpiredAt = sess.ExpiredAt
	d.segmentsFromSession(sess)
	return session, sess
}

func (d *Domain) expireSession(token string, out *ResponseCommon) int64 {
	session := wcAuth.NewSessionsMutator(d.AuthOltp)
	session.SessionToken = token
	now := fastime.UnixNow()
	if session.FindBySessionToken() {
		out.SessionToken = conf.CookieLogoutValue
		if session.ExpiredAt > now {
			session.SetExpiredAt(now)
			if !session.DoUpdateBySessionToken() {
				out.SetError(500, ErrUserSessionRemovalFailed)
				return 0
			}
			return now
		}
		return session.ExpiredAt
	}
	return 0
}

const (
	ErrSessionTokenEmpty     = `sessionToken empty`
	ErrSessionTokenInvalid   = `sessionToken invalid`
	ErrSessionTokenExpired   = `sessionToken expired`
	ErrSessionTokenNotFound  = `sessionToken not found`
	ErrSessionTokenLoggedOut = `sessionToken already logged out`

	ErrSegmentNotAllowed = `session segment not allowed`

	ErrSessionUserNotSuperAdmin = `session email is not superadmin`
)

func (d *Domain) MustLogin(in RequestCommon, out *ResponseCommon) (res *Session) {
	if in.SessionToken == `` {
		out.SetError(403, ErrSessionTokenEmpty)
		return nil
	}
	defer func() {
		if res == nil {
			// force user to clear cookie
			out.SessionToken = conf.CookieLogoutValue
		}
	}()
	sess := &Session{}
	if !sess.Decrypt(in.SessionToken, in.UserAgent) {
		out.SetError(403, ErrSessionTokenInvalid)
		return nil
	}
	now := fastime.UnixNow()
	if sess.ExpiredAt < now {
		out.SetError(403, ErrSessionTokenExpired)
		return nil
	}

	session := rqAuth.NewSessions(d.AuthOltp)
	session.SessionToken = in.SessionToken
	if !session.FindBySessionToken() {
		out.SetError(403, ErrSessionTokenNotFound)
		return nil
	}
	if session.ExpiredAt <= now {
		out.SetError(403, ErrSessionTokenLoggedOut)
		return nil
	}

	// fill segment
	d.segmentsFromSession(sess)

	// check allowed to hit/access this segment
	firstSegment := in.FirstSegment()
	if firstSegment != `` && // never filled on test
		!sess.IsSuperAdmin &&
		!sess.Segments[firstSegment] {
		out.SetError(403, ErrSegmentNotAllowed)
		return nil
	}

	out.actor = sess.UserId
	return sess
}

func (d *Domain) MustAdmin(in RequestCommon, out *ResponseCommon) (sess *Session) {
	sess = d.MustLogin(in, out)
	if sess == nil {
		return nil
	}
	if !sess.IsSuperAdmin {
		out.SetError(403, ErrSessionUserNotSuperAdmin)
		return nil
	}
	sess.IsSuperAdmin = true
	return sess
}
