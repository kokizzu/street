package mProperty

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TablePropertyUsa Tt.TableName = `property_usa`

	Version    = `version`
	PropertyId = `propertyId` // ID of Property ID in Redfin

	// Basic Info
	Street                  = `street`
	City                    = `city`
	State                   = `state`
	Zip                     = `zip`
	CountryCode             = `countryCode`
	PropertyTypeName        = `propertyTypeName`
	YearBuilt               = `yearBuilt`
	YearRenovated           = `yearRenovated`
	TotalSqft               = `totalSqft`
	Apn                     = `apn` // APN stands for Assessor’s Parcel Number. It is a unique number assigned by the Tax Assessor to a piece of property or each tract of land within a county. It’s a numbering system that distinguishes one property from the next.
	PropertyLastUpdatedDate = `propertyLastUpdatedDate`
	DisplayTimeZone         = `displayTimeZone`

	// Latest tax info
	TaxableLandValue        = `taxableLandValue`
	TaxableImprovementValue = `taxableImprovementValue`
	RollYear                = `rollYear`
	TaxesDue                = `taxesDue`

	// County
	CountyUrl      = `countyUrl`
	CountyName     = `countyName`
	CountyIsActive = `countyIsActive`

	// Tax notes
	TaxNote = `taxNote`

	// Amenity
	AmenitySuperGroups = `amenitySuperGroups`

	// Source photo
	MediaSource = `mediaSource`

	// Zone data info
	ZoneName            = `zoneName`            // Ex: Residential Apartment
	ZoneType            = `zoneType`            // Ex: Residential
	ZoneSubType         = `zoneSubType`         // Ex: Single Family
	ZoneDisplay         = `zoneDisplay`         // Ex: Residential Single Family
	ZoneCode            = `zoneCode`            // Ex: RA-2
	PermittedLandUse    = `permittedLandUse`    // Ex: ["Single-Family","Multi-Family","Short-Term Rentals","Commercial"]
	NotPermittedLandUse = `notPermittedLandUse` // Ex: ["Two-Family", "ADU", "Industrial"]

	// Agent & Broker
	ListingBrokerName   = `listingBrokerName`
	ListingBrokerNumber = `listingBrokerNumber`
	ListingAgentName    = `listingAgentName`
	ListingAgentNumber  = `listingAgentNumber`
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

	TableUserPropLikes Tt.TableName = `userPropLikes`

	PropId = `propId`
	UserId = `userId`

	TablePropLikeCount Tt.TableName = `propLikeCount`

	Count = `count`
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
	TableUserPropLikes: {
		Fields: []Tt.Field{
			{PropId, Tt.Unsigned},
			{UserId, Tt.Unsigned},
			{CreatedAt, Tt.Integer},
		},
		Engine:  Tt.Vinyl,
		Uniques: []string{UserId, PropId},
	},
	TablePropLikeCount: {
		Fields: []Tt.Field{
			{PropId, Tt.Unsigned},
			{Count, Tt.Integer},
		},
		Unique1: PropId,
		Engine:  Tt.Memtx,
	},
	TablePropertyUsa: {
		Fields: []Tt.Field{

			{Id, Tt.Unsigned},

			// Basic info
			{Version, Tt.Integer},
			{PropertyId, Tt.Unsigned},
			{Street, Tt.String},
			{City, Tt.String},
			{State, Tt.String},
			{Zip, Tt.String},
			{CountryCode, Tt.String},
			{PropertyTypeName, Tt.String},
			{YearBuilt, Tt.Integer},
			{YearRenovated, Tt.Integer},
			{TotalSqft, Tt.Double},
			{Apn, Tt.String},
			{PropertyLastUpdatedDate, Tt.Integer},
			{DisplayTimeZone, Tt.String},
			{Coord, Tt.Array},
			{Bedroom, Tt.Integer},
			{Bathroom, Tt.Integer},

			// Tax info
			{TaxableLandValue, Tt.Integer},
			{TaxableImprovementValue, Tt.Integer},
			{RollYear, Tt.Integer},
			{TaxesDue, Tt.Double},

			// Amenity Group
			{AmenitySuperGroups, Tt.Array},

			// County
			{CountyUrl, Tt.String},
			{CountyName, Tt.String},
			{CountyIsActive, Tt.Boolean},

			// Media sources
			{MediaSource, Tt.Array},

			// Zone data info
			{ZoneName, Tt.String},
			{ZoneType, Tt.String},
			{ZoneSubType, Tt.String},
			{ZoneDisplay, Tt.String},
			{ZoneCode, Tt.String},
			{PermittedLandUse, Tt.Array},
			{NotPermittedLandUse, Tt.Array},

			{Note, Tt.String},
			{TaxNote, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},

			// Agent & Broker
			{ListingBrokerName, Tt.String},
			{ListingBrokerNumber, Tt.String},
			{ListingAgentName, Tt.String},
			{ListingAgentNumber, Tt.String},

			{AgencyFeePercent, Tt.Double},
			{Country, Tt.String},
		},
		AutoIncrementId: true,
		Unique1:         PropertyId,
		Indexes:         []string{PropertyId},
		Engine:          Tt.Memtx,
		Spatial:         Coord,
	},
}

const (
	TablePropertyLogs Ch.TableName = `propertyLogs`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{}
