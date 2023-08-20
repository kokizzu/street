package mProperty

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TableProperty Tt.TableName = `property`

	Id                     = `id`
	UniqPropKey            = `uniqPropKey`
	SerialNumber           = `serialNumber`
	SizeM2                 = `sizeM2`  // used also for rent/sell
	MainUse                = `mainUse` // this is facilities for rent/sell
	MainBuildingMaterial   = `mainBuildingMaterial`
	ConstructCompletedDate = `constructCompletedDate`
	NumberOfFloors         = `numberOfFloors`
	BuildingLamination     = `buildingLamination`
	Address                = `address`
	FormattedAddress       = `formattedAddress` // used also for rent/sell
	District               = `district`
	Note                   = `note`  // used also for rent/sell
	Coord                  = `coord` // [latitude, longitude] used also for rent/sell
	CreatedAt              = `createdAt`
	CreatedBy              = `createdBy`
	UpdatedAt              = `updatedAt`
	UpdatedBy              = `updatedBy`
	DeletedAt              = `deletedAt`
	LastPrice              = `lastPrice`
	PriceHistoriesSell     = `priceHistoriesSell`
	PriceHistoriesRent     = `priceHistoriesRent`

	TablePropertyHistory Tt.TableName = `property_history`

	PropertyKey           = `propertyKey` // refer to UniqPropKey?
	TransactionKey        = `transactionKey`
	TransactionType       = `transactionType`
	TransactionSign       = `transactionSign`
	TransactionTime       = `transactionTime`
	TransactionDateNormal = `transactionDateNormal`
	TransactionNumber     = `transactionNumber`
	PriceNTD              = `priceNtd`
	PricePerUnit          = `pricePerUnit`
	Price                 = `price`

	// mostly this one for rent/sell
	Purpose          = `purpose`   // rent/sell (empty means not for rent/sell)
	HouseType        = `houseType` // apartment/house
	Images           = `images`
	Bedroom          = `bedroom`
	Bathroom         = `bathroom`
	AgencyFeePercent = `agencyFeePercent`
	FloorList        = `floorList`
	Country          = `country`
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
			{Coord, Tt.Array},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{FormattedAddress, Tt.String},
			{LastPrice, Tt.String},
			{PriceHistoriesSell, Tt.Array},
			{PriceHistoriesRent, Tt.Array},
			{Purpose, Tt.String},
			{HouseType, Tt.String},
			{Images, Tt.Array},
			{Bedroom, Tt.Integer},
			{Bathroom, Tt.Integer},
			{AgencyFeePercent, Tt.Double},
			{FloorList, Tt.Array},
			{Country, Tt.String},
		},
		AutoIncrementId: true,
		Unique1:         UniqPropKey,
		Indexes:         []string{SerialNumber},
		Engine:          Tt.Memtx,
		Spatial:         Coord,
	},
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
			{SerialNumber, Tt.String},
		},
		AutoIncrementId: true,
		Unique1:         TransactionKey,
		Indexes:         []string{PropertyKey},
		Engine:          Tt.Memtx,
	},
}

const (
	TablePropertyLogs Ch.TableName = `propertyLogs`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{}
