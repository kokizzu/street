package domain

import (
	"fmt"
	"io"
	"os"
	"street/model/mProperty/rqProperty"
	"street/model/mStorage/wcStorage"

	"github.com/gabriel-vasile/mimetype"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserUploadFile.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserUploadFile.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserUploadFile.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserUploadFile.go
//go:generate farify doublequote --file UserUploadFile.go

type (
	UserUpload3DFileIn struct {
		RequestCommon
		PropertyId uint64 `json:"propertyId" form:"propertyId" query:"propertyId" long:"propertyId" msg:"propertyId"`
		Country    string `json:"country" form:"country" query:"country" long:"country" msg:"country"`
	}

	UserUpload3DFileOut struct {
		ResponseCommon
		ImageURL string `json:"imageUrl" form:"imageUrl" query:"imageUrl" long:"imageUrl" msg:"imageUrl"`
	}
)

const (
	UserUpload3DFileAction = `user/upload3DFile`

	ErrUserUpload3DFileInvalidExtension = `invalid 3D file extension (must be .obj/.fbx/.stl/.amf/.iges)`
	ErrUserUpload3DPropertyNotFound     = `cannot find property`
	ErrUserUpload3DPropertyTWNotFound   = `cannot find property (TW)`
	ErrUserUpload3DPropertyUSNotFound   = `cannot find property (US)`
	ErrUserUpload3DFileNoRawfile        = `no rawfile in multipart`
	ErrUserUpload3DFileTooSmall         = `3D file too small`
	ErrUserUpload3DFileUnreadable       = `3D file unreadable`
	ErrUserUpload3DFileUnsupportedMime  = `unsupported mime, not obj/fbx/stl/amf/iges`
	ErrUserUpload3DFileUnableToSeek     = `unable to seek content`
)

func (d *Domain) UserUpload3DFile(in *UserUpload3DFileIn) (out UserUpload3DFileOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.Country {
	case `TW`:
		propertyTw := rqProperty.NewPropertyTW(d.AuthOltp)
		propertyTw.Id = in.PropertyId
		if !propertyTw.FindById() {
			out.SetError(400, ErrUserUpload3DPropertyTWNotFound)
			return
		}
	case `US`:
		propertyUs := rqProperty.NewPropertyUS(d.AuthOltp)
		propertyUs.Id = in.PropertyId
		if !propertyUs.FindById() {
			out.SetError(400, ErrUserUpload3DPropertyUSNotFound)
			return
		}
	default:
		property := rqProperty.NewProperty(d.AuthOltp)
		property.Id = in.PropertyId
		if !property.FindById() {
			out.SetError(400, ErrUserUpload3DPropertyNotFound)
			return
		}
	}

	if in.RawFile == nil {
		out.SetError(400, ErrUserUpload3DFileNoRawfile)
		return
	}

	if in.RawFile.Size <= 0 {
		out.SetError(400, ErrUserUpload3DFileTooSmall)
		return
	}

	reader, err := in.RawFile.openFunc()
	if err != nil {
		out.SetError(500, ErrUserUpload3DFileUnreadable)
		return
	}
	defer reader.Close()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		out.SetError(400, ErrUserUpload3DFileUnsupportedMime)
		return
	}
	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		out.SetError(400, ErrUserUpload3DFileUnableToSeek)
		return
	}

	ext := mime.Extension()
	L.Print(`File Extension:`, ext)
	switch ext {
	case `.obj`, `.fbx`, `.stl`, `.amf`, `.iges`, `.glb`:
		break
	default:
		d.Log.Warn().Msgf(`uploaded mime: %v %v`, mime.Extension(), mime.String())
		out.SetError(400, ErrUserUpload3DFileInvalidExtension)
		return
	}

	filePath := in.RequestId + `_` + in.RawFile.FileName

	file := wcStorage.NewDesignFilesMutator(d.StorOltp)
	file.SetCountryPropId(fmt.Sprintf("%s:%d", in.Country, in.PropertyId))
	file.SetCreatedAt(in.UnixNow())
	file.SetCreatedBy(sess.UserId)
	if !file.DoInsert() {
		out.SetError(500, ErrUserUploadFailedSaveMetadata)
		return
	}

	suffix := S.ValidateFilename(filePath)
	suffix = S.LeftOfLast(suffix, `.`)
	suffix = S.Right(suffix, 80)

	file.SetFilePath(fmt.Sprintf(`%d_%s%s`, file.Id, suffix, ext))

	writer, err := os.Create(d.UploadDir + file.FilePath)
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadFileCreationFailed)
		out.SetError(500, ErrUserUploadFileCreationFailed)
		return
	}

	_, err = io.Copy(writer, reader)
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadCopyFromSourceFailed)
		out.SetError(500, ErrUserUploadCopyFromSourceFailed)
		return
	}
	err = writer.Close()
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadFailedFinalize)
		out.SetError(500, ErrUserUploadFailedFinalize)
	}

	return
}
