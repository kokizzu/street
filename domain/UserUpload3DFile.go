package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserUploadFile.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserUploadFile.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserUploadFile.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserUploadFile.go
//go:generate farify doublequote --file UserUploadFile.go

const (
	UserUpload3DFileAction = `user/upload3DFile`
)

type (
	UserUpload3DFileIn struct {
		RequestCommon
		PropertyId string `json:"propertyId" form:"propertyId" query:"propertyId" long:"propertyId" msg:"propertyId"`
		Country		 string `json:"country" form:"country" query:"country" long:"country" msg:"country"`
	}

	UserUpload3DFileOut struct {
		ResponseCommon
		ImageURL  string `json:"imageUrl" form:"imageUrl" query:"imageUrl" long:"imageUrl" msg:"imageUrl"`
	}
)

func (d *Domain) UserUpload3DFile(in *UserUpload3DFileIn) (out UserUpload3DFileOut) {
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

	// TODO: validate 3D file

	return
}