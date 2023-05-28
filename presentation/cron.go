package presentation

import (
	"github.com/rs/zerolog"

	"street/domain"
	"street/model/xMailer"
)

type Cron struct {
	*domain.Domain
	Log          *zerolog.Logger
	SendMailFunc xMailer.SendMailFunc
	Mailer       xMailer.Mailer
}

func (c *Cron) Start() {

}
