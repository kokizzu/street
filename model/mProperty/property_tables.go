package mProperty

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TablePropertyUs Tt.TableName = `propertyUS`

	Version = `version`

	// Basic Info
	Street                  = `street`
	City                    = `city`
	State                   = `state`
	Zip                     = `zip`
	CountryCode             = `countryCode`
	YearBuilt               = `yearBuilt`
	YearRenovated           = `yearRenovated`
	TotalSqft               = `totalSqft`
	PropertyLastUpdatedDate = `propertyLastUpdatedDate`

	TaxInfo        = `taxInfo`
	HistoryTaxInfo = `historyTaxInfo`

	// County
	CountyUrl      = `countyUrl`
	CountyName     = `countyName`
	CountyIsActive = `countyIsActive`

	// Tax notes
	TaxNote = `taxNote`

	// Amenity
	AmenitySuperGroupsJson = `amenitySuperGroups`

	// Source photo
	MediaSource     = `mediaSource`
	MediaSourceJson = `mediaSourceJson`

	// Zone data info
	ZoneDataInfoJson = `zoneDataInfo` // String - store full json string

	// Agent & Broker (Disclaimed info)
	MlsDisclaimerInfoJson = `mlsDisclaimerInfo`

	// Facilitates Nearby
	FacilityInfoJson = `facilityInfo`
	RiskInfoJson     = `riskInfo`
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

	PropertyKey            = `propertyKey` // refer to UniqPropKey?
	TransactionKey         = `transactionKey`
	TransactionType        = `transactionType`
	TransactionSign        = `transactionSign`
	TransactionTime        = `transactionTime`
	TransactionDescription = `transactionDescription`
	TransactionDateNormal  = `transactionDateNormal`
	TransactionNumber      = `transactionNumber`
	PriceNTD               = `priceNtd`
	PricePerUnit           = `pricePerUnit`
	Price                  = `price`

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

func buildStandardPropertySchema() []Tt.Field {
	schema := map[int]Tt.Field{
		0: {Id, Tt.Unsigned},

		// Metadata key info
		29: {Version, Tt.Double}, // Version of data (specific for US)
		1:  {UniqPropKey, Tt.String},
		2:  {SerialNumber, Tt.String},

		// Basic Info
		44: {Street, Tt.String}, // New field based on US data
		45: {City, Tt.String},   // New field based on US data
		46: {State, Tt.String},  // New field based on US data
		47: {Zip, Tt.String},    // New field based on US data
		9:  {Address, Tt.String},
		18: {FormattedAddress, Tt.String},

		// County (District)
		10: {District, Tt.String},
		12: {Coord, Tt.Array},
		41: {CountyUrl, Tt.String},
		42: {CountyName, Tt.String},
		43: {CountyIsActive, Tt.Boolean},
		48: {CountryCode, Tt.String}, // New field based on US data

		// Property size area
		3:  {SizeM2, Tt.String},
		39: {TotalSqft, Tt.Double}, // Specific for US data

		// Property description
		4:  {MainUse, Tt.String},
		5:  {MainBuildingMaterial, Tt.String},
		6:  {ConstructCompletedDate, Tt.String},
		7:  {NumberOfFloors, Tt.String},
		8:  {BuildingLamination, Tt.String},
		22: {Purpose, Tt.String},
		23: {HouseType, Tt.String},
		24: {Images, Tt.Array},
		25: {Bedroom, Tt.Integer},
		26: {Bathroom, Tt.Integer},
		27: {AgencyFeePercent, Tt.Double},
		28: {FloorList, Tt.Array},

		// Media images
		30: {MediaSourceJson, Tt.String},

		// Other data
		31: {ZoneDataInfoJson, Tt.String}, // Specific for US data
		32: {TaxInfo, Tt.String},          // Tax info
		33: {HistoryTaxInfo, Tt.String},   // Full history tax of tax info
		37: {YearBuilt, Tt.Integer},
		38: {YearRenovated, Tt.Integer},
		40: {AmenitySuperGroupsJson, Tt.String}, // Amenity super groups and health condition/quality of equipments and property in general
		49: {PropertyLastUpdatedDate, Tt.Integer},
		50: {TaxNote, Tt.String},

		11: {Note, Tt.String},

		// Price related to property
		19: {LastPrice, Tt.String},
		20: {PriceHistoriesSell, Tt.Array},
		21: {PriceHistoriesRent, Tt.Array},

		// Agent & Broker for Realtor
		34: {MlsDisclaimerInfoJson, Tt.String},
		35: {FacilityInfoJson, Tt.String},
		36: {RiskInfoJson, Tt.String},

		13: {CreatedAt, Tt.Integer},
		14: {CreatedBy, Tt.Unsigned},
		15: {UpdatedAt, Tt.Integer},
		16: {UpdatedBy, Tt.Unsigned},
		17: {DeletedAt, Tt.Integer},
	}

	listFieldsInPropScheme := make([]Tt.Field, len(schema))
	for i := 0; i < len(listFieldsInPropScheme); i++ {
		if schema[i].Name == "" && schema[i].Type == "" {
			continue
		} else {
			listFieldsInPropScheme[i] = schema[i]
		}
	}

	return listFieldsInPropScheme
}

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableProperty: {
		Fields:          buildStandardPropertySchema(),
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
			{TransactionDescription, Tt.String},
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
	TablePropertyUs: {
		Fields:          buildStandardPropertySchema(),
		AutoIncrementId: true,
		Unique1:         UniqPropKey,
		Indexes:         []string{SerialNumber},
		Engine:          Tt.Memtx,
		Spatial:         Coord,
	},
}

const (
	TablePropertyLogs Ch.TableName = `propertyLogs`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{}
