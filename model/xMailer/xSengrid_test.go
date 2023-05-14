package xMailer

import (
	"testing"

	"github.com/zeebo/assert"

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
	err = m.SendRegistrationEmail(c.DefaultFromEmail, conf.WEB_PROTO_DOMAIN+`/guest/loginVerify?secretCode=justTesting`)
	assert.NoError(t, err)
}
