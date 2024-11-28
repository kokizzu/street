package domain

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"street/model/mProperty/rqProperty"
	"street/model/mStorage/rqStorage"
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
		PropKey string `json:"propKey" form:"propKey" query:"propKey" long:"propKey" msg:"propKey"`
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
	ErrUserUpload3DFileExist						= `3D file already exist`
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
		propertyTw.UniqPropKey = in.PropKey
		if in.PropKey == `` {
			if !propertyTw.FindById() {
				out.SetError(400, ErrUserUpload3DPropertyTWNotFound)
				return
			}
		} else {
			if !propertyTw.FindByUniqPropKey() {
				out.SetError(400, ErrUserUpload3DPropertyTWNotFound)
				return
			}
		}
	case `US`:
		propertyUs := rqProperty.NewPropertyUS(d.AuthOltp)
		propertyUs.Id = in.PropertyId
		propertyUs.UniqPropKey = in.PropKey
		if in.PropKey == `` {
			if !propertyUs.FindById() {
				out.SetError(400, ErrUserUpload3DPropertyUSNotFound)
				return
			}
		} else {
			if !propertyUs.FindByUniqPropKey() {
				out.SetError(400, ErrUserUpload3DPropertyUSNotFound)
				return
			}
		}
	default:
		property := rqProperty.NewProperty(d.AuthOltp)
		property.Id = in.PropertyId
		property.UniqPropKey = in.PropKey
		if in.PropKey == `` {
			if !property.FindById() {
				out.SetError(400, ErrUserUpload3DPropertyNotFound)
				return
			}
		} else {
			if !property.FindByUniqPropKey() {
				out.SetError(400, ErrUserUpload3DPropertyNotFound)
				return
			}
		}
	}

	countryPropId := fmt.Sprintf("%s:%d", in.Country, in.PropertyId)
	img3d := rqStorage.NewDesignFiles(d.StorOltp)
	img3d.CountryPropId = countryPropId
	if img3d.FindByCountryPropId() {
		out.SetError(400, ErrUserUpload3DFileExist)
		return
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
	switch ext {
	case `.obj`, `.fbx`, `.stl`, `.amf`, `.iges`, `.glb`:
		break
	default:
		d.Log.Warn().Msgf(`uploaded mime: %v %v`, mime.Extension(), mime.String())
		ext = filepath.Ext(in.RawFile.FileName)
		if ext == `` {
			ext = `.obj`
		}
	}

	filePath := in.RequestId + `_` + in.RawFile.FileName

	file := wcStorage.NewDesignFilesMutator(d.StorOltp)
	file.SetCountryPropId(countryPropId)
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
	pathName := d.UploadDir + file.FilePath
	writer, err := os.Create(pathName)
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadFileCreationFailed)
		out.SetError(500, ErrUserUploadFileCreationFailed)
		return
	}

	L.Print(`Pathname:`, pathName)

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

	if !file.DoUpdateById() {
		out.SetError(500, ErrUserUploadFailedSaveMetadata)
		return
	}

	out.ImageURL = file.FilePath

	return
}
