package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
)

func TestUserRegister(t *testing.T) {
	d := Domain{
		UserOltp: testTt,
		UserOlap: testCh,
	}

	t.Run("emptyInput", func(t *testing.T) {
		in := UserRegisterIn{}
		out := d.UserRegister(&in)
		assert.Equal(t, out.Error, "email must be valid")
	})

	t.Run("invalidEmail", func(t *testing.T) {
		in := UserRegisterIn{
			Email: "09s8djmf6",
		}
		out := d.UserRegister(&in)
		assert.Equal(t, out.Error, ErrUserRegisterEmailInvalid)
	})

	t.Run("validEmail,noPass", func(t *testing.T) {
		in := UserRegisterIn{
			Email: "a@b.c",
		}
		out := d.UserRegister(&in)
		assert.Equal(t, out.Error, ErrUserRegisterPasswordTooShort)
	})

	t.Run("validEmail,shortPass", func(t *testing.T) {
		in := UserRegisterIn{
			Email:    "a@b.c",
			Password: "12345678",
		}
		out := d.UserRegister(&in)
		assert.Equal(t, out.Error, ErrUserRegisterPasswordTooShort)
	})

	t.Run("validEmail,validPass", func(t *testing.T) {
		in := UserRegisterIn{
			Email:    "a@b.c",
			Password: "123456789012",
		}
		out := d.UserRegister(&in)
		assert.Equal(t, out.Error, "")
		require.NotZero(t, out.User.Id)
	})
}
