package conf

import (
	"os"
)

type GmapConf struct {
	PublicApiKey string
}

func EnvGmap() GmapConf {
	return GmapConf{
		PublicApiKey: os.Getenv("GMAP_PUBLIC_API_KEY"),
	}
}
