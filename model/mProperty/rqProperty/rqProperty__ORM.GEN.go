package rqProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mProperty"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type rqProperty__ORM.GEN.go
// go:generate msgp -tests=false -file rqProperty__ORM.GEN.go -o rqProperty__MSG.GEN.go

// Property DAO reader/query struct
type Property struct {
	Adapter                *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id                     uint64
	UniqPropKey            string
	SerialNumber           string
	SizeM2                 string
	MainUse                string
	MainBuildingMaterial   string
	ConstructCompletedDate string
	NumberOfFloors         string
	BuildingLamination     string
	Address                string
	District               string
	Note                   string
	Coord                  []any
	CreatedAt              int64
	CreatedBy              uint64
	UpdatedAt              int64
	UpdatedBy              uint64
	DeletedAt              int64
	FormattedAddress       string
	LastPrice              string
	PriceHistories         []any
}

// NewProperty create new ORM reader/query object
func NewProperty(adapter *Tt.Adapter) *Property {
	return &Property{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *Property) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableProperty) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *Property) SqlTableName() string { //nolint:dupl false positive
	return `"property"`
}

func (p *Property) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *Property) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{p.Id})
	if L.IsError(err, `Property.FindById failed: `+p.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		p.FromArray(rows[0])
		return true
	}
	return false
}

// SpatialIndexCoord return spatial index name
func (p *Property) SpatialIndexCoord() string { //nolint:dupl false positive
	return `coord`
}

// UniqueIndexUniqPropKey return unique index name
func (p *Property) UniqueIndexUniqPropKey() string { //nolint:dupl false positive
	return `uniqPropKey`
}

