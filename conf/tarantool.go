package conf

import (
	"errors"
	"fmt"
	"os"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/tarantool/go-tarantool"
)

type TarantoolConf struct {
	User string
	Pass string
	Host string
	Port int
}

func EnvTarantool() TarantoolConf {
	return TarantoolConf{
		User: os.Getenv("TARANTOOL_USER"),
		Pass: os.Getenv("TARANTOOL_PASS"),
		Host: os.Getenv("TARANTOOL_HOST"),
		Port: S.ToInt(os.Getenv("TARANTOOL_PORT")),
	}
}

var ErrConnectTarantool = errors.New(`TarantoolConf) Connect`)

func (c TarantoolConf) Connect() (a *Tt.Adapter, err error) {
	hostPort := fmt.Sprintf("%s:%d", c.Host, c.Port)
	connectFunc := func() *tarantool.Connection {
		taran, err := tarantool.Connect(hostPort, tarantool.Opts{
			User: c.User,
			Pass: c.Pass,
		})
		L.IsError(err, `tarantool.Connect `+hostPort)
		return taran
	}
	taran := connectFunc()
	_, err = taran.Ping()
	if err != nil {
		return nil, WrapError(ErrConnectTarantool, err)
	}
	a = &Tt.Adapter{
		Connection: taran,
		Reconnect:  connectFunc,
	}
	return
}
