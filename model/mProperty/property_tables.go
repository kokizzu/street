package mProperty

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TableProperty Tt.TableName = `property`

	Id                     = `id`
	UniqPropKey            = `UniqPropKey`
	SerialNumber           = `serialNumber`
	SizeM2                 = `sizeM2`
	MainUse                = `mainUse`
	MainBuildingMaterial   = `mainBuildingMaterial`
	ConstructCompletedDate = `constructCompletedDate`
	NumberOfFloors         = `numberOfFloors`
	BuildingLamination     = `buildingLamination`
	Address                = `address`
	District               = `district`
	Note                   = `note`
	// TODO: change to spatial column
	Lat       = `lat`
	Long      = `long`
	CreatedAt = `createdAt`
	CreatedBy = `createdBy`
	UpdatedAt = `updatedAt`
	UpdatedBy = `updatedBy`
	DeletedAt = `deletedAt`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableProperty: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{UniqPropKey, Tt.String},
			{SerialNumber, Tt.String},
			{SizeM2, Tt.String},
			{MainUse, Tt.String},
			{MainBuildingMaterial, Tt.String},
			{ConstructCompletedDate, Tt.String},
			{NumberOfFloors, Tt.String},
			{BuildingLamination, Tt.String},
			{Address, Tt.String},
			{District, Tt.String},
			{Note, Tt.String},
			{Lat, Tt.String},
			{Long, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		AutoIncrementId: true,
		Unique1:         UniqPropKey,
		Indexes:         []string{SerialNumber},
		Engine:          Tt.Memtx,
	},
}

const (
	TablePropertyLogs Ch.TableName = `tablePropertyLogs`

	RequestId = `requestId`
	Error     = `error`
	ActorId   = `actorId`
	IpAddr4   = `ipAddr4`
	IpAddr6   = `ipAddr6`
	UserAgent = `userAgent`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TablePropertyLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{RequestId, Ch.String},
			{Error, Ch.String},
			{ActorId, Ch.UInt64},
			{IpAddr4, Ch.IPv4},
			{IpAddr6, Ch.IPv6},
			{UserAgent, Ch.String},
		},
		Orders: []string{ActorId, RequestId},
	},
}