// FindByUniqPropKey Find one by UniqPropKey
func (p *Property) FindByUniqPropKey() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexUniqPropKey(), 0, 1, tarantool.IterEq, A.X{p.UniqPropKey})
	if L.IsError(err, `Property.FindByUniqPropKey failed: `+p.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		p.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (p *Property) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "uniqPropKey"
	, "serialNumber"
	, "sizeM2"
	, "mainUse"
	, "mainBuildingMaterial"
	, "constructCompletedDate"
	, "numberOfFloors"
	, "buildingLamination"
	, "address"
	, "district"
	, "note"
	, "coord"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "formattedAddress"
	, "lastPrice"
	, "priceHistories"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Property) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "uniqPropKey"
	, "serialNumber"
	, "sizeM2"
	, "mainUse"
	, "mainBuildingMaterial"
	, "constructCompletedDate"
	, "numberOfFloors"
	, "buildingLamination"
	, "address"
	, "district"
	, "note"
	, "coord"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "formattedAddress"
	, "lastPrice"
	, "priceHistories"
	`
}

// ToUpdateArray generate slice of update command
func (p *Property) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, p.Id},
		A.X{`=`, 1, p.UniqPropKey},
		A.X{`=`, 2, p.SerialNumber},
		A.X{`=`, 3, p.SizeM2},
		A.X{`=`, 4, p.MainUse},
		A.X{`=`, 5, p.MainBuildingMaterial},
		A.X{`=`, 6, p.ConstructCompletedDate},
		A.X{`=`, 7, p.NumberOfFloors},
		A.X{`=`, 8, p.BuildingLamination},
		A.X{`=`, 9, p.Address},
		A.X{`=`, 10, p.District},
		A.X{`=`, 11, p.Note},
		A.X{`=`, 12, p.Coord},
		A.X{`=`, 13, p.CreatedAt},
		A.X{`=`, 14, p.CreatedBy},
		A.X{`=`, 15, p.UpdatedAt},
		A.X{`=`, 16, p.UpdatedBy},
		A.X{`=`, 17, p.DeletedAt},
		A.X{`=`, 18, p.FormattedAddress},
		A.X{`=`, 19, p.LastPrice},
		A.X{`=`, 20, p.PriceHistories},
	}
}

// IdxId return name of the index
func (p *Property) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *Property) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxUniqPropKey return name of the index
func (p *Property) IdxUniqPropKey() int { //nolint:dupl false positive
	return 1
}

// SqlUniqPropKey return name of the column being indexed
func (p *Property) SqlUniqPropKey() string { //nolint:dupl false positive
	return `"uniqPropKey"`
}

// IdxSerialNumber return name of the index
func (p *Property) IdxSerialNumber() int { //nolint:dupl false positive
	return 2
}

// SqlSerialNumber return name of the column being indexed
func (p *Property) SqlSerialNumber() string { //nolint:dupl false positive
	return `"serialNumber"`
}

// IdxSizeM2 return name of the index
func (p *Property) IdxSizeM2() int { //nolint:dupl false positive
	return 3
}

// SqlSizeM2 return name of the column being indexed
func (p *Property) SqlSizeM2() string { //nolint:dupl false positive
	return `"sizeM2"`
}

// IdxMainUse return name of the index
func (p *Property) IdxMainUse() int { //nolint:dupl false positive
	return 4
}

// SqlMainUse return name of the column being indexed
func (p *Property) SqlMainUse() string { //nolint:dupl false positive
	return `"mainUse"`
}

// IdxMainBuildingMaterial return name of the index
func (p *Property) IdxMainBuildingMaterial() int { //nolint:dupl false positive
	return 5
}

// SqlMainBuildingMaterial return name of the column being indexed
func (p *Property) SqlMainBuildingMaterial() string { //nolint:dupl false positive
	return `"mainBuildingMaterial"`
}

// IdxConstructCompletedDate return name of the index
func (p *Property) IdxConstructCompletedDate() int { //nolint:dupl false positive
	return 6
}

// SqlConstructCompletedDate return name of the column being indexed
func (p *Property) SqlConstructCompletedDate() string { //nolint:dupl false positive
	return `"constructCompletedDate"`
}

// IdxNumberOfFloors return name of the index
func (p *Property) IdxNumberOfFloors() int { //nolint:dupl false positive
	return 7
}

// SqlNumberOfFloors return name of the column being indexed
func (p *Property) SqlNumberOfFloors() string { //nolint:dupl false positive
	return `"numberOfFloors"`
}

// IdxBuildingLamination return name of the index
func (p *Property) IdxBuildingLamination() int { //nolint:dupl false positive
	return 8
}

// SqlBuildingLamination return name of the column being indexed
func (p *Property) SqlBuildingLamination() string { //nolint:dupl false positive
	return `"buildingLamination"`
}

// IdxAddress return name of the index
func (p *Property) IdxAddress() int { //nolint:dupl false positive
	return 9
}

// SqlAddress return name of the column being indexed
func (p *Property) SqlAddress() string { //nolint:dupl false positive
	return `"address"`
}

// IdxDistrict return name of the index
func (p *Property) IdxDistrict() int { //nolint:dupl false positive
	return 10
}

// SqlDistrict return name of the column being indexed
func (p *Property) SqlDistrict() string { //nolint:dupl false positive
	return `"district"`
}

// IdxNote return name of the index
func (p *Property) IdxNote() int { //nolint:dupl false positive
	return 11
}

// SqlNote return name of the column being indexed
func (p *Property) SqlNote() string { //nolint:dupl false positive
	return `"note"`
}

// IdxCoord return name of the index
func (p *Property) IdxCoord() int { //nolint:dupl false positive
	return 12
}

// SqlCoord return name of the column being indexed
func (p *Property) SqlCoord() string { //nolint:dupl false positive
	return `"coord"`
}

// IdxCreatedAt return name of the index
func (p *Property) IdxCreatedAt() int { //nolint:dupl false positive
	return 13
}

// SqlCreatedAt return name of the column being indexed
func (p *Property) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *Property) IdxCreatedBy() int { //nolint:dupl false positive
	return 14
}

// SqlCreatedBy return name of the column being indexed
func (p *Property) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *Property) IdxUpdatedAt() int { //nolint:dupl false positive
	return 15
}

// SqlUpdatedAt return name of the column being indexed
func (p *Property) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *Property) IdxUpdatedBy() int { //nolint:dupl false positive
	return 16
}

// SqlUpdatedBy return name of the column being indexed
func (p *Property) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *Property) IdxDeletedAt() int { //nolint:dupl false positive
	return 17
}

// SqlDeletedAt return name of the column being indexed
func (p *Property) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxFormattedAddress return name of the index
func (p *Property) IdxFormattedAddress() int { //nolint:dupl false positive
	return 18
}

// SqlFormattedAddress return name of the column being indexed
func (p *Property) SqlFormattedAddress() string { //nolint:dupl false positive
	return `"formattedAddress"`
}

// IdxLastPrice return name of the index
func (p *Property) IdxLastPrice() int { //nolint:dupl false positive
	return 19
}

// SqlLastPrice return name of the column being indexed
func (p *Property) SqlLastPrice() string { //nolint:dupl false positive
	return `"lastPrice"`
}

// IdxPriceHistories return name of the index
func (p *Property) IdxPriceHistories() int { //nolint:dupl false positive
	return 20
}

// SqlPriceHistories return name of the column being indexed
func (p *Property) SqlPriceHistories() string { //nolint:dupl false positive
	return `"priceHistories"`
}

// ToArray receiver fields to slice
func (p *Property) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.UniqPropKey,            // 1
		p.SerialNumber,           // 2
		p.SizeM2,                 // 3
		p.MainUse,                // 4
		p.MainBuildingMaterial,   // 5
		p.ConstructCompletedDate, // 6
		p.NumberOfFloors,         // 7
		p.BuildingLamination,     // 8
		p.Address,                // 9
		p.District,               // 10
		p.Note,                   // 11
		p.Coord,                  // 12
		p.CreatedAt,              // 13
		p.CreatedBy,              // 14
		p.UpdatedAt,              // 15
		p.UpdatedBy,              // 16
		p.DeletedAt,              // 17
		p.FormattedAddress,       // 18
		p.LastPrice,              // 19
		p.PriceHistories,         // 20
	}
}

// FromArray convert slice to receiver fields
func (p *Property) FromArray(a A.X) *Property { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.UniqPropKey = X.ToS(a[1])
	p.SerialNumber = X.ToS(a[2])
	p.SizeM2 = X.ToS(a[3])
	p.MainUse = X.ToS(a[4])
	p.MainBuildingMaterial = X.ToS(a[5])
	p.ConstructCompletedDate = X.ToS(a[6])
	p.NumberOfFloors = X.ToS(a[7])
	p.BuildingLamination = X.ToS(a[8])
	p.Address = X.ToS(a[9])
	p.District = X.ToS(a[10])
	p.Note = X.ToS(a[11])
	p.Coord = X.ToArr(a[12])
	p.CreatedAt = X.ToI(a[13])
	p.CreatedBy = X.ToU(a[14])
	p.UpdatedAt = X.ToI(a[15])
	p.UpdatedBy = X.ToU(a[16])
	p.DeletedAt = X.ToI(a[17])
	p.FormattedAddress = X.ToS(a[18])
	p.LastPrice = X.ToS(a[19])
	p.PriceHistories = X.ToArr(a[20])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *Property) FromUncensoredArray(a A.X) *Property { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.UniqPropKey = X.ToS(a[1])
	p.SerialNumber = X.ToS(a[2])
	p.SizeM2 = X.ToS(a[3])
	p.MainUse = X.ToS(a[4])
	p.MainBuildingMaterial = X.ToS(a[5])
	p.ConstructCompletedDate = X.ToS(a[6])
	p.NumberOfFloors = X.ToS(a[7])
	p.BuildingLamination = X.ToS(a[8])
	p.Address = X.ToS(a[9])
	p.District = X.ToS(a[10])
	p.Note = X.ToS(a[11])
	p.Coord = X.ToArr(a[12])
	p.CreatedAt = X.ToI(a[13])
	p.CreatedBy = X.ToU(a[14])
	p.UpdatedAt = X.ToI(a[15])
	p.UpdatedBy = X.ToU(a[16])
	p.DeletedAt = X.ToI(a[17])
	p.FormattedAddress = X.ToS(a[18])
	p.LastPrice = X.ToS(a[19])
	p.PriceHistories = X.ToArr(a[20])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *Property) FindOffsetLimit(offset, limit uint32, idx string) []Property { //nolint:dupl false positive
	var rows []Property
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Property.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Property{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *Property) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Property.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (p *Property) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PropertyFieldTypeMap returns key value of field name and key
var PropertyFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                     Tt.Unsigned,
	`uniqPropKey`:            Tt.String,
	`serialNumber`:           Tt.String,
	`sizeM2`:                 Tt.String,
	`mainUse`:                Tt.String,
	`mainBuildingMaterial`:   Tt.String,
	`constructCompletedDate`: Tt.String,
	`numberOfFloors`:         Tt.String,
	`buildingLamination`:     Tt.String,
	`address`:                Tt.String,
	`district`:               Tt.String,
	`note`:                   Tt.String,
	`coord`:                  Tt.Array,
	`createdAt`:              Tt.Integer,
	`createdBy`:              Tt.Unsigned,
	`updatedAt`:              Tt.Integer,
	`updatedBy`:              Tt.Unsigned,
	`deletedAt`:              Tt.Integer,
	`formattedAddress`:       Tt.String,
	`lastPrice`:              Tt.String,
	`priceHistories`:         Tt.Array,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyHistory DAO reader/query struct
type PropertyHistory struct {
	Adapter               *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id                    uint64
	PropertyKey           string
	TransactionKey        string
	TransactionType       string
	TransactionSign       string
	TransactionTime       string
	TransactionDateNormal string
	TransactionNumber     string
	PriceNtd              int64
	PricePerUnit          int64
	Price                 int64
	Address               string
	District              string
	Note                  string
	CreatedAt             int64
	CreatedBy             uint64
	UpdatedAt             int64
	UpdatedBy             uint64
	DeletedAt             int64
	SerialNumber          string
}

// NewPropertyHistory create new ORM reader/query object
func NewPropertyHistory(adapter *Tt.Adapter) *PropertyHistory {
	return &PropertyHistory{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *PropertyHistory) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TablePropertyHistory) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *PropertyHistory) SqlTableName() string { //nolint:dupl false positive
	return `"property_history"`
}

func (p *PropertyHistory) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *PropertyHistory) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{p.Id})
	if L.IsError(err, `PropertyHistory.FindById failed: `+p.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		p.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexTransactionKey return unique index name
func (p *PropertyHistory) UniqueIndexTransactionKey() string { //nolint:dupl false positive
	return `transactionKey`
}

// FindByTransactionKey Find one by TransactionKey
func (p *PropertyHistory) FindByTransactionKey() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexTransactionKey(), 0, 1, tarantool.IterEq, A.X{p.TransactionKey})
	if L.IsError(err, `PropertyHistory.FindByTransactionKey failed: `+p.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		p.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (p *PropertyHistory) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "propertyKey"
	, "transactionKey"
	, "transactionType"
	, "transactionSign"
	, "transactionTime"
	, "transactionDateNormal"
	, "transactionNumber"
	, "priceNtd"
	, "pricePerUnit"
	, "price"
	, "address"
	, "district"
	, "note"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "serialNumber"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *PropertyHistory) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "propertyKey"
	, "transactionKey"
	, "transactionType"
	, "transactionSign"
	, "transactionTime"
	, "transactionDateNormal"
	, "transactionNumber"
	, "priceNtd"
	, "pricePerUnit"
	, "price"
	, "address"
	, "district"
	, "note"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "serialNumber"
	`
}

