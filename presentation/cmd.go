package presentation

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"github.com/rs/zerolog"

	"street/domain"
)

type CLI struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
	Log      *zerolog.Logger
}

func (c *CLI) Run(args []string) {
	if len(args) < 1 {
		c.Log.Print(`must start with: `, allCommands)
		return
	}

	b := &domain.Domain{
		UserOltp: c.AuthOltp,
		UserOlap: c.AuthOlap,
	}
	action := args[0]
	switch action {
	case domain.UserRegisterAction: //TODO: use codegen
		if len(args) < 2 {
			c.Log.Print(`must provide json payload`)
			return
		}
		in := domain.UserRegisterIn{}
		err := json.Unmarshal([]byte(args[1]), &in)
		if L.IsError(err, `json.Unmarshal`) {
			return
		}
		out := b.UserRegister(&in)
		fmt.Println(X.ToJsonPretty(out))
	}
}
