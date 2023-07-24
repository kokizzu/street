package domain

import (
	"net"
	"time"

	chBuffer "github.com/kokizzu/ch-timed-buffer"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/M"
	"github.com/rs/zerolog"

	"street/conf"
	"street/model/mAuth"
	"street/model/mAuth/saAuth"
	"street/model/xMailer"
)

type Domain struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter

	PropOltp *Tt.Adapter
	PropOlap *Ch.Adapter

	StorOltp *Tt.Adapter

	Mailer xMailer.Mailer

	IsBgSvc bool // long-running program
	Oauth   conf.OauthConf

	// oauth related cache
	googleUserInfoEndpointCache string

	// timed buffer
	authLogs *chBuffer.TimedBuffer

	// logger
	Log *zerolog.Logger

	// list of superadmin emails
	Superadmins M.SB

	UploadDir string
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
	d.Log.Debug().Msg(`timed buffer flushed`)
}

func (d *Domain) InsertActionLog(in *RequestCommon, out *ResponseCommon) bool {
	ip := net.ParseIP(in.IpAddress)
	v4 := ip.To4()
	var ip4, ip6 string
	if v4 == nil {
		ip4 = `0.0.0.0`
	} else {
		ip4 = v4.String()
	}
	v6 := ip.To16()
	if v6 == nil {
		ip6 = `0:0:0:0:0:0:0:0`
	} else {
		ip6 = v6.String()
	}
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
		Lat:        in.Lat,
		Long:       in.Long,
		Latency:    in.Latency(),
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
		row.Lat,
		row.Long,
		row.Latency,
	})
}
