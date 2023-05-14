package conf

import (
	"os"
)

type MailjetConf struct {
	MailerConf
	ApiKeyPublic  string
	ApiKeyPrivate string
}

func EnvMailjet() MailjetConf {
	return MailjetConf{
		ApiKeyPublic:  os.Getenv("MAILJET_APIKEY_PUBLIC"),
		ApiKeyPrivate: os.Getenv("MAILJET_APIKEY_PRIVATE"),
		MailerConf:    EnvMailer(),
	}
}
