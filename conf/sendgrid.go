package conf

import (
	"os"
)

type SendgridConf struct {
	ApiKey           string
	DefaultFromEmail string
	DefaultFromName  string
	ReplyToEmail     string
}

func EnvSendgrid() SendgridConf {
	return SendgridConf{
		ApiKey:           os.Getenv("SENDGRID_API_KEY"),
		DefaultFromEmail: os.Getenv("MAILER_DEFAULT_FROM_EMAIL"),
		DefaultFromName:  os.Getenv("MAILER_DEFAULT_FROM_NAME"),
		ReplyToEmail:     os.Getenv("MAILER_REPLY_TO_EMAIL"),
	}
}
