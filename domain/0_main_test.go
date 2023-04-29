package domain

import (
	"database/sql"
	"testing"

	"github.com/kokizzu/gotro/D"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/ory/dockertest/v3"
	"github.com/tarantool/go-tarantool"

	"street/model"
)

// create dockertest instance

var testTt *Tt.Adapter
var testCh *Ch.Adapter

func TestMain(m *testing.M) {
	// setup dockertest instance
	dockerPool := D.InitDockerTest("")
	defer dockerPool.Cleanup()

	// attach tarantool
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

	// attach clickhouse
	cdt := &Ch.ChDockerTest{
		User:     "testC",
		Password: "passC",
		Database: "default",
	}
	img = cdt.ImageLatest(dockerPool)
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

	// run migration
	model.RunMigration(testTt, testCh)

	// run tests
	m.Run()

	// teardown dockertest instance
}
