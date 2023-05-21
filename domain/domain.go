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
	d.authLogs = chBuffer.NewTimedBuffer(d.AuthOlap.DB, 100_000, 1*time.Second, saAuth.Preparators[mAuth.TableUserLogs])
}

func (d *Domain) WaitTimedBufferFinalFlush() {
	<-d.authLogs.WaitFinalFlush
	log.Println(`timed buffer flushed`)
}

func (d *Domain) InsertActionLog(in *RequestCommon, out *ResponseCommon) bool {
	ip := net.ParseIP(in.IpAddress)
	ip4 := S.IfEmpty(ip.To4().String(), `0.0.0.0`)
	ip6 := S.IfEmpty(ip.To16().String(), `0:0:0:0:0:0:0:0`)
	row := saAuth.UserLogs{
		Adapter:    nil,
		CreatedAt:  in.TimeNow(),
		RequestId:  in.RequestId,
		Error:      out.Error,
		ActorId:    out.actor,
		IpAddr4:    ip4,
		IpAddr6:    ip6,
		UserAgent:  in.UserAgent,
		Action:     in.Action,
		Traces:     out.Traces(),
		StatusCode: int16(out.StatusCode),
	}
	return d.authLogs.Insert([]any{
		row.CreatedAt,
		row.RequestId,
		row.Error,
		row.ActorId,
		row.IpAddr4,
		row.IpAddr6,
		row.UserAgent,
		row.Action,
		row.Traces,
		row.StatusCode,
	})
}
