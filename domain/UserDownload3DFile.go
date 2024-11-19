package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserDownload3DFile.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserDownload3DFile.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserDownload3DFile.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserDownload3DFile.go
//go:generate farify doublequote --file UserDownload3DFile.go

type (
	UserDownload3DFileIn struct {
		RequestCommon
	}

	UserDownload3DFileOut struct {
		ResponseCommon
	}
)

const (
	UserDownload3DFileAction = `user/download3DFile`
)

func (d *Domain) UserDownload3DFile(in *UserDownload3DFileIn) (out UserDownload3DFileOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	return
}
