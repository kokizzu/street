package mPropertyHistory

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TablePropertyHistory Tt.TableName = `property_history`

	Id                    = `id`
	PropertyKey           = `propertyKey`
	TransactionKey        = `transactionKey`
	TransactionType       = `transactionType`
	TransactionSign       = `transactionSign`
	TransactionTime       = `transactionTime`
	TransactionDateNormal = `transactionDateNormal`
	TransactionNumber     = `transactionNumber`
	PriceNTD              = `priceNtd`
	PricePerUnit          = `pricePerUnit`
	Price                 = `price`
	Address               = `address`
	District              = `district`
	Note                  = `note`
	CreatedAt             = `createdAt`
	CreatedBy             = `createdBy`
	UpdatedAt             = `updatedAt`
	UpdatedBy             = `updatedBy`
	DeletedAt             = `deletedAt`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TablePropertyHistory: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{PropertyKey, Tt.String},
			{TransactionKey, Tt.String},
			{TransactionType, Tt.String},
			{TransactionSign, Tt.String},
			{TransactionTime, Tt.String},
			{TransactionDateNormal, Tt.String},
			{TransactionNumber, Tt.String},
			{PriceNTD, Tt.Integer},
			{PricePerUnit, Tt.Integer},
			{Price, Tt.Integer},
			{Address, Tt.String},
			{District, Tt.String},
			{Note, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		AutoIncrementId: true,
		Unique1:         TransactionKey,
		Indexes:         []string{Id},
		Engine:          Tt.Memtx,
	},
}

const (
	TablePropertyHistoryLogs Ch.TableName = `tablePropertyHistoryLogs`

	RequestId = `requestId`
	Error     = `error`
	ActorId   = `actorId`
	IpAddr4   = `ipAddr4`
	IpAddr6   = `ipAddr6`
	UserAgent = `userAgent`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TablePropertyHistoryLogs: {
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
