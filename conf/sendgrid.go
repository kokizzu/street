package conf

import (
	"os"
)

type SendgridConf struct {
	MailerConf
	ApiKey string
}

func EnvSendgrid() SendgridConf {
	return SendgridConf{
		ApiKey:     os.Getenv("SENDGRID_API_KEY"),
		MailerConf: EnvMailer(),
	}
}
