package presentation

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/rs/zerolog"
)

type Cron struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
	Log      *zerolog.Logger
}

func (c *Cron) Start() {

}
