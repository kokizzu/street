package conf

import (
	"fmt"
	"os"

	"github.com/kokizzu/gotro/S"
)

type WebConf struct {
	Port int
}

func EnvWebConf() WebConf {
	return WebConf{
		Port: S.ToInt(os.Getenv("WEB_PORT")),
	}
}

func (w WebConf) ListenAddr() string {
	return ":" + fmt.Sprint(w.Port)
}
