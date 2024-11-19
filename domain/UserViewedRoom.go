package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserViewedRoom.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserViewedRoom.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserViewedRoom.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserViewedRoom.go
//go:generate farify doublequote --file UserViewedRoom.go

type (
	UserViewedRoomIn struct {
		RequestCommon
	}

	UserViewedRoomOut struct {
		ResponseCommon
	}
)

const (
	UserViewedRoomAction = `user/viewedRoom`
)

func (d *Domain) UserViewedRoom(in *UserViewedRoomIn) (out UserViewedRoomOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	return
}
