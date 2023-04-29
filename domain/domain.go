package domain

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

type Domain struct {
	UserOltp *Tt.Adapter
	UserOlap *Ch.Adapter
}
