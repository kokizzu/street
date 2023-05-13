package domain

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"

	"street/model/xMailer"
)

type Domain struct {
	AuthOltp     *Tt.Adapter
	AuthOlap     *Ch.Adapter
	SendMailFunc xMailer.SendMailFunc
	Mailer       xMailer.Mailer
}
