package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"street/conf"
	"street/domain"
	"street/model"
	"street/model/xMailer"
	"street/model/zImport"
	"street/presentation"
)

var VERSION = ``
var log *zerolog.Logger

func main() {
	conf.VERSION = VERSION

	// note: set instance id when there's multiple instance
	// lexid.Config.Separator = `~THE_INSTANCE_ID`

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

	L.PanicIf(eg.Wait(), `eg.Wait`) // if error, make sure no error on: docker compose up
	for _, closer := range closers {
		closer := closer
		defer closer()
	}

	// create domain object
	d := &domain.Domain{
		AuthOltp: tConn,
		AuthOlap: cConn,
		PropOltp: tConn,
		PropOlap: cConn,
		StorOltp: tConn,
		Mailer: xMailer.Mailer{
			SendMailFunc: mailer.SendMailFunc,
		},
		IsBgSvc: false,
		Oauth:   oauth,
		Log:     log,

		UploadDir: conf.UploadDir(),

		Superadmins: conf.EnvSuperAdmins(),
	}

	mode := S.ToLower(os.Args[1])

	// check table existence
	if mode != `migrate` {
		model.VerifyTables(tConn, cConn, tConn, cConn, tConn)
	}

	// start
	switch mode {
	case `web`:
		ws := &presentation.WebServer{
			Domain: d,
			Cfg:    conf.EnvWebConf(),
		}
		ws.Start(log)
	case `cli`:
		cli := &presentation.CLI{
			Domain: d,
		}
		cli.Run(os.Args[2:], log)
	case `cron`:
		cron := &presentation.Cron{
			Domain: d,
		}
		cron.Start(log)
	case `migrate`:
		model.RunMigration(log, tConn, cConn, tConn, cConn, tConn)
	case `import`:
		excelFile, _ := filepath.Abs(`./static/house_data/House_Data_Full_Version_v1.xlsx`)
		jsonCoordFile, _ := filepath.Abs(`./static/house_data/coordinates.json`)
		zImport.ImportExcelData(tConn, excelFile, jsonCoordFile)
		zImport.PatchPropertiesPrice(tConn)
		getStreetViewImg("800", "400", "47.5763831", "-122.4211769", "85", "70", "3")
	case `import_location`:
		zImport.ImportHouseLocation(tConn)
	//case `upgradememtx`:
	//	zUpgrade.UserSessionToMemtx(tConn)
	//case `fix_time`:
	//	zImport.FixCreatedUpdatedAt(tConn)
	//case `patch_property_price`:
	//	zImport.PatchPropertiesPrice(tConn)
	//case `patch_serial_number_history`:
	//	zImport.PatchSerialNumberForHouseHistory(tConn)
	default:
		log.Error().Str(`mode`, mode).Msg(`unknown mode`)
	}

}

func getStreetViewImg(width, height, lat, lng, fov, heading, pitch string) {
	const gMapsApiKey string = "AIzaSyBKF5w6NExgYbmNMvlbMqF6sH2X4dFvMBg"
	url := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/streetview?size=%sx%s&location=%s,%s&fov=%s&heading=%s&pitch=%s&key=%s",
		width, height, lat, lng, fov, heading, pitch, gMapsApiKey,
	)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP Request failed with status: ", resp.Status)
	}
	file, err := os.Create("tmp/streetview_img.png")
	if err != nil {
		fmt.Println("Failed")
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Success")
}
