package mProperty

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TableProperty Tt.TableName = `property`

	Id                     = `id`
	SerialNumber           = `serialNumber`
	SizeM2                 = `sizeM2`
	MainUse                = `mainUse`
	MainBuildingMaterial   = `mainBuildingMaterial`
	ConstructCompletedDate = `constructCompletedDate`
	NumberOfFloors         = `numberOfFloors`
	BuildingLamination     = `buildingLamination`
	Note                   = `note`
	CreatedAt              = `createdAt`
	CreatedBy              = `createdBy`
	UpdatedAt              = `updatedAt`
	UpdatedBy              = `updatedBy`
	DeletedAt              = `deletedAt`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableProperty: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{SerialNumber, Tt.String},
			{SizeM2, Tt.Double},
			{MainUse, Tt.String},
			{MainBuildingMaterial, Tt.String},
			{ConstructCompletedDate, Tt.Integer},
			{NumberOfFloors, Tt.Double}, // float?
			{BuildingLamination, Tt.String},
			{Note, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		AutoIncrementId: true,
		Unique1:         Id,
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