// ToUpdateArray generate slice of update command
func (p *PropertyHistory) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, p.Id},
		A.X{`=`, 1, p.PropertyKey},
		A.X{`=`, 2, p.TransactionKey},
		A.X{`=`, 3, p.TransactionType},
		A.X{`=`, 4, p.TransactionSign},
		A.X{`=`, 5, p.TransactionTime},
		A.X{`=`, 6, p.TransactionDateNormal},
		A.X{`=`, 7, p.TransactionNumber},
		A.X{`=`, 8, p.PriceNtd},
		A.X{`=`, 9, p.PricePerUnit},
		A.X{`=`, 10, p.Price},
		A.X{`=`, 11, p.Address},
		A.X{`=`, 12, p.District},
		A.X{`=`, 13, p.Note},
		A.X{`=`, 14, p.CreatedAt},
		A.X{`=`, 15, p.CreatedBy},
		A.X{`=`, 16, p.UpdatedAt},
		A.X{`=`, 17, p.UpdatedBy},
		A.X{`=`, 18, p.DeletedAt},
		A.X{`=`, 19, p.SerialNumber},
	}
}

// IdxId return name of the index
func (p *PropertyHistory) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *PropertyHistory) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxPropertyKey return name of the index
func (p *PropertyHistory) IdxPropertyKey() int { //nolint:dupl false positive
	return 1
}

