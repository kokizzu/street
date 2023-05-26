package domain

import (
	"math"
	"testing"

	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/id64"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"street/model/mAuth/rqAuth"
)

func TestGuestRegister(t *testing.T) {
	d := testDomain()

	t.Run("emptyInput", func(t *testing.T) {
		in := GuestRegisterIn{}
		out := d.GuestRegister(&in)
		assert.Equal(t, out.Error, ErrGuestRegisterEmailInvalid)
	})

	t.Run("invalidEmail", func(t *testing.T) {
		in := GuestRegisterIn{
			Email: "09s8djmf6",
		}
		out := d.GuestRegister(&in)
		assert.Equal(t, out.Error, ErrGuestRegisterEmailInvalid)
	})

	t.Run("validEmail,noPass", func(t *testing.T) {
		in := GuestRegisterIn{
			Email: "a@b.c",
		}
		out := d.GuestRegister(&in)
		assert.Equal(t, out.Error, ErrGuestRegisterPasswordTooShort)
	})

	t.Run("validEmail,shortPass", func(t *testing.T) {
		in := GuestRegisterIn{
			Email:    "a@b.c",
			Password: "12345678",
		}
		out := d.GuestRegister(&in)
		assert.Equal(t, out.Error, ErrGuestRegisterPasswordTooShort)
	})

	t.Run("validEmail,validPass", func(t *testing.T) {
		in := GuestRegisterIn{
			Email:    "a@b.c",
			Password: "123456789012",
		}
		out := d.GuestRegister(&in)
		require.Empty(t, out.Error)
		require.NotZero(t, out.User.Id)

		secretCode, hash := parseSecretCodeHashUrl(t, out.verifyEmailUrl)

		t.Run(`verifyEmail,invalidHash`, func(t *testing.T) {
			in := GuestVerifyEmailIn{
				SecretCode: secretCode,
				Hash:       `#*()@&MN%#9834`,
			}
			out := d.GuestVerifyEmail(&in)
			assert.Equal(t, out.Error, ErrGuestVerifyEmailInvalidHash)
		})

		t.Run(`verifyEmail,invalidSecretCode`, func(t *testing.T) {
			in := GuestVerifyEmailIn{
				SecretCode: `@#*)($*@#%)(@`,
				Hash:       hash,
			}
			out := d.GuestVerifyEmail(&in)
			assert.Equal(t, out.Error, ErrGuestVerifyEmailSecretCodeMismatch)
		})

		t.Run(`verifyEmail,valid`, func(t *testing.T) {
			in := GuestVerifyEmailIn{
				SecretCode: secretCode,
				Hash:       hash,
			}
			out := d.GuestVerifyEmail(&in)
			require.Empty(t, out.Error)
			assert.True(t, out.Ok)
		})

		t.Run(`verifyEmail,alreadyVerified`, func(t *testing.T) {
			in := GuestVerifyEmailIn{
				SecretCode: `whatever`,
				Hash:       hash,
			}
			out := d.GuestVerifyEmail(&in)
			require.Empty(t, out.Error)
		})
	})
}

func TestGuestRequestVerificationEmail(t *testing.T) {
	d := testDomain()

	const email = "a@b.e"
	in := GuestRegisterIn{
		Email:    email,
		Password: "123456789012",
	}
	out := d.GuestRegister(&in)
	require.Empty(t, out.Error)
	require.NotZero(t, out.User.Id)

	t.Run("userNotFound", func(t *testing.T) {
		in := GuestResendVerificationEmailIn{
			Email: "x@x.com",
		}
		out := d.GuestResendVerificationEmail(&in)
		assert.Equal(t, out.Error, ErrGuestResendVerificationEmailUserNotFound)
	})

	t.Run("firstRequest", func(t *testing.T) {
		in := GuestResendVerificationEmailIn{
			Email: email,
		}
		out := d.GuestResendVerificationEmail(&in)
		require.Empty(t, out.Error)
		assert.True(t, out.Ok)
		secret, hash := parseSecretCodeHashUrl(t, out.verifyEmailUrl)

		t.Run("triggeredAgain", func(t *testing.T) {
			in := GuestResendVerificationEmailIn{
				Email: email,
			}
			out := d.GuestResendVerificationEmail(&in)
			assert.Equal(t, out.Error, ErrGuestResendVerificationEmailTriggeredTooFrequently)
		})

		t.Run(`verify`, func(t *testing.T) {
			in := GuestVerifyEmailIn{
				SecretCode: secret,
				Hash:       hash,
			}
			out := d.GuestVerifyEmail(&in)
			require.Empty(t, out.Error)

			t.Run(`requestVerificationEmail,alreadyVerified`, func(t *testing.T) {
				in := GuestResendVerificationEmailIn{
					Email: email,
				}
				out := d.GuestResendVerificationEmail(&in)
				assert.Equal(t, out.Error, ErrGuestResendVerificationEmailUserAlreadyVerified)
			})
		})
	})

}

