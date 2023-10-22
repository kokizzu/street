package domain

import (
	"fmt"
	"image/jpeg"
	"io"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gabriel-vasile/mimetype"
	"github.com/jxskiss/base62"
	"github.com/kokizzu/gotro/S"

	"street/model/mStorage/wcStorage"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserUploadFile.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserUploadFile.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserUploadFile.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserUploadFile.go
//go:generate farify doublequote --file UserUploadFile.go

type (
	UserUploadFileIn struct {
		RequestCommon

		// retrieve rawFile from body
		Purpose string `json:"purpose" form:"purpose" query:"purpose" long:"purpose" msg:"purpose"`
	}

	UserUploadFileOut struct {
		ResponseCommon
		ResizedUrl  string `json:"resizedUrl" form:"resizedUrl" query:"resizedUrl" long:"resizedUrl" msg:"resizedUrl"`
		OriginalUrl string `json:"originalUrl" form:"originalUrl" query:"originalUrl" long:"originalUrl" msg:"originalUrl"`
		UrlPattern  string `json:"urlPattern" form:"urlPattern" query:"urlPattern" long:"urlPattern" msg:"urlPattern"`
	}
)

const (
	UserUploadFileAction = `user/uploadFile`

	ErrUserUploadNoRawFile            = `user upload no rawFile on multipart`
	ErrUserUploadFailedSaveMetadata   = `user upload save metadata failed`
	ErrUserUploadGarbageCollect       = `user upload remove metadata failed`
	ErrUserUploadFailedUpdateMetadata = `user upload update metadata failed`
	ErrUserUploadInvalidPurpose       = `user upload purpose invalid`
	ErrUserUploadTooSmall             = `user upload too small`
	ErrUserUploadFileUnreadable       = `user upload file unreadable`
	ErrUserUploadUnabledToDetectMime  = `user upload unable detect mime`
	ErrUserUploadUnableSeekContent    = `user upload unable seek content`
	ErrUserUploadUnsupportedMime      = `user upload unsupported mime, not jpg or png`
	ErrUserUploadFileCreationFailed   = `user upload file creation failed`
	ErrUserUploadCopyFromSourceFailed = `user upload copy from source failed`
	ErrUserUploadUnableDecodeImage    = `user upload unable decode image`
	ErrUserUploadThumbnailCreation    = `user upload thumbnail creation`
	ErrUserUploadUnableEncodeImage    = `user upload unable encode image`
	ErrUserUploadFailedFinalize       = `user upload failed finalize`
	ErrUserUploadThumbnailFailed      = `user upload failed thumbnail creation`
	ErrUserUploadThumbnailUnreadable  = `user upload thumbnail unreadable`

	UserUploadFile_PurposeProperty  = `property`
	UserUploadFile_PurposeFloorPlan = `floorPlan`
)

func (d *Domain) UserUploadFile(in *UserUploadFileIn) (out UserUploadFileOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.RawFile == nil {
		out.SetError(400, ErrUserUploadNoRawFile)
		return
	}

	if in.RawFile.Size <= 0 {
		out.SetError(400, ErrUserUploadTooSmall)
		return
	}

	switch in.Purpose {
	case UserUploadFile_PurposeProperty, UserUploadFile_PurposeFloorPlan:
	//TODO: check for ownership of property before processing
	default:
		out.SetError(400, ErrUserUploadInvalidPurpose)
		return
	}

	reader, err := in.RawFile.openFunc()
	if err != nil {
		out.SetError(500, ErrUserUploadFileUnreadable)
		return
	}
	defer reader.Close()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		out.SetError(400, ErrUserUploadUnabledToDetectMime)
		return
	}
	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		out.SetError(400, ErrUserUploadUnableSeekContent)
		return
	}

	ext := mime.Extension()
	switch ext {
	case `.png`, `.jpg`:
	default:
		d.Log.Warn().Msgf(`uploaded mime: %v %v`, mime.Extension(), mime.String())
		out.SetError(400, ErrUserUploadUnsupportedMime)
		return
	}

	file := wcStorage.NewFilesMutator(d.StorOltp)
	file.Mime = in.RawFile.Mime
	file.OriginalSize = uint64(in.RawFile.Size)
	file.OriginalPath = in.RequestId + `_` + in.RawFile.FileName
	file.CreatedAt = in.UnixNow()
	file.CreatedBy = sess.UserId
	file.Purpose = in.Purpose
	if !file.DoInsert() {
		out.SetError(500, ErrUserUploadFailedSaveMetadata)
		return
	}
	out.refId = file.Id
	defer func() {
		if out.HasError() {
			if !file.DoDeletePermanentById() {
				out.SetError(500, ErrUserUploadGarbageCollect)
			}
			return
		}
		if !file.DoUpdateById() {
			out.SetError(500, ErrUserUploadFailedUpdateMetadata)
			return
		}

		b62id := string(base62.FormatInt(int64(file.Id)))
		out.UrlPattern = fmt.Sprintf(`/`+GuestFilesAction+`/%s-___%s`, b62id, ext)
		out.ResizedUrl = fmt.Sprintf(`/`+GuestFilesAction+`/%s-small%s`, b62id, ext)
		out.OriginalUrl = fmt.Sprintf(`/`+GuestFilesAction+`/%s-orig%s`, b62id, ext)
	}()
	file.SetAccessCount(1)
	file.SetLastAccessAt(in.UnixNow())

	// save uploaded file

	suffix := S.ValidateFilename(file.OriginalPath)
	suffix = S.LeftOfLast(suffix, `.`)
	suffix = S.Right(suffix, 80)
	file.SetOriginalPath(fmt.Sprintf(`%d_%s%s`, file.Id, suffix, ext))

	writer, err := os.Create(d.UploadDir + file.OriginalPath)
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

	origImg, err := imaging.Open(d.UploadDir + file.OriginalPath)
	if err != nil {
		out.SetError(500, ErrUserUploadUnableDecodeImage)
		return
	}

	// create thumbnail: 640x400

	smallImg := imaging.Fit(origImg, 640, 400, imaging.Box)
	file.SetResizedPath(fmt.Sprintf(`%d_%s-small%s`, file.Id, suffix, ext))

	writer, err = os.Create(d.UploadDir + file.ResizedPath)
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadThumbnailCreation)
		out.SetError(500, ErrUserUploadThumbnailCreation)
		return
	}

	err = jpeg.Encode(writer, smallImg, &jpeg.Options{Quality: 90})
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadUnableEncodeImage)
		out.SetError(500, ErrUserUploadUnableEncodeImage)
		return
	}

	stat, err := writer.Stat()
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadThumbnailUnreadable)
		out.SetError(500, ErrUserUploadThumbnailUnreadable)
		return
	}

	err = writer.Close()
	if err != nil {
		d.Log.Err(err).Msg(ErrUserUploadThumbnailFailed)
		out.SetError(500, ErrUserUploadThumbnailFailed)
		return
	}

	file.SetResizedSize(uint64(stat.Size()))

	return
}
