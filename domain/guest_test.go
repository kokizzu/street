package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"

	"street/model/mAuth/rqAuth"
)

func TestGuestRegister(t *testing.T) {
	d := Domain{
		AuthOltp: testTt,
		AuthOlap: testCh,
	}

	t.Run("emptyInput", func(t *testing.T) {
		in := GuestRegisterIn{}
		out := d.GuestRegister(&in)
		assert.Equal(t, out.Error, "email must be valid")
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
		assert.Equal(t, out.Error, "")
		require.NotZero(t, out.User.Id)
	})
}

func TestGuestLogin(t *testing.T) {
	d := Domain{
		AuthOltp: testTt,
		AuthOlap: testCh,
	}

	t.Run("emptyInput", func(t *testing.T) {
		in := GuestLoginIn{}
		out := d.GuestLogin(&in)
		assert.Equal(t, out.Error, "email must be valid")
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
		assert.Equal(t, out.Error, "")
		assert.Equal(t, out.User.Id, reg.Id)
		require.NotEmpty(t, out.SessionToken)
	})
}
