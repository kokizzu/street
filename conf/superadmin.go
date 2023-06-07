package conf

import (
	"os"

	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
)

func EnvSuperadmins() (res M.SB) {
	emailsStr := os.Getenv(`SUPERADMIN_EMAILS`)
	emails := S.Split(emailsStr, `,`)
	for _, email := range emails {
		res[email] = true
	}
	return res
}
