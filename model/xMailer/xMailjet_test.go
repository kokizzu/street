package xMailer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"street/conf"
)

func TestMailjet(t *testing.T) {
	t.Skip() // uncomment if need to check API key
	conf.LoadEnv()
	c := conf.EnvMailjet()
	sg, err := NewMailjet(c)
	assert.NoError(t, err)
	m := Mailer{
		SendMailFunc: sg.SendEmail,
	}
	// hardcoded for local, normally use WebConf.WebProtoDomain
	err = m.SendRegistrationEmail(c.DefaultFromEmail, `http://localhost:1234/guest/loginVerify?secretCode=justTestingMailJet`)
	assert.NoError(t, err)
}
