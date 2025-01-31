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
			`mailhog`:    conf.EnvMailhog(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	eg, _ := errgroup.WithContext(ctx)
	var closers []func() error

	// mailer
	var mailer xMailer.Mailer
	eg.Go(func() error {
		mailer.Conf = conf.EnvMailer()
		switch mailer.Conf.DefaultMailer {
		case `sendgrid`:
			sg := &xMailer.Sengrid{SendgridConf: conf.EnvSendgrid()}
			L.PanicIf(sg.Connect(), `Sengrid.Connect`)
			mailer.SendMailFunc = sg.SendEmail
		case `mailjet`:
			mj := &xMailer.Mailjet{MailjetConf: conf.EnvMailjet()}
			L.PanicIf(mj.Connect(), `Mailjet.Connect`)
			mailer.SendMailFunc = mj.SendEmail
		case `dockermailserver`:
			dms := xMailer.Dockermailserver{DockermailserverConf: conf.EnvDockermailserver()}
			L.PanicIf(dms.Connect(), `Dockermailserver.Connect`)
			mailer.SendMailFunc = dms.SendEmail
		case `mailhog`:
			mh, err := xMailer.NewMailhog(conf.EnvMailhog())
			L.PanicIf(err, `NewMailhog`)
			mailer.SendMailFunc = mh.SendEmail
		default:
			L.Panic(`unknown DefaultMailer`)
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
		AuthOltp:     tConn,
		AuthOlap:     cConn,
		PropOltp:     tConn,
		PropOlap:     cConn,
		BusinessOltp: tConn,
		StorOltp:     tConn,
		Mailer:       mailer,
		IsBgSvc:      false,
		Oauth:        oauth,
		Gmap:         gmap,
		Log:          log,

		UploadDir: conf.UploadDir(),
		CacheDir:  conf.CacheDir(),

		Superadmins: conf.EnvSuperAdmins(),

		WebCfg: conf.EnvWebConf(),
	}
	d.InitTimedBuffer()
	defer d.CloseTimedBuffer()

	mode := S.ToLower(os.Args[1])

	// check table existence
	if mode != `migrate` {
		L.Print(`verifying table schema, if failed, run: go run main.go migrate`)
		model.VerifyTables(tConn, cConn, tConn, cConn, tConn, tConn)
	}

	// start
	switch mode {
	case `web`:
		ws := &presentation.WebServer{
			Domain: d,
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
		model.RunMigration(log, tConn, cConn, tConn, cConn, tConn, tConn)
	case `import`: // 2023-06-01
		excelFile, _ := filepath.Abs(`./static/house_data/House_Data_Full_Version_v1.xlsx`)
		jsonCoordFile, _ := filepath.Abs(`./static/house_data/coordinates.json`)
		zImport.ImportExcelData(tConn, excelFile, jsonCoordFile)
		zImport.PatchPropertiesPrice(tConn)
	case `import_property_us`: // 2023-09-19
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
	case `import_property_us_new`: // 2024-10-21
		zImport.ReadPropertyUS_TruliaCom(tConn, `./static/house_data/house_data_trulia.com.tsv`)
		zImport.ReadPropertyUS_ZillowCom(tConn, `./static/house_data/house_data_zillow.com.tsv`)
		zImport.ReadPropertyUS_ZillowCom(tConn, `./static/house_data/house_data_zillow.com_2.tsv`)
	case `import_property_pt`: // 2024-11-13 lisbon/portugal
		zImport.ReadPropertyPT_RightmoveCoUk(tConn, `./static/house_data/rightmove.co.uk_lisbon_final.tsv`)
		zImport.ReadPropertyPT_RightmoveCoUk2(tConn, `./static/house_data/Rightmove.co.uk_Final.tsv`)
	case `import_property_pt_new`: // 2024-12-01 portugal
		zImport.ReadPropertyPT_Buy_ZomePT(tConn, `./static/house_data/Zome.pt_buy_Final.tsv`)
		zImport.ReadPropertyPT_Rent_ZomePT(tConn, `./static/house_data/Zome.pt_rent_Final.tsv`)
	case `import_property_uk_split_rent_buy`:
		zImport.ReadPropertyUK_Rent_RightmoveCoUk(tConn, `./static/house_data/Rightmove.co.uk_Rent_Final.tsv`)
		zImport.ReadPropertyUK_Sale_RightmoveCoUk(tConn, `./static/house_data/Rightmove.co.uk_Sale_Final.tsv`)
	case `import_property_history_us`: // 2023-11-12
		const baseUrl = "https://www.redfin.com/stingray/api/home/details/belowTheFold"
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
		zImport.ImportPropertyHistoryUsData(tConn, baseUrl, minPropertyIdNumber, maxPropertyIdNumber)
	case `import_property_tw`: // 2023-11-30
		// requirement: copy crawled data to static/property_tw_data/props_tw.jsonl
		zImport.ImportPropertyTwData(tConn, gmap)
	case `import_streetview_image_tw`: // 2023-12-12
		zImport.ImportStreetViewImageTW(d, gmap)
	case `migrate_property_us_image`: // 2023-11-12
		zImport.MigratePropertyUSImage(tConn, 4000001, 8000000)
	case `clean_excessive_attr_property_extra_us`: // 2023-10-15
		zImport.CleanExcessiveAttrPropertyExtraUs(tConn)
	case `import_location`: // 2023-07-02
		zImport.ImportHouseLocation(tConn, gmap)
	case `import_streetview_image`: // 2023-09-04
		zImport.ImportStreetViewImage(d, gmap)
	case `import_translation`:
		// https://docs.google.com/spreadsheets/d/1XnbE1ERv-jGjEOh-Feibtlb-drTjgzqOrcHqTCCmE3Y/edit#gid=0
		zImport.GoogleSheetTranslationToJson(`1XnbE1ERv-jGjEOh-Feibtlb-drTjgzqOrcHqTCCmE3Y`)
	case `import_geolocation`:
		zImport.ImportGeolocationDatabase(cConn, `./static/cities.tsv`)
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
