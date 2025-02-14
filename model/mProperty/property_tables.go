package mProperty

import (
	"sort"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	PropertyPurposeSale string = `for sale`
	PropertyPurposeRent string = `for rent`
)

const (
	TablePropertyUS      Tt.TableName = `propertyUS`
	TablePropertyExtraUS Tt.TableName = `propertyExtraUS`

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
	MediaSourceJson = `mediaSourceJson`

	// Zone data info
	ZoneDataInfoJson = `zoneDataInfo` // String - store full json string

	// Agent & Broker (Disclaimed info)
	MlsDisclaimerInfoJson = `mlsDisclaimerInfo`

	// Facilitates Nearby
	FacilityInfoJson = `facilityInfo`
	RiskInfoJson     = `riskInfo`

	// Approval State
	ApprovalState        = `approvalState` // empty = approved, otherwise it's the reason of rejection
	ApprovalStatePending = `pending`
)

const (
	TablePropertyTW Tt.TableName = `propertyTW`
)

const (
	TableProperty Tt.TableName = `property`

	Id                     = `id`
	UniqPropKey            = `uniqPropKey`
	SerialNumber           = `serialNumber`
	SizeM2                 = `sizeM2`  // used also for rent/sell // TODO: deprecate, use totalSqft
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
	Altitude               = `altitude`
	CreatedAt              = `createdAt`
	CreatedBy              = `createdBy`
	UpdatedAt              = `updatedAt`
	UpdatedBy              = `updatedBy`
	DeletedAt              = `deletedAt`
	About                  = `about`
	ContactEmail           = `contactEmail`
	ContactPhone           = `contactPhone`
	Attribute              = `attribute`

	TablePropertyHistory   Tt.TableName = `property_history`
	TablePropertyHistoryUS Tt.TableName = `property_historyUS`

	PropertyKey            = `propertyKey` // refer to UniqPropKey?
	TransactionKey         = `transactionKey`
	TransactionType        = `transactionType`
	TransactionSign        = `transactionSign`
	TransactionTime        = `transactionTime`
	TransactionDescription = `transactionDescription`
	TransactionDateNormal  = `transactionDateNormal`
	TransactionNumber      = `transactionNumber`

	PriceNTD           = `priceNtd` // TODO: deprecate
	PricePerUnit       = `pricePerUnit`
	Price              = `price`
	LastPrice          = `lastPrice`
	DepositFee         = `depositFee`
	OtherFees          = `otherFees` // array of array ['label', fee]
	PriceHistoriesSell = `priceHistoriesSell`
	PriceHistoriesRent = `priceHistoriesRent`

	// mostly this one for rent/sell
	Purpose             = `purpose`     // rent/sell (empty means not for rent/sell)
	HouseType           = `houseType`   // apartment/house
	Images              = `images`      // array of string url
	ImageLabels         = `imageLabels` // array of string
	Bedroom             = `bedroom`
	Bathroom            = `bathroom`
	Livingroom          = `livingroom`
	Parking             = `parking`
	AgencyFeePercent    = `agencyFeePercent`
	MinimumDurationYear = `minimumDurationYear` // only for rent?
	FloorList           = `floorList`           // TODO: deprecate

	TableUserPropLikes Tt.TableName = `userPropLikes`

	PropId = `propId`
	UserId = `userId`

	TablePropLikeCount Tt.TableName = `propLikeCount`

	Count = `count`
)

func buildPropertyExtraUs() []Tt.Field {

	schema := map[int]Tt.Field{
		0: {Id, Tt.Unsigned},

		// Property ID
		1: {PropertyKey, Tt.String},

		// County
		2: {CountyUrl, Tt.String},
		3: {CountyIsActive, Tt.Boolean},

		// Others data
		4: {ZoneDataInfoJson, Tt.String}, // Specific for US data,
		5: {TaxInfo, Tt.String},
		6: {HistoryTaxInfo, Tt.String},
		7: {AmenitySuperGroupsJson, Tt.String},
		8: {MlsDisclaimerInfoJson, Tt.String},

		9:  {FacilityInfoJson, Tt.String},
		10: {RiskInfoJson, Tt.String},

		11: {MediaSourceJson, Tt.String},
		12: {TaxNote, Tt.String},
	}

	listFields := make([]Tt.Field, len(schema))
	for i := 0; i < len(listFields); i++ {
		if schema[i].Name == "" && schema[i].Type == "" {
			continue
		} else {
			listFields[i] = schema[i]
		}
	}

	return listFields
}

