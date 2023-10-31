package conf

import (
	"os"
	"path"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
)

func UploadDir() string {
	// upload directory logic
	uploadDir := os.Getenv(`UPLOAD_DIR`)
	return createDirIfNotExists(uploadDir, `upload`)
}

func CacheDir() string {
	// upload directory logic
	cacheDir := os.Getenv(`CACHE_DIR`)
	return createDirIfNotExists(cacheDir, `cache`)
}

func createDirIfNotExists(dir, what string) string {
	if !S.StartsWith(dir, `/`) {
		workdDir, err := os.Getwd()
		L.PanicIf(err, `failed get working directory for `+what)
		dir = path.Join(workdDir, dir)
	}
	dirStat, err := os.Stat(dir)
	if err != nil {
		err = os.MkdirAll(dir, 0770)
		L.PanicIf(err, `failed create `+what+` directory: `+dir)
		dirStat, _ = os.Stat(dir)
	}
	if !dirStat.IsDir() {
		panic(what + ` dir is not a directory: ` + dir)
	}
	if !S.EndsWith(dir, `/`) {
		dir += `/`
	}
	return dir
}
