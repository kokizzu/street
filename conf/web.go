package conf

import (
	"fmt"
	"os"

	"github.com/kokizzu/gotro/S"
)

type WebConf struct {
	Port int
}

var WEB_PROTO_DOMAIN string

func init() {
	WEB_PROTO_DOMAIN = os.Getenv("WEB_PROTO_DOMAIN")
	if WEB_PROTO_DOMAIN == `` {
		WEB_PROTO_DOMAIN = `http://localhost`
	}
}

func EnvWebConf() WebConf {
	return WebConf{
		Port: S.ToInt(os.Getenv("WEB_PORT")),
	}
}

func (w WebConf) ListenAddr() string {
	return ":" + fmt.Sprint(w.Port)
}
