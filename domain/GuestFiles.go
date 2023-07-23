package domain

import (
	"io"
	"os"

	"github.com/jxskiss/base62"

	"street/model/mStorage/wcStorage"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file GuestFiles.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type GuestFiles.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type GuestFiles.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type GuestFiles.go
//go:generate farify doublequote --file GuestFiles.go

type (
	GuestFilesIn struct {
		RequestCommon

		Base62id string `json:"base62Id,string" form:"base62Id" query:"base62Id" long:"base62Id" msg:"base62Id"`
		Modifier string `json:"modifier" form:"modifier" query:"modifier" long:"modifier" msg:"modifier"`
		Ext      string `json:"ext" form:"ext" query:"ext" long:"ext" msg:"ext"`
	}
	GuestFilesOut struct {
		ResponseCommon
		Request RequestCommon `json:"request" form:"request" query:"request" long:"request" msg:"request"`

		Raw         []byte `json:"raw" form:"raw" query:"raw" long:"raw" msg:"raw"`
		ContentType string `json:"contentType" form:"contentType" query:"contentType" long:"contentType" msg:"contentType"`
	}
)

const (
	GuestFilesAction = `guest/files`

	ErrGuestFilesInvalidId    = `invalid guest files id`
	ErrGuestFilesNotInDb      = `file not in database`
	ErrGuestFilesInaccessible = `file inaccessible`
	ErrGuestFilesUnreadable   = `file unreadable`
)

func (d *Domain) GuestFiles(in *GuestFilesIn) (out GuestFilesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.Request = in.RequestCommon

	id, err := base62.ParseUint([]byte(in.Base62id))
	if err != nil {
		out.SetError(400, ErrGuestFilesInvalidId)
	}

	file := wcStorage.NewFilesMutator(d.StorOltp)
	file.Id = id
	if !file.FindById() {
		out.SetError(404, ErrGuestFilesNotInDb)
		return
	}

	fsPath := file.OriginalPath
	if in.Modifier == `small` {
		fsPath = file.ResizedPath
	}

	fs, err := os.Open(d.UploadDir + fsPath)
	if err != nil {
		d.Log.Error().Msgf(`file cannot be opend: %v`, fsPath)
		out.SetError(404, ErrGuestFilesInaccessible)
		return
	}
	defer fs.Close()

	out.Raw, err = io.ReadAll(fs)
	if err != nil {
		d.Log.Error().Msgf(`file cannot be read: %v`, fsPath)
		out.SetError(500, ErrGuestFilesUnreadable)
		return
	}

	go file.DoIncStat(in.UnixNow())

	out.ContentType = file.Mime

	return
}
