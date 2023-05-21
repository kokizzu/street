package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"street/conf"
	"street/model"
	"street/model/xMailer"
	"street/presentation"
)

var VERSION = ``
var log *zerolog.Logger

func main() {
	conf.VERSION = VERSION

	// note: set instancec id when there's multiple instance
	// lexid.Config.Separator = `~[ID]`

	fmt.Println(conf.PROJECT_NAME + ` ` + S.IfEmpty(VERSION, `local-dev`))

	log = conf.InitLogger()
	conf.LoadEnv()

	args := os.Args
	if len(args) < 2 {
		L.Print(`must start with: run, web, cron, migrate, or config as first argument`)
		L.Print(args)
		return
	}

	if args[1] == `config` {
		L.Describe(M.SX{
			`web`:        conf.EnvWebConf(),
			`tarantool`:  conf.EnvTarantool(),
			`clickhouse`: conf.EnvClickhouse(),
			`sendgrid`:   conf.EnvSendgrid(),
			`mailjet`:    conf.EnvMailjet(),
			`mailhog`:    conf.EnvMailhog(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	var closers []func() error

	// mailer
	var mailer xMailer.Mailer
	eg.Go(func() error {
		mailerCfg := conf.EnvMailer()
		fmt.Println(`mailer: ` + mailerCfg.DefaultMailer)
		switch mailerCfg.DefaultMailer {
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
		return nil
	})

	var err error

	// connect to tarantool
	var tConn *Tt.Adapter
	eg.Go(func() error {
		tConf := conf.EnvTarantool()
		tConn, err = tConf.Connect()
		if tConn != nil {
			closers = append(closers, tConn.Close)
			fmt.Println(`tarantool connected: ` + tConf.DebugStr())
		}
		return err
	})

	// connect to clickhouse
	var cConn *Ch.Adapter
	eg.Go(func() error {
		cConf := conf.EnvClickhouse()
		cConn, err = cConf.Connect()
		if cConn != nil {
			closers = append(closers, cConn.Close)
			fmt.Println(`clickhouse connected: ` + cConf.DebugStr())
		}
		return err
	})

	oauth := conf.EnvOauth()

	L.PanicIf(eg.Wait(), `eg.Wait`)
	for _, closer := range closers {
		closer := closer
		defer closer()
	}

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
			Oauth:    oauth,
		}
		ws.Start()
	case `cli`:
		cli := &presentation.CLI{
			AuthOltp: tConn,
			AuthOlap: cConn,
			Log:      log,
			Mailer:   mailer,
			Oauth:    oauth,
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
