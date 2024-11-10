package domain

import (
	"testing"

	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/id64"
	"github.com/kpango/fastime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"street/model/mProperty/rqProperty"
	"street/model/zCrud"
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
		oldUsername := out.User.UserName
		oldEmail := out.User.Email
		oldFullname := out.User.FullName

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

		t.Run(`updateProfileUserName`, func(t *testing.T) {
			in := &UserUpdateProfileIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
				UserName: S.RandomPassword(10),
			}
			require.NotEqual(t, in.UserName, oldUsername)
			out := d.UserUpdateProfile(in)
			require.Empty(t, out.Error)
			if out.User != nil {
				require.NotZero(t, out.User.Id)
				require.Equal(t, out.User.UserName, in.UserName)
			}
		})

		t.Run(`updateProfileEmail`, func(t *testing.T) {
			in := &UserUpdateProfileIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
				Email: S.RandomPassword(10) + `@yahoo.com`,
			}
			require.NotEqual(t, in.Email, oldEmail)
			out := d.UserUpdateProfile(in)
			require.Empty(t, out.Error)
			if out.User != nil {
				require.NotZero(t, out.User.Id)
				require.Equal(t, out.User.Email, in.Email)
				require.Empty(t, out.User.VerifiedAt)
				email = out.User.Email
			}
		})

		t.Run(`updateProfileFullName`, func(t *testing.T) {
			in := &UserUpdateProfileIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
				FullName: S.RandomPassword(10),
			}
			require.NotEqual(t, in.FullName, oldFullname)
			out := d.UserUpdateProfile(in)
			require.Empty(t, out.Error)
			if out.User != nil {
				require.NotZero(t, out.User.Id)
				require.Equal(t, out.User.FullName, in.FullName)
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
			require.Empty(t, out.Error)

			t.Run(`loginAfterChangePass,useOldPass`, func(t *testing.T) {
				in := &GuestLoginIn{
					Email:    email,
					Password: pass,
				}
				out := d.GuestLogin(in)
				assert.Equal(t, out.Error, ErrGuestLoginPasswordIncorrect)
			})

			t.Run(`loginAfterChangePass,newPass`, func(t *testing.T) {
				in := &GuestLoginIn{
					Email:    email,
					Password: newPass,
				}
				out := d.GuestLogin(in)
				require.Empty(t, out.Error)

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
					require.Empty(t, out.Error)

					// TODO: check login again, should not able
					// TODO: use session to do profile after deactivate, should not able
				})

			})
		})

		t.Run(`userGpsCountry`, func(t *testing.T) {
			t.Skip(`skip call google API on CI`)
			in := &UserGpsCountryIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
				CenterLat:  -34.397,
				CenterLong: 150.644,
			}
			out := d.UserGpsCountry(in)
			assert.Equal(t, out.Error, ErrUserGpsCountryFailedGetCountry)
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

			t.Run(`updateProfileAfterLogout`, func(t *testing.T) {
				in := &UserUpdateProfileIn{
					RequestCommon: RequestCommon{
						SessionToken: sessionToken,
					},
					UserName: "ayaya",
					FullName: "",
					Email:    "",
				}
				out := d.UserUpdateProfile(in)
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

func TestUserSearchPropAndLikes(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	// create property as admin
	out1 := d.AdminProperties(&AdminPropertiesIn{
		RequestCommon: testAdminRequestCommon(AdminPropertiesAction),
		Cmd:           zCrud.CmdUpsert,
		Property: rqProperty.Property{
			UniqPropKey:            "TEST234",
			SerialNumber:           "1001",
			SizeM2:                 "200.5",
			MainUse:                "Residential",
			MainBuildingMaterial:   "Paper",
			ConstructCompletedDate: "105年9月7日",
			NumberOfFloors:         "20",
			Address:                "4-55 Matsushigecho, Nakamura Ward, Nagoya, Aichi 450-0004",
			Coord:                  []any{35.16, 136.9},
		},
		WithMeta: false,
	})

	assert.Empty(t, out1.Error)
	assert.NotNil(t, out1.Property)

	// find prop as admin
	outs1 := d.UserSearchProp(&UserSearchPropIn{
		RequestCommon: testAdminRequestCommon(UserSearchPropAction),
		CenterLat:     35.16,
		CenterLong:    136.9,
		Offset:        0,
	})
	assert.Empty(t, outs1.Error)
	assert.Len(t, outs1.Properties, 1)

	// create property as admin 2
	out2 := d.AdminProperties(&AdminPropertiesIn{
		RequestCommon: testAdminRequestCommon(AdminPropertiesAction),
		Cmd:           zCrud.CmdUpsert,
		Property: rqProperty.Property{
			UniqPropKey:            "TEST123",
			SerialNumber:           "1002",
			SizeM2:                 "123.5",
			MainUse:                "Residential",
			MainBuildingMaterial:   "Cardboard",
			ConstructCompletedDate: "105年12月4日",
			NumberOfFloors:         "10",
			Address:                "1-12 Matsushigecho, Nakamura Ward, Nagoya, Aichi 450-0004",
			Coord:                  []any{35.165, 136.89},
		},
		WithMeta: false,
	})

	assert.Empty(t, out2.Error)
	assert.NotNil(t, out2.Property)

	// find prop as admin
	outs2 := d.UserSearchProp(&UserSearchPropIn{
		RequestCommon: testAdminRequestCommon(UserSearchPropAction),
		CenterLat:     35.16,
		CenterLong:    136.89,
		Offset:        0,
	})
	assert.Empty(t, outs2.Error)
	assert.Len(t, outs2.Properties, 2)

	fetchLike := func(pid1, pid2 uint64) (likedMap map[uint64]bool, countMap map[uint64]int64) {
		upl := rqProperty.NewUserPropLikes(d.PropOltp)
		upl.UserId = testAdmin.Id
		likedMap = upl.LikedMap([]uint64{pid1, pid2})
		plc := rqProperty.NewPropLikeCount(d.PropOltp)
		countMap = plc.CountMap([]uint64{pid1, pid2})
		return

	}

	propId1 := out1.Property.Id
	propId2 := out2.Property.Id

	t.Run(`likeProp1`, func(t *testing.T) {
		in1 := &UserLikePropIn{
			RequestCommon: testAdminRequestCommon(UserLikePropAction),
			PropId:        propId1,
			Like:          true,
		}
		out := d.UserLikeProp(in1)
		assert.Empty(t, out.Error)

		// assert like
		likedMap, countMap := fetchLike(propId1, propId2)
		assert.Equal(t, likedMap[propId1], true)
		assert.Equal(t, countMap[propId1], int64(1))
		assert.Equal(t, likedMap[propId2], false)
		assert.Equal(t, countMap[propId2], int64(0))

		t.Run(`likeProp2`, func(t *testing.T) {
			in2 := &UserLikePropIn{
				RequestCommon: testAdminRequestCommon(UserLikePropAction),
				PropId:        out2.Property.Id,
				Like:          true,
			}
			out := d.UserLikeProp(in2)
			assert.Empty(t, out.Error)

			// assert like
			likedMap, countMap := fetchLike(propId1, propId2)
			assert.Equal(t, likedMap[propId1], true)
			assert.Equal(t, countMap[propId1], int64(1))
			assert.Equal(t, likedMap[propId2], true)
			assert.Equal(t, countMap[propId2], int64(1))

			t.Run(`unlikeProp2`, func(t *testing.T) {

				in := &UserLikePropIn{
					RequestCommon: testAdminRequestCommon(UserLikePropAction),
					PropId:        out2.Property.Id,
					Like:          false,
				}
				out := d.UserLikeProp(in)
				assert.Empty(t, out.Error)

				// being liked removed
				likedMap, countMap := fetchLike(propId1, propId2)
				assert.Equal(t, likedMap[propId1], true)
				assert.Equal(t, countMap[propId1], int64(1))
				assert.Equal(t, likedMap[propId2], false)
				assert.Equal(t, countMap[propId2], int64(0))
			})
		})

		t.Run(`unlikeProp1`, func(t *testing.T) {
			in := &UserLikePropIn{
				RequestCommon: testAdminRequestCommon(UserLikePropAction),
				PropId:        out1.Property.Id,
				Like:          false,
			}
			out := d.UserLikeProp(in)
			assert.Empty(t, out.Error)

			// both like removed
			likedMap, countMap := fetchLike(propId1, propId2)
			assert.Equal(t, likedMap[propId1], false)
			assert.Equal(t, countMap[propId1], int64(0))
			assert.Equal(t, likedMap[propId2], false)
			assert.Equal(t, countMap[propId2], int64(0))

		})

	})
}
