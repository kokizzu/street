package xMailer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"street/conf"
)

func TestSengrid(t *testing.T) {
	t.Skip() // uncomment if need to check API key
	conf.LoadEnv()
	c := conf.EnvSendgrid()
	sg, err := NewSendgrid(c)
	assert.NoError(t, err)
	m := Mailer{
		SendMailFunc: sg.SendEmail,
	}
	// hardcoded for local, normally use WebConf.WebProtoDomain
	err = m.SendRegistrationEmail(c.DefaultFromEmail, `http://localhost:1234/guest/loginVerify?secretCode=justTestingSendGrid`)
	assert.NoError(t, err)
}
