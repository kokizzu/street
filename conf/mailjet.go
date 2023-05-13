package conf

import (
	"os"
)

type MailjetConf struct {
	ApiKeyPublic     string
	ApiKeyPrivate    string
	DefaultFromEmail string
	DefaultFromName  string
	ReplyToEmail     string
}

func EnvMailjet() MailjetConf {
	return MailjetConf{
		ApiKeyPublic:     os.Getenv("MJ_APIKEY_PUBLIC"),
		ApiKeyPrivate:    os.Getenv("MJ_APIKEY_PRIVATE"),
		DefaultFromEmail: os.Getenv("MAILER_DEFAULT_FROM_EMAIL"),
		DefaultFromName:  os.Getenv("MAILER_DEFAULT_FROM_NAME"),
		ReplyToEmail:     os.Getenv("MAILER_REPLY_TO_EMAIL"),
	}
}