func buildPropertyHistorySchema() []Tt.Field {

	schema := map[int]Tt.Field{
		0:  {Id, Tt.Unsigned},
		1:  {PropertyKey, Tt.String},
		2:  {TransactionKey, Tt.String},
		3:  {TransactionType, Tt.String},
		4:  {TransactionSign, Tt.String},
		5:  {TransactionTime, Tt.String},
		6:  {TransactionDateNormal, Tt.String},
		7:  {TransactionNumber, Tt.String},
		8:  {PriceNTD, Tt.Integer},
		9:  {PricePerUnit, Tt.Integer},
		10: {Price, Tt.Integer},
		11: {Address, Tt.String},
		12: {District, Tt.String},
		13: {Note, Tt.String},
		14: {CreatedAt, Tt.Integer},
		15: {CreatedBy, Tt.Unsigned},
		16: {UpdatedAt, Tt.Integer},
		17: {UpdatedBy, Tt.Unsigned},
		18: {DeletedAt, Tt.Integer},
		19: {SerialNumber, Tt.String},
		20: {TransactionDescription, Tt.String},
	}

	listFields := make([]Tt.Field, len(schema))
	for i := 0; i < len(listFields); i++ {
		if schema[i].Name == "" && schema[i].Type == "" {
			continue
		} else {
			listFields[i] = schema[i]
		}
	}

	return listFields
}