func TestGuestLogin(t *testing.T) {
	d := testDomain()

	t.Run("emptyInput", func(t *testing.T) {
		in := GuestLoginIn{}
		out := d.GuestLogin(&in)
		assert.Equal(t, out.Error, ErrGuestLoginEmailInvalid)
	})

	t.Run("invalidEmail", func(t *testing.T) {
		in := GuestLoginIn{
			Email: "09s8djmf6",
		}
		out := d.GuestLogin(&in)
		assert.Equal(t, out.Error, ErrGuestLoginEmailInvalid)
	})

	var reg rqAuth.Users
	in := GuestRegisterIn{
		Email:    "a@c.id",
		Password: "test12345678",
	}
	out := d.GuestRegister(&in)
	reg = out.User

	t.Run("validEmail,userNotFound", func(t *testing.T) {
		in := GuestLoginIn{
			Email:    "b@c.id",
			Password: "test",
		}
		out := d.GuestLogin(&in)
		assert.Equal(t, out.Error, ErrGuestLoginEmailOrPasswordIncorrect)
	})

	t.Run("validEmail,wrongPass", func(t *testing.T) {
		in := GuestLoginIn{
			Email:    "a@c.id",
			Password: "test",
		}
		out := d.GuestLogin(&in)
		assert.Equal(t, out.Error, ErrGuestLoginPasswordOrEmailIncorrect)
	})

	t.Run("validEmail,validPass", func(t *testing.T) {
		in := GuestLoginIn{
			Email:    "a@c.id",
			Password: "test12345678",
		}
		out := d.GuestLogin(&in)
		require.Empty(t, out.Error)
		assert.Equal(t, out.User.Id, reg.Id)
		require.NotEmpty(t, out.SessionToken)
	})
}

func TestForgotResetPassword(t *testing.T) {
	d := testDomain()

	email := id64.SID() + `@reset`
	const pass = `012345678901`

	// register first
	in := GuestRegisterIn{
		Email:    email,
		Password: pass,
	}
	out := d.GuestRegister(&in)
	require.Empty(t, out.Error)

	t.Run(`forgotPassword`, func(t *testing.T) {
		in := &GuestForgotPasswordIn{
			Email: email,
		}
		out := d.GuestForgotPassword(in)
		require.Empty(t, out.Error)
		require.NotEmpty(t, out.resetPassUrl)

		t.Run(`forgotPasswordAgain`, func(t *testing.T) {
			out := d.GuestForgotPassword(in)
			assert.Equal(t, out.Error, ErrGuestForgotPasswordTriggeredTooFrequently)
		})

		secretCode, hash := parseSecretCodeHashUrl(t, out.resetPassUrl)

		t.Run(`resetPassword,invalidHash`, func(t *testing.T) {
			in := &GuestResetPasswordIn{
				SecretCode: secretCode,
				Hash:       `!@#$%^&*()_+{}:"><?~"`,
			}
			out := d.GuestResetPassword(in)
			assert.Equal(t, out.Error, ErrGuestResetPasswordInvalidHash)
		})

		t.Run(`resetPassword,userNotFound`, func(t *testing.T) {
			in := &GuestResetPasswordIn{
				SecretCode: secretCode,
				Hash:       S.EncodeCB63[uint64](math.MaxInt64, 0),
			}
			out := d.GuestResetPassword(in)
			assert.Equal(t, out.Error, ErrGuestResetPasswordUserNotFound)
		})

		t.Run(`resetPassword,passTooShort`, func(t *testing.T) {
			in := &GuestResetPasswordIn{
				SecretCode: secretCode,
				Hash:       hash,
			}
			out := d.GuestResetPassword(in)
			assert.Equal(t, out.Error, ErrGuestResetPasswordTooShort)
		})

		t.Run(`resetPassword,invalidSecretCode`, func(t *testing.T) {
			in := &GuestResetPasswordIn{
				SecretCode: `1234567`,
				Hash:       hash,
				Password:   pass,
			}
			out := d.GuestResetPassword(in)
			assert.Equal(t, out.Error, ErrGuestResetPasswordWrongSecret)
		})

		t.Run(`resetPassword,valid`, func(t *testing.T) {
			in := &GuestResetPasswordIn{
				SecretCode: secretCode,
				Hash:       hash,
				Password:   pass,
			}
			out := d.GuestResetPassword(in)
			require.Empty(t, out.Error)

			t.Run(`resetPasswordAgain`, func(t *testing.T) {
				out := d.GuestResetPassword(in)
				assert.Equal(t, out.Error, ErrGuestResetPasswordExpiredLink)
			})
		})
	})

}

func parseSecretCodeHashUrl(t *testing.T, url string) (secretCode string, hash string) {
	resetToken := S.RightOf(url, `?`)
	tokens := S.Split(resetToken, `&`)
	require.Len(t, tokens, 2)
	secretCode = S.RightOf(tokens[0], `=`)
	hash = S.RightOf(tokens[1], `=`)
	return
}
