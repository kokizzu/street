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
	"street/model/mProperty"
	"street/model/mProperty/saProperty"
	"street/model/xGmap"
	"street/model/xMailer"
)

type Domain struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter

	PropOltp *Tt.Adapter
	PropOlap *Ch.Adapter

	BusinessOltp *Tt.Adapter

	StorOltp *Tt.Adapter

	Mailer xMailer.Mailer

	IsBgSvc bool // long-running program

	// 3rd party
	Oauth conf.OauthConf
	Gmap  xGmap.Gmap

	// oauth related cache
	googleUserInfoEndpointCache string

	// timed buffer
	authLogs *chBuffer.TimedBuffer
	scannedAreasLogs *chBuffer.TimedBuffer
	scannedPropsLogs *chBuffer.TimedBuffer
	viewedRoomsLogs *chBuffer.TimedBuffer

	// logger
	Log *zerolog.Logger

	// list of superadmin emails
	Superadmins M.SB
	UploadDir   string
	CacheDir    string

	// web related, need to be here so Mailer/CLI also know the domain/urls
	WebCfg conf.WebConf
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
	d.scannedAreasLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableScannedAreas])
	d.scannedPropsLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableScannedProperties])
	d.viewedRoomsLogs = chBuffer.NewTimedBuffer(d.PropOlap.DB, 100_000, 1*time.Second, saProperty.Preparators[mProperty.TableViewedRooms])
}

func (d *Domain) WaitTimedBufferFinalFlush() {
	<-d.authLogs.WaitFinalFlush
	d.Log.Debug().Msg(`timed buffer flushed`)
}

var defaultIP4 = net.ParseIP(`0.0.0.0`).To4()
var defaultIP6 = net.ParseIP(`0:0:0:0:0:0:0:0`).To16()

func (d *Domain) InsertActionLog(in *RequestCommon, out *ResponseCommon) bool {
	ip := net.ParseIP(in.IpAddress)
	ip4 := ip.To4()
	if ip4 == nil {
		ip4 = defaultIP4
	}
	ip6 := ip.To16()
	if ip6 == nil {
		ip6 = defaultIP6
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
		RefId: 			in.RefId,
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
		row.RefId,
	})
}

func (d *Domain) CloseTimedBuffer() {
	go d.authLogs.Close()
	d.WaitTimedBufferFinalFlush()
}

func (d *Domain) insertScannedAreas(in saProperty.ScannedAreas) bool {
	return d.scannedAreasLogs.Insert([]any{
		in.ActorId,
		in.CreatedAt,
		in.Latitude,
		in.Longitude,
		in.City,
		in.State,
	})
}

func (d *Domain) insertScannedProps(in saProperty.ScannedProperties) bool {
	return d.scannedPropsLogs.Insert([]any{
		in.ActorId,
		in.CreatedAt,
		in.CountryCode,
		in.PropertyId,
	})
}

func (d *Domain) insertViewedRooms(in saProperty.ViewedRooms) bool {
	return d.viewedRoomsLogs.Insert([]any{
		in.ActorId,
		in.CreatedAt,
		in.PropertyId,
		in.RoomLabel,
		in.Country,
	})
}