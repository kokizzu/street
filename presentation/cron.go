package presentation

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/rs/zerolog"

	"street/model/xMailer"
)

type Cron struct {
	AuthOltp     *Tt.Adapter
	AuthOlap     *Ch.Adapter
	Log          *zerolog.Logger
	SendMailFunc xMailer.SendMailFunc
	Mailer       xMailer.Mailer
	PropOltp     *Tt.Adapter
	PropOlap     *Ch.Adapter
}

func (c *Cron) Start() {

}
