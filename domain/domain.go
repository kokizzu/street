package domain

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

type Domain struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
}
