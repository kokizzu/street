package domain

import (
	"log"
	"net"
	"time"

	chBuffer "github.com/kokizzu/ch-timed-buffer"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/S"

	"street/conf"
	"street/model/mAuth"
	"street/model/mAuth/saAuth"
	"street/model/xMailer"
)

type Domain struct {
	AuthOltp     *Tt.Adapter
	AuthOlap     *Ch.Adapter
	SendMailFunc xMailer.SendMailFunc
	Mailer       xMailer.Mailer

	IsBgSvc bool // long-running program
	Oauth   conf.OauthConf

	GoogleUserInfoEndpointCache string

	// timed buffer
	authLogs *chBuffer.TimedBuffer
}

// will run in background if background service
func (d *Domain) runSubtask(subTask func()) {
	if d.IsBgSvc {
		go subTask()
	} else {
		subTask()
	}
}

func (d *Domain) InitTimedBuffer() {
	d.authLogs = chBuffer.NewTimedBuffer(d.AuthOlap.DB, 100_000, 1*time.Second, saAuth.Preparators[mAuth.TableActionLogs])
}

func (d *Domain) WaitTimedBufferFinalFlush() {
	<-d.authLogs.WaitFinalFlush
	log.Println(`timed buffer flushed`)
}

func (d *Domain) InsertActionLog(in *RequestCommon, out *ResponseCommon) bool {
	ip := net.ParseIP(in.IpAddress)
	ip4 := S.IfEmpty(ip.To4().String(), `0.0.0.0`)
	ip6 := S.IfEmpty(ip.To16().String(), `0:0:0:0:0:0:0:0`)
	row := saAuth.ActionLogs{
		CreatedAt:  in.TimeNow(),
		RequestId:  in.RequestId,
		ActorId:    out.actor,
		Action:     in.Action,
		StatusCode: int16(out.StatusCode),
		Traces:     out.Traces(),
		Error:      out.Error,
		IpAddr4:    ip4,
		IpAddr6:    ip6,
		UserAgent:  in.UserAgent,
	}
	return d.authLogs.Insert([]any{
		row.CreatedAt,
		row.RequestId,
		row.ActorId,
		row.Action,
		row.StatusCode,
		row.Traces,
		row.Error,
		row.IpAddr4,
		row.IpAddr6,
		row.UserAgent,
	})
}
