package conf

import (
	"os"

	"github.com/kokizzu/gotro/S"
)

type MailhogConf struct {
	MailhogHost      string
	MailhogPort      int
	DefaultFromEmail string
	DefaultFromName  string
	ReplyToEmail     string
}

func EnvMailhog() MailhogConf {
	return MailhogConf{
		MailhogHost:      os.Getenv("MAILHOG_HOST"),
		MailhogPort:      S.ToInt(os.Getenv("MAILHOG_PORT")),
		DefaultFromEmail: os.Getenv("MAILER_DEFAULT_FROM_EMAIL"),
		DefaultFromName:  os.Getenv("MAILER_DEFAULT_FROM_NAME"),
		ReplyToEmail:     os.Getenv("MAILER_REPLY_TO_EMAIL"),
	}
}
