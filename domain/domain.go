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

	IsBgSvc bool // long running program
}

// will run in background if background service
func (d *Domain) runSubtask(subTask func()) {
	if d.IsBgSvc {
		go subTask()
	} else {
		subTask()
	}
}
