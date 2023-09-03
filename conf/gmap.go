package conf

import (
	"os"
)

type GmapConf struct {
	PublicApiKey  string
	PrivateApiKey string
}

func EnvGmap() GmapConf {
	return GmapConf{
		PublicApiKey:  os.Getenv("GMAP_PUBLIC_API_KEY"),
		PrivateApiKey: os.Getenv(`GOOGLE_API_KEY`),
	}
}