// SqlPropertyKey return name of the column being indexed
func (p *PropertyHistory) SqlPropertyKey() string { //nolint:dupl false positive
	return `"propertyKey"`
}

// IdxTransactionKey return name of the index
func (p *PropertyHistory) IdxTransactionKey() int { //nolint:dupl false positive
	return 2
}

// SqlTransactionKey return name of the column being indexed
func (p *PropertyHistory) SqlTransactionKey() string { //nolint:dupl false positive
	return `"transactionKey"`
}

// IdxTransactionType return name of the index
func (p *PropertyHistory) IdxTransactionType() int { //nolint:dupl false positive
	return 3
}

// SqlTransactionType return name of the column being indexed
func (p *PropertyHistory) SqlTransactionType() string { //nolint:dupl false positive
	return `"transactionType"`
}

// IdxTransactionSign return name of the index
func (p *PropertyHistory) IdxTransactionSign() int { //nolint:dupl false positive
	return 4
}

// SqlTransactionSign return name of the column being indexed
func (p *PropertyHistory) SqlTransactionSign() string { //nolint:dupl false positive
	return `"transactionSign"`
}

// IdxTransactionTime return name of the index
func (p *PropertyHistory) IdxTransactionTime() int { //nolint:dupl false positive
	return 5
}

// SqlTransactionTime return name of the column being indexed
func (p *PropertyHistory) SqlTransactionTime() string { //nolint:dupl false positive
	return `"transactionTime"`
}

