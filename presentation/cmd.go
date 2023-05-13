package presentation

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/rs/zerolog"

	"street/domain"
	"street/model/xMailer"
)

type CLI struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
	Log      *zerolog.Logger
	Mailer   xMailer.Mailer
}

func (c *CLI) Run(args []string) {
	if len(args) < 1 {
		c.Log.Print(`must start with one of: `, allCommands)
		return
	}
	if len(args) < 2 {
		c.Log.Print(`must provide json payload`)
		return
	}

	b := &domain.Domain{
		AuthOltp: c.AuthOltp,
		AuthOlap: c.AuthOlap,
		Mailer:   c.Mailer,
	}
	cmdRun(b, args[0], []byte(args[1]))

}
