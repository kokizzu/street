package domain

import (
	"database/sql"
	"fmt"
	"net"
	"testing"

	"github.com/kokizzu/gotro/D"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/ory/dockertest/v3"
	"github.com/tarantool/go-tarantool"
	"golang.org/x/sync/errgroup"

	"street/conf"
	"street/model"
	"street/model/xMailer"
)

// create dockertest instance

var testTt *Tt.Adapter
var testCh *Ch.Adapter
var testMailer xMailer.Mailer

func TestMain(m *testing.M) {
	// setup dockertest instance
	dockerPool := D.InitDockerTest("")
	defer dockerPool.Cleanup()

	eg := errgroup.Group{}

	// attach tarantool
	eg.Go(func() error {
		tdt := &Tt.TtDockerTest{
			User:     "testT",
			Password: "passT",
		}
		img := tdt.ImageVersion(dockerPool, ``)
		dockerPool.Spawn(img, func(res *dockertest.Resource) error {
			t, err := tdt.ConnectCheck(res)
			if err != nil {
				return err
			}
			testTt = &Tt.Adapter{
				Connection: t,
				Reconnect: func() *tarantool.Connection {
					t, err := tdt.ConnectCheck(res)
					L.IsError(err, `tdt.ConnectCheck`)
					return t
				},
			}
			return nil
		})
		return nil
	})

	// attach clickhouse
	eg.Go(func() error {
		cdt := &Ch.ChDockerTest{
			User:     "testC",
			Password: "passC",
			Database: "default",
		}
		img := cdt.ImageLatest(dockerPool)
		dockerPool.Spawn(img, func(res *dockertest.Resource) error {
			c, err := cdt.ConnectCheck(res)
			if err != nil {
				return err
			}
			testCh = &Ch.Adapter{
				DB: c,
				Reconnect: func() *sql.DB {
					c, err := cdt.ConnectCheck(res)
					L.IsError(err, `cdt.ConnectCheck`)
					return c
				},
			}
			return nil
		})
		return nil
	})

	// mailer
	eg.Go(func() error {
		mailhogConf := conf.MailhogConf{
			MailhogHost: `localhost`,
			MailhogPort: 1025,
		}
		mailhogPort := fmt.Sprint(mailhogConf.MailhogPort)
		dockerPool.Spawn(&dockertest.RunOptions{
			Name:       `dockertest-mailhog-` + dockerPool.Uniq,
			Repository: "mailhog/mailhog",
			Tag:        `latest`,
			NetworkID:  dockerPool.Network.ID,
		}, func(res *dockertest.Resource) error {
			_, err := net.Dial("tcp", res.GetHostPort(mailhogPort+"/tcp"))
			if err != nil {
				return err
			}
			mailHog, err := xMailer.NewMailhog(mailhogConf)
			L.PanicIf(err, `xMailer.NewMailhog`)
			testMailer.SendMailFunc = mailHog.SendEmail
			return nil
		})
		return nil
	})

	_ = eg.Wait()

	// run migration
	model.RunMigration(testTt, testCh, testTt, testCh)

	// run tests
	m.Run()

	// teardown dockertest instance
}

func testDomain() *Domain {
	return &Domain{
		AuthOltp: testTt,
		AuthOlap: testCh,
		Mailer:   testMailer,
	}
}