// IdxTransactionDateNormal return name of the index
func (p *PropertyHistory) IdxTransactionDateNormal() int { //nolint:dupl false positive
	return 6
}

// SqlTransactionDateNormal return name of the column being indexed
func (p *PropertyHistory) SqlTransactionDateNormal() string { //nolint:dupl false positive
	return `"transactionDateNormal"`
}

// IdxTransactionNumber return name of the index
func (p *PropertyHistory) IdxTransactionNumber() int { //nolint:dupl false positive
	return 7
}

// SqlTransactionNumber return name of the column being indexed
func (p *PropertyHistory) SqlTransactionNumber() string { //nolint:dupl false positive
	return `"transactionNumber"`
}

// IdxPriceNtd return name of the index
func (p *PropertyHistory) IdxPriceNtd() int { //nolint:dupl false positive
	return 8
}

// SqlPriceNtd return name of the column being indexed
func (p *PropertyHistory) SqlPriceNtd() string { //nolint:dupl false positive
	return `"priceNtd"`
}

// IdxPricePerUnit return name of the index
func (p *PropertyHistory) IdxPricePerUnit() int { //nolint:dupl false positive
	return 9
}

// SqlPricePerUnit return name of the column being indexed
func (p *PropertyHistory) SqlPricePerUnit() string { //nolint:dupl false positive
	return `"pricePerUnit"`
}

// IdxPrice return name of the index
func (p *PropertyHistory) IdxPrice() int { //nolint:dupl false positive
	return 10
}

// SqlPrice return name of the column being indexed
func (p *PropertyHistory) SqlPrice() string { //nolint:dupl false positive
	return `"price"`
}

// IdxAddress return name of the index
func (p *PropertyHistory) IdxAddress() int { //nolint:dupl false positive
	return 11
}

// SqlAddress return name of the column being indexed
func (p *PropertyHistory) SqlAddress() string { //nolint:dupl false positive
	return `"address"`
}

// IdxDistrict return name of the index
func (p *PropertyHistory) IdxDistrict() int { //nolint:dupl false positive
	return 12
}

// SqlDistrict return name of the column being indexed
func (p *PropertyHistory) SqlDistrict() string { //nolint:dupl false positive
	return `"district"`
}

// IdxNote return name of the index
func (p *PropertyHistory) IdxNote() int { //nolint:dupl false positive
	return 13
}

// SqlNote return name of the column being indexed
func (p *PropertyHistory) SqlNote() string { //nolint:dupl false positive
	return `"note"`
}

// IdxCreatedAt return name of the index
func (p *PropertyHistory) IdxCreatedAt() int { //nolint:dupl false positive
	return 14
}

// SqlCreatedAt return name of the column being indexed
func (p *PropertyHistory) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *PropertyHistory) IdxCreatedBy() int { //nolint:dupl false positive
	return 15
}

// SqlCreatedBy return name of the column being indexed
func (p *PropertyHistory) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *PropertyHistory) IdxUpdatedAt() int { //nolint:dupl false positive
	return 16
}

// SqlUpdatedAt return name of the column being indexed
func (p *PropertyHistory) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *PropertyHistory) IdxUpdatedBy() int { //nolint:dupl false positive
	return 17
}

// SqlUpdatedBy return name of the column being indexed
func (p *PropertyHistory) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *PropertyHistory) IdxDeletedAt() int { //nolint:dupl false positive
	return 18
}

// SqlDeletedAt return name of the column being indexed
func (p *PropertyHistory) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxSerialNumber return name of the index
func (p *PropertyHistory) IdxSerialNumber() int { //nolint:dupl false positive
	return 19
}

