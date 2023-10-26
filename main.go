package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	"street/model/xGmap"
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
		case `dockermailserver`:
			dms, err := xMailer.NewDockermailserver(conf.EnvDockermailserver())
			L.PanicIf(err, `NewDockermailserver`)
			mailer.SendMailFunc = dms.SendEmail
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

	// init gmap
	gmapConf := conf.EnvGmap()
	gmap := xGmap.Gmap{gmapConf}

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
		Gmap:    gmap,
		Log:     log,

		UploadDir: conf.UploadDir(),

		Superadmins: conf.EnvSuperAdmins(),
	}
	d.InitTimedBuffer()
	defer d.CloseTimedBuffer()

	mode := S.ToLower(os.Args[1])

	// check table existence
	if mode != `migrate` {
		L.Print(`verifying table schema, if failed, run: go run main.go migrate`)
		model.VerifyTables(tConn, cConn, tConn, cConn, tConn)
	}

	// start
	switch mode {
	case `web`:
		ws := &presentation.WebServer{
			Domain: d,
			Cfg:    conf.EnvWebConf(),
		}
		conf.LoadCountries("./static/country_data/countries.tsv")
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
	case `import_property_us`:
		const baseUrl = "https://www.redfin.com/stingray/api/home/details/belowTheFold"
		// const minPropertyId = 1
		// const maxPropertyId = 10000000
		args := os.Args
		// Check the number of arguments.
		if len(args) < 2 {
			fmt.Println("Usage: go run main.go import_property_us -minPropertyId -maxPropertyId")
			return
		}

		// Process the arguments
		minPropertyId := args[2]
		maxPropertyId := args[3]

		minPropertyIdNumber, err := strconv.Atoi(minPropertyId)
		if err != nil {
			fmt.Println("Error: -minPropertyId is required a number and not a string")
			return
		}

		maxPropertyIdNumber, err := strconv.Atoi(maxPropertyId)
		if err != nil {
			fmt.Println("Error: -maxPropertyIdNumber is required a number and not a string")
			return
		}

		if minPropertyIdNumber >= maxPropertyIdNumber {
			fmt.Println("Error: -minPropertyId is required smaller than -maxPropertyIdNumber")
			return
		}

		zImport.ImportPropertyUsData(tConn, baseUrl, minPropertyIdNumber, maxPropertyIdNumber)
	case `clean_excessive_attr_property_extra_us`:
		zImport.CleanExcessiveAttrPropertyExtraUs(tConn)
	case `import_location`:
		zImport.ImportHouseLocation(tConn, gmap)
	case `import_streetview_image`:
		zImport.ImportStreetViewImage(d, gmap)
	case `import_translation`:
		// https://docs.google.com/spreadsheets/d/1XnbE1ERv-jGjEOh-Feibtlb-drTjgzqOrcHqTCCmE3Y/edit#gid=0
		zImport.GoogleSheetTranslationToJson(`1XnbE1ERv-jGjEOh-Feibtlb-drTjgzqOrcHqTCCmE3Y`)
	case `manual_test`: // how to manual test, it's better to use unit test, except for third party
		const UA = `LocalTesting/1.0`
		const sessionSavePath = `/tmp/session1.txt` // simulate cookie
		savedSession := L.ReadFile(sessionSavePath)
		if savedSession == `` { // if expired, please remove that file
			session, _ := d.CreateSession(1, `admin@localhost`, UA, `127.0.0.1`)
			if !session.DoInsert() {
				L.Print(`failed inserting session`)
				return
			}
			L.CreateFile(sessionSavePath, session.SessionToken)
			savedSession = session.SessionToken
		}
		rc := domain.NewLocalRequestCommon(savedSession, UA)
		out := d.UserGpsCountry(&domain.UserGpsCountryIn{
			RequestCommon: rc,
			CenterLat:     40.730610,
			CenterLong:    -73.935242,
			CheckOnly:     true,
		})
		L.Describe(out)
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
