package domain

import (
	"testing"

	"github.com/kokizzu/id64"
	"github.com/kpango/fastime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogout(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	email := id64.SID() + `@login`
	const pass = `012345678901`

	// register first
	in := GuestRegisterIn{
		Email:    email,
		Password: pass,
	}
	out := d.GuestRegister(&in)
	require.Empty(t, out.Error)
	require.NotZero(t, out.User.Id)

	t.Run(`emptySessionToken`, func(t *testing.T) {
		in := UserProfileIn{}
		out := d.UserProfile(&in)
		assert.Equal(t, out.Error, ErrSessionTokenEmpty)
	})

	t.Run(`invalidSessionToken`, func(t *testing.T) {
		in := UserProfileIn{
			RequestCommon: RequestCommon{
				SessionToken: `invalid`,
			},
		}
		out := d.UserProfile(&in)
		assert.Equal(t, out.Error, ErrSessionTokenInvalid)
	})

	t.Run(`expiredSessionToken`, func(t *testing.T) {
		const userAgent = ``
		sess := Session{
			UserId:    1,
			ExpiredAt: fastime.UnixNow() - 1,
			Email:     email,
		}

		in := UserProfileIn{
			RequestCommon: RequestCommon{
				SessionToken: sess.Encrypt(userAgent),
			},
		}
		out := d.UserProfile(&in)
		assert.Equal(t, out.Error, ErrSessionTokenExpired)
	})

	t.Run(`login`, func(t *testing.T) {
		in := GuestLoginIn{
			Email:    email,
			Password: pass,
		}
		out := d.GuestLogin(&in)
		require.Empty(t, out.Error)
		require.NotZero(t, out.User.Id)
		require.NotEmpty(t, out.SessionToken)

		sessionToken := out.SessionToken

		// check profile after login must succeed
		t.Run(`profile`, func(t *testing.T) {
			in := &UserProfileIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
			}
			out := d.UserProfile(in)
			require.Empty(t, out.Error)
			if out.User != nil {
				require.NotZero(t, out.User.Id)
			}
		})

		t.Run(`wrongUserAgent`, func(t *testing.T) {
			in := &UserProfileIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
					UserAgent:    `otherUserAgent`,
				},
			}
			out := d.UserProfile(in)
			assert.Equal(t, out.Error, ErrSessionTokenInvalid)
		})

		t.Run(`changePass,noNewPass`, func(t *testing.T) {
			in := &UserChangePasswordIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
			}
			out := d.UserChangePassword(in)
			assert.Equal(t, out.Error, ErrUserChangePasswordNewPassTooShort)
		})

		const newPass = `1234567890123`

		t.Run(`changePass,wrongOldPass`, func(t *testing.T) {
			in := &UserChangePasswordIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
				NewPass: newPass,
			}
			out := d.UserChangePassword(in)
			assert.Equal(t, out.Error, ErrUserChangePasswordWrongOldPass)
		})

		t.Run(`changePass,valid`, func(t *testing.T) {
			in := &UserChangePasswordIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
				NewPass: newPass,
				OldPass: pass,
			}
			out := d.UserChangePassword(in)
			assert.Empty(t, out.Error)

			t.Run(`loginAfterChangePass,useOldPass`, func(t *testing.T) {
				in := &GuestLoginIn{
					Email:    email,
					Password: pass,
				}
				out := d.GuestLogin(in)
				assert.Equal(t, out.Error, ErrGuestLoginPasswordOrEmailIncorrect)
			})

			t.Run(`loginAfterChangePass,newPass`, func(t *testing.T) {
				in := &GuestLoginIn{
					Email:    email,
					Password: newPass,
				}
				out := d.GuestLogin(in)
				assert.Empty(t, out.Error)

				newSessionToken := out.SessionToken

				t.Run(`deactivate,wrongPass`, func(t *testing.T) {
					in := &UserDeactivateIn{
						RequestCommon: RequestCommon{
							SessionToken: newSessionToken,
						},
						Password: pass,
					}
					out := d.UserDeactivate(in)
					assert.Equal(t, out.Error, ErrUserDeactivateWrongPass)
				})

				t.Run(`deactivate,correctPass`, func(t *testing.T) {
					in := &UserDeactivateIn{
						RequestCommon: RequestCommon{
							SessionToken: newSessionToken,
						},
						Password: newPass,
					}
					out := d.UserDeactivate(in)
					assert.Empty(t, out.Error)

					// TODO: check login again, should not able
					// TODO: use session to do profile after deactivate, should not able
				})

			})
		})

		t.Run(`logout`, func(t *testing.T) {
			in := &UserLogoutIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
			}
			out := d.UserLogout(in)
			require.Empty(t, out.Error)
			require.NotZero(t, out.LogoutAt)
			previousLogoutAt := out.LogoutAt

			t.Run(`profilAfterLogout`, func(t *testing.T) {
				in := &UserProfileIn{
					RequestCommon: RequestCommon{
						SessionToken: sessionToken,
					},
				}
				out := d.UserProfile(in)
				assert.Equal(t, out.Error, ErrSessionTokenLoggedOut)
			})

			// idempotent
			t.Run(`logoutAgain`, func(t *testing.T) {
				out := d.UserLogout(in)
				assert.Equal(t, out.Error, ``)
				assert.Equal(t, out.LogoutAt, previousLogoutAt)
			})

			t.Run(`changePassword,notLoggedIn`, func(t *testing.T) {
				in := &UserChangePasswordIn{
					RequestCommon: RequestCommon{
						SessionToken: sessionToken,
					},
				}
				out := d.UserChangePassword(in)
				assert.Equal(t, out.Error, ErrSessionTokenLoggedOut)
			})

			t.Run(`deactivate,notLoggedIn`, func(t *testing.T) {
				in := &UserDeactivateIn{
					RequestCommon: RequestCommon{
						SessionToken: sessionToken,
					},
					Password: pass,
				}
				out := d.UserDeactivate(in)
				assert.Equal(t, out.Error, ErrSessionTokenLoggedOut)
			})
		})
	})
}