// SqlSerialNumber return name of the column being indexed
func (p *PropertyHistory) SqlSerialNumber() string { //nolint:dupl false positive
	return `"serialNumber"`
}

// ToArray receiver fields to slice
func (p *PropertyHistory) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.PropertyKey,           // 1
		p.TransactionKey,        // 2
		p.TransactionType,       // 3
		p.TransactionSign,       // 4
		p.TransactionTime,       // 5
		p.TransactionDateNormal, // 6
		p.TransactionNumber,     // 7
		p.PriceNtd,              // 8
		p.PricePerUnit,          // 9
		p.Price,                 // 10
		p.Address,               // 11
		p.District,              // 12
		p.Note,                  // 13
		p.CreatedAt,             // 14
		p.CreatedBy,             // 15
		p.UpdatedAt,             // 16
		p.UpdatedBy,             // 17
		p.DeletedAt,             // 18
		p.SerialNumber,          // 19
	}
}

// FromArray convert slice to receiver fields
func (p *PropertyHistory) FromArray(a A.X) *PropertyHistory { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PropertyKey = X.ToS(a[1])
	p.TransactionKey = X.ToS(a[2])
	p.TransactionType = X.ToS(a[3])
	p.TransactionSign = X.ToS(a[4])
	p.TransactionTime = X.ToS(a[5])
	p.TransactionDateNormal = X.ToS(a[6])
	p.TransactionNumber = X.ToS(a[7])
	p.PriceNtd = X.ToI(a[8])
	p.PricePerUnit = X.ToI(a[9])
	p.Price = X.ToI(a[10])
	p.Address = X.ToS(a[11])
	p.District = X.ToS(a[12])
	p.Note = X.ToS(a[13])
	p.CreatedAt = X.ToI(a[14])
	p.CreatedBy = X.ToU(a[15])
	p.UpdatedAt = X.ToI(a[16])
	p.UpdatedBy = X.ToU(a[17])
	p.DeletedAt = X.ToI(a[18])
	p.SerialNumber = X.ToS(a[19])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *PropertyHistory) FromUncensoredArray(a A.X) *PropertyHistory { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PropertyKey = X.ToS(a[1])
	p.TransactionKey = X.ToS(a[2])
	p.TransactionType = X.ToS(a[3])
	p.TransactionSign = X.ToS(a[4])
	p.TransactionTime = X.ToS(a[5])
	p.TransactionDateNormal = X.ToS(a[6])
	p.TransactionNumber = X.ToS(a[7])
	p.PriceNtd = X.ToI(a[8])
	p.PricePerUnit = X.ToI(a[9])
	p.Price = X.ToI(a[10])
	p.Address = X.ToS(a[11])
	p.District = X.ToS(a[12])
	p.Note = X.ToS(a[13])
	p.CreatedAt = X.ToI(a[14])
	p.CreatedBy = X.ToU(a[15])
	p.UpdatedAt = X.ToI(a[16])
	p.UpdatedBy = X.ToU(a[17])
	p.DeletedAt = X.ToI(a[18])
	p.SerialNumber = X.ToS(a[19])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *PropertyHistory) FindOffsetLimit(offset, limit uint32, idx string) []PropertyHistory { //nolint:dupl false positive
	var rows []PropertyHistory
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropertyHistory.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := PropertyHistory{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *PropertyHistory) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropertyHistory.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (p *PropertyHistory) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PropertyHistoryFieldTypeMap returns key value of field name and key
var PropertyHistoryFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                    Tt.Unsigned,
	`propertyKey`:           Tt.String,
	`transactionKey`:        Tt.String,
	`transactionType`:       Tt.String,
	`transactionSign`:       Tt.String,
	`transactionTime`:       Tt.String,
	`transactionDateNormal`: Tt.String,
	`transactionNumber`:     Tt.String,
	`priceNtd`:              Tt.Integer,
	`pricePerUnit`:          Tt.Integer,
	`price`:                 Tt.Integer,
	`address`:               Tt.String,
	`district`:              Tt.String,
	`note`:                  Tt.String,
	`createdAt`:             Tt.Integer,
	`createdBy`:             Tt.Unsigned,
	`updatedAt`:             Tt.Integer,
	`updatedBy`:             Tt.Unsigned,
	`deletedAt`:             Tt.Integer,
	`serialNumber`:          Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
