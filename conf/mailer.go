package conf

import (
	"os"
)

type MailerConf struct {
	DefaultFromEmail string
	DefaultFromName  string
	ReplyToEmail     string
}

func EnvMailer() MailerConf {
	return MailerConf{
		DefaultFromEmail: os.Getenv("MAILER_DEFAULT_FROM_EMAIL"),
		DefaultFromName:  os.Getenv("MAILER_DEFAULT_FROM_NAME"),
		ReplyToEmail:     os.Getenv("MAILER_REPLY_TO_EMAIL"),
	}
}