func buildStandardPropertySchema() []Tt.Field {
	schema := map[int]Tt.Field{
		0: {Id, Tt.Unsigned},

		// Metadata key info
		29: {Version, Tt.String}, // Version of data (specific for US)
		1:  {UniqPropKey, Tt.String},
		2:  {SerialNumber, Tt.String},

		// Basic Info
		34: {Street, Tt.String}, // Address field based on US data
		35: {City, Tt.String},   // New field based on US data
		36: {State, Tt.String},  // New field based on US data
		37: {Zip, Tt.String},    // New field based on US data
		9:  {Address, Tt.String},
		18: {FormattedAddress, Tt.String},

		// County (District)
		10: {District, Tt.String},
		12: {Coord, Tt.Array},
		42: {Altitude, Tt.Double},
		33: {CountyName, Tt.String},
		40: {CountryCode, Tt.String}, // temporary until all records splitted into per country table

		// Property size area
		3:  {SizeM2, Tt.String},
		32: {TotalSqft, Tt.Double}, // Specific for US data

		// Property description
		4:  {MainUse, Tt.String},
		5:  {MainBuildingMaterial, Tt.String},
		6:  {ConstructCompletedDate, Tt.String},
		7:  {NumberOfFloors, Tt.String},
		8:  {BuildingLamination, Tt.String},
		22: {Purpose, Tt.String},
		23: {HouseType, Tt.String},
		24: {Images, Tt.Array},
		47: {ImageLabels, Tt.Array},

		25: {Bedroom, Tt.Integer},
		26: {Bathroom, Tt.Integer},
		41: {Livingroom, Tt.Integer},
		43: {Parking, Tt.Double},
		27: {AgencyFeePercent, Tt.Double},
		28: {FloorList, Tt.Array},

		30: {YearBuilt, Tt.Integer},
		31: {YearRenovated, Tt.Integer},
		38: {PropertyLastUpdatedDate, Tt.Integer},

		11: {Note, Tt.String},

		// Price related to property
		19: {LastPrice, Tt.String},
		44: {DepositFee, Tt.Double},
		46: {OtherFees, Tt.Array},
		45: {MinimumDurationYear, Tt.Double},
		20: {PriceHistoriesSell, Tt.Array},
		21: {PriceHistoriesRent, Tt.Array},

		13: {CreatedAt, Tt.Integer},
		14: {CreatedBy, Tt.Unsigned},
		15: {UpdatedAt, Tt.Integer},
		16: {UpdatedBy, Tt.Unsigned},
		17: {DeletedAt, Tt.Integer},

		// Approval state
		39: {ApprovalState, Tt.String},

		// Property Attribute
		48: {Attribute, Tt.String},
	}

	keys := make([]int, 0, len(schema))
	for k := range schema {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	listPropertyFields := make([]Tt.Field, len(schema))

	for i := 0; i < len(keys); i++ {
		listPropertyFields[i] = schema[keys[i]]
	}

	return listPropertyFields
}

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableProperty: {
		Fields:          buildStandardPropertySchema(),
		AutoIncrementId: true,
		Unique1:         UniqPropKey,
		Indexes:         []string{SerialNumber, ApprovalState},
		Engine:          Tt.Memtx,
		Spatial:         Coord,
	},
	TablePropertyHistory: {
		Fields:          buildPropertyHistorySchema(),
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
	TablePropertyUS: {
		Fields:          buildStandardPropertySchema(),
		AutoIncrementId: true,
		Unique1:         UniqPropKey,
		Indexes:         []string{SerialNumber},
		Engine:          Tt.Memtx,
		Spatial:         Coord,
	},
	TablePropertyExtraUS: {
		Fields:          buildPropertyExtraUs(),
		AutoIncrementId: true,
		Unique1:         PropertyKey,
		Engine:          Tt.Vinyl,
	},
	TablePropertyHistoryUS: {
		Fields:          buildPropertyHistorySchema(),
		AutoIncrementId: true,
		Unique1:         TransactionKey,
		Indexes:         []string{PropertyKey},
		Engine:          Tt.Memtx,
	},
	TablePropertyTW: {
		Fields:          buildStandardPropertySchema(),
		AutoIncrementId: true,
		Unique1:         UniqPropKey,
		Indexes:         []string{SerialNumber},
		Engine:          Tt.Memtx,
		Spatial:         Coord,
	},
}

const (
	TableViewedRooms Ch.TableName = `viewedRooms`

	ActorId    = `actorId`
	PropertyId = `propertyId`
	RoomLabel  = `roomLabel`
	Country    = `country`
)

const (
	TableScannedAreas Ch.TableName = `scannedAreas`

	Latitude  = `latitude`
	Longitude = `longitude`
)

const (
	TableScannedProperties Ch.TableName = `scannedProperties`
)

const (
	TableGeolocation Ch.TableName = `geolocation`

	StateId     = `stateId`
	StateCode   = `stateCode`
	StateName   = `stateName`
	CountryId   = `countryId`
	CountryName = `countryName`
	WikiDataId  = `wikiDataId`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TableViewedRooms: {
		Fields: []Ch.Field{
			{ActorId, Ch.UInt64},
			{CreatedAt, Ch.DateTime},
			{PropertyId, Ch.UInt64},
			{RoomLabel, Ch.String},
			{Country, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableScannedAreas: {
		Fields: []Ch.Field{
			{ActorId, Ch.UInt64},
			{CreatedAt, Ch.DateTime},
			{Latitude, Ch.Float64},
			{Longitude, Ch.Float64},
			{City, Ch.String},
			{State, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableScannedProperties: {
		Fields: []Ch.Field{
			{ActorId, Ch.UInt64},
			{CreatedAt, Ch.DateTime},
			{CountryCode, Ch.String},
			{PropertyId, Ch.UInt64},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableGeolocation: {
		Fields: []Ch.Field{
			{Id, Ch.UInt64},
			{City, Ch.String},
			{StateId, Ch.UInt64},
			{StateCode, Ch.String},
			{StateName, Ch.String},
			{CountryId, Ch.UInt64},
			{CountryCode, Ch.String},
			{CountryName, Ch.String},
			{Latitude, Ch.Float64},
			{Longitude, Ch.Float64},
			{WikiDataId, Ch.String},
		},
		Orders: []string{Id},
	},
}
