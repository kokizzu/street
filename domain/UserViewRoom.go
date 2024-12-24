package domain

import (
	"street/model/mProperty/rqProperty"
	"street/model/mProperty/saProperty"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserViewRoom.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserViewRoom.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserViewRoom.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserViewRoom.go
//go:generate farify doublequote --file UserViewRoom.go

type (
	UserViewRoomIn struct {
		RequestCommon
		ViewedRoom saProperty.ViewedRooms `json:"viewedRoom" form:"viewedRoom" query:"viewedRoom" long:"viewedRoom" msg:"viewedRoom"`
	}

	UserViewRoomOut struct {
		ResponseCommon
	}
)

const (
	UserViewRoomAction = `user/viewRoom`

	ErrUserViewRoomPropertyNotFound = `property not found`
	ErrUserViewRoomSaveFailed       = `failed to save viewed room`
	ErrUserViewRoomSaveReturn       = `failed to return saved viewed room`
)

func (d *Domain) UserViewRoom(in *UserViewRoomIn) (out UserViewRoomOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.ViewedRoom.Country {
	case `US`:
		propUS := rqProperty.NewPropertyUS(d.PropOltp)
		propUS.Id = in.ViewedRoom.PropertyId
		if !propUS.FindById() {
			out.SetError(400, ErrUserViewRoomPropertyNotFound)
			return
		}
	case `TW`:
		propTW := rqProperty.NewPropertyTW(d.PropOltp)
		propTW.Id = in.ViewedRoom.PropertyId
		if !propTW.FindById() {
			out.SetError(400, ErrUserViewRoomPropertyNotFound)
			return
		}
	default:
		prop := rqProperty.NewProperty(d.PropOltp)
		prop.Id = in.ViewedRoom.PropertyId
		if !prop.FindById() {
			out.SetError(400, ErrUserViewRoomPropertyNotFound)
			return
		}
	}

	row := saProperty.ViewedRooms{
		ActorId:    sess.UserId,
		CreatedAt:  time.Now(),
		PropertyId: in.ViewedRoom.PropertyId,
		RoomLabel:  in.ViewedRoom.RoomLabel,
		Country:    in.ViewedRoom.Country,
	}

	if !d.viewedRoomsLogs.Insert([]any{
		row.ActorId,
		row.CreatedAt,
		row.PropertyId,
		row.RoomLabel,
		row.Country,
	}) {
		out.SetError(500, ErrUserViewRoomSaveFailed)
		return
	}

	out.StatusCode = 200

	return
}
