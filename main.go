package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/rs/zerolog"

	"street/conf"
	"street/model"
	"street/model/xMailer"
	"street/presentation"
)

var VERSION = ``
var log *zerolog.Logger

func main() {
	conf.VERSION = VERSION
	log = conf.InitLogger()

	err := godotenv.Overload(`.env`)
	L.PanicIf(err, `godotenv.Load .env`)
	err = godotenv.Overload(`.env.override`)
	L.PanicIf(err, `godotenv.Load .env.override`)

	args := os.Args
	if len(args) < 2 {
		L.Print(`must start with: run, web, cron, or migrate as first argument`)
		L.Print(args)
		return
	}

	// mailer
	var mailer xMailer.Mailer
	switch xMailer.GetMailer() {
	case `sendgrid`:
		sg := &xMailer.Sengrid{SendgridConf: conf.EnvSendgrid()}
		L.PanicIf(sg.Connect(), `Sengrid.Connect`)
		mailer.SendMailFunc = sg.SendEmail
	case `mailjet`:
		mj := &xMailer.Mailjet{MailjetConf: conf.EnvMailjet()}
		L.PanicIf(mj.Connect(), `Mailjet.Connect`)
		mailer.SendMailFunc = mj.SendEmail
	default: // use mailhog
		mh, err := xMailer.NewMailhog(conf.EnvMailhog())
		L.PanicIf(err, `NewMailhog`)
		mailer.SendMailFunc = mh.SendEmail
	}

	// connect to tarantool
	tConf := conf.EnvTarantool()
	tConn, err := tConf.Connect()
	L.PanicIf(err, `tConf.Connect`)
	defer tConn.Close()

	// connect to clickhouse
	cConf := conf.EnvClickhouse()
	cConn, err := cConf.Connect()
	L.PanicIf(err, `cConf.Connect`)
	defer cConn.Close()

	// start
	mode := S.ToLower(os.Args[1])
	switch mode {
	case `web`:
		ws := &presentation.WebServer{
			AuthOltp: tConn,
			AuthOlap: cConn,
			Log:      log,
			Cfg:      conf.EnvWebConf(),
			Mailer:   mailer,
		}
		ws.Start()
	case `cli`:
		cli := &presentation.CLI{
			AuthOltp: tConn,
			AuthOlap: cConn,
			Log:      log,
			Mailer:   mailer,
		}
		cli.Run(os.Args[2:])
	case `cron`:
		cron := &presentation.Cron{
			AuthOltp: tConn,
			AuthOlap: cConn,
			Log:      log,
			Mailer:   mailer,
		}
		cron.Start()
	case `migrate`:
		model.RunMigration(tConn, cConn)
	default:
		log.Error().Str(`mode`, mode).Msg(`unknown mode`)
	}

}
