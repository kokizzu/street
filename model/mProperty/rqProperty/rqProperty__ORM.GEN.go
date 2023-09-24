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

// PropLikeCount DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqProperty__ORM.GEN.go
type PropLikeCount struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	PropId  uint64      `json:"propId,string" form:"propId" query:"propId" long:"propId" msg:"propId"`
	Count   int64       `json:"count" form:"count" query:"count" long:"count" msg:"count"`
}

// NewPropLikeCount create new ORM reader/query object
func NewPropLikeCount(adapter *Tt.Adapter) *PropLikeCount {
	return &PropLikeCount{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *PropLikeCount) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TablePropLikeCount) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *PropLikeCount) SqlTableName() string { //nolint:dupl false positive
	return `"propLikeCount"`
}

// UniqueIndexPropId return unique index name
func (p *PropLikeCount) UniqueIndexPropId() string { //nolint:dupl false positive
	return `propId`
}

// FindByPropId Find one by PropId
func (p *PropLikeCount) FindByPropId() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexPropId(), 0, 1, tarantool.IterEq, A.X{p.PropId})
	if L.IsError(err, `PropLikeCount.FindByPropId failed: `+p.SpaceName()) {
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
func (p *PropLikeCount) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "propId"
	, "count"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *PropLikeCount) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "propId"
	, "count"
	`
}

// ToUpdateArray generate slice of update command
func (p *PropLikeCount) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, p.PropId},
		A.X{`=`, 1, p.Count},
	}
}

// IdxPropId return name of the index
func (p *PropLikeCount) IdxPropId() int { //nolint:dupl false positive
	return 0
}

// SqlPropId return name of the column being indexed
func (p *PropLikeCount) SqlPropId() string { //nolint:dupl false positive
	return `"propId"`
}

// IdxCount return name of the index
func (p *PropLikeCount) IdxCount() int { //nolint:dupl false positive
	return 1
}

// SqlCount return name of the column being indexed
func (p *PropLikeCount) SqlCount() string { //nolint:dupl false positive
	return `"count"`
}

// ToArray receiver fields to slice
func (p *PropLikeCount) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		p.PropId, // 0
		p.Count,  // 1
	}
}

// FromArray convert slice to receiver fields
func (p *PropLikeCount) FromArray(a A.X) *PropLikeCount { //nolint:dupl false positive
	p.PropId = X.ToU(a[0])
	p.Count = X.ToI(a[1])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *PropLikeCount) FromUncensoredArray(a A.X) *PropLikeCount { //nolint:dupl false positive
	p.PropId = X.ToU(a[0])
	p.Count = X.ToI(a[1])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *PropLikeCount) FindOffsetLimit(offset, limit uint32, idx string) []PropLikeCount { //nolint:dupl false positive
	var rows []PropLikeCount
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropLikeCount.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := PropLikeCount{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *PropLikeCount) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropLikeCount.FindOffsetLimit failed: `+p.SpaceName()) {
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
func (p *PropLikeCount) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PropLikeCountFieldTypeMap returns key value of field name and key
var PropLikeCountFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`propId`: Tt.Unsigned,
	`count`:  Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Property DAO reader/query struct
type Property struct {
	Adapter                 *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                      uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	UniqPropKey             string      `json:"uniqPropKey" form:"uniqPropKey" query:"uniqPropKey" long:"uniqPropKey" msg:"uniqPropKey"`
	SerialNumber            string      `json:"serialNumber" form:"serialNumber" query:"serialNumber" long:"serialNumber" msg:"serialNumber"`
	SizeM2                  string      `json:"sizeM2" form:"sizeM2" query:"sizeM2" long:"sizeM2" msg:"sizeM2"`
	MainUse                 string      `json:"mainUse" form:"mainUse" query:"mainUse" long:"mainUse" msg:"mainUse"`
	MainBuildingMaterial    string      `json:"mainBuildingMaterial" form:"mainBuildingMaterial" query:"mainBuildingMaterial" long:"mainBuildingMaterial" msg:"mainBuildingMaterial"`
	ConstructCompletedDate  string      `json:"constructCompletedDate" form:"constructCompletedDate" query:"constructCompletedDate" long:"constructCompletedDate" msg:"constructCompletedDate"`
	NumberOfFloors          string      `json:"numberOfFloors" form:"numberOfFloors" query:"numberOfFloors" long:"numberOfFloors" msg:"numberOfFloors"`
	BuildingLamination      string      `json:"buildingLamination" form:"buildingLamination" query:"buildingLamination" long:"buildingLamination" msg:"buildingLamination"`
	Address                 string      `json:"address" form:"address" query:"address" long:"address" msg:"address"`
	District                string      `json:"district" form:"district" query:"district" long:"district" msg:"district"`
	Note                    string      `json:"note" form:"note" query:"note" long:"note" msg:"note"`
	Coord                   []any       `json:"coord" form:"coord" query:"coord" long:"coord" msg:"coord"`
	CreatedAt               int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy               uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt               int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy               uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt               int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	FormattedAddress        string      `json:"formattedAddress" form:"formattedAddress" query:"formattedAddress" long:"formattedAddress" msg:"formattedAddress"`
	LastPrice               string      `json:"lastPrice" form:"lastPrice" query:"lastPrice" long:"lastPrice" msg:"lastPrice"`
	PriceHistoriesSell      []any       `json:"priceHistoriesSell" form:"priceHistoriesSell" query:"priceHistoriesSell" long:"priceHistoriesSell" msg:"priceHistoriesSell"`
	PriceHistoriesRent      []any       `json:"priceHistoriesRent" form:"priceHistoriesRent" query:"priceHistoriesRent" long:"priceHistoriesRent" msg:"priceHistoriesRent"`
	Purpose                 string      `json:"purpose" form:"purpose" query:"purpose" long:"purpose" msg:"purpose"`
	HouseType               string      `json:"houseType" form:"houseType" query:"houseType" long:"houseType" msg:"houseType"`
	Images                  []any       `json:"images" form:"images" query:"images" long:"images" msg:"images"`
	Bedroom                 int64       `json:"bedroom" form:"bedroom" query:"bedroom" long:"bedroom" msg:"bedroom"`
	Bathroom                int64       `json:"bathroom" form:"bathroom" query:"bathroom" long:"bathroom" msg:"bathroom"`
	AgencyFeePercent        float64     `json:"agencyFeePercent" form:"agencyFeePercent" query:"agencyFeePercent" long:"agencyFeePercent" msg:"agencyFeePercent"`
	FloorList               []any       `json:"floorList" form:"floorList" query:"floorList" long:"floorList" msg:"floorList"`
	Version                 string      `json:"version" form:"version" query:"version" long:"version" msg:"version"`
	YearBuilt               int64       `json:"yearBuilt" form:"yearBuilt" query:"yearBuilt" long:"yearBuilt" msg:"yearBuilt"`
	YearRenovated           int64       `json:"yearRenovated" form:"yearRenovated" query:"yearRenovated" long:"yearRenovated" msg:"yearRenovated"`
	TotalSqft               float64     `json:"totalSqft" form:"totalSqft" query:"totalSqft" long:"totalSqft" msg:"totalSqft"`
	CountyName              string      `json:"countyName" form:"countyName" query:"countyName" long:"countyName" msg:"countyName"`
	Street                  string      `json:"street" form:"street" query:"street" long:"street" msg:"street"`
	City                    string      `json:"city" form:"city" query:"city" long:"city" msg:"city"`
	State                   string      `json:"state" form:"state" query:"state" long:"state" msg:"state"`
	Zip                     string      `json:"zip" form:"zip" query:"zip" long:"zip" msg:"zip"`
	PropertyLastUpdatedDate int64       `json:"propertyLastUpdatedDate" form:"propertyLastUpdatedDate" query:"propertyLastUpdatedDate" long:"propertyLastUpdatedDate" msg:"propertyLastUpdatedDate"`
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
	, "priceHistoriesSell"
	, "priceHistoriesRent"
	, "purpose"
	, "houseType"
	, "images"
	, "bedroom"
	, "bathroom"
	, "agencyFeePercent"
	, "floorList"
	, "version"
	, "yearBuilt"
	, "yearRenovated"
	, "totalSqft"
	, "countyName"
	, "street"
	, "city"
	, "state"
	, "zip"
	, "propertyLastUpdatedDate"
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
	, "priceHistoriesSell"
	, "priceHistoriesRent"
	, "purpose"
	, "houseType"
	, "images"
	, "bedroom"
	, "bathroom"
	, "agencyFeePercent"
	, "floorList"
	, "version"
	, "yearBuilt"
	, "yearRenovated"
	, "totalSqft"
	, "countyName"
	, "street"
	, "city"
	, "state"
	, "zip"
	, "propertyLastUpdatedDate"
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
		A.X{`=`, 20, p.PriceHistoriesSell},
		A.X{`=`, 21, p.PriceHistoriesRent},
		A.X{`=`, 22, p.Purpose},
		A.X{`=`, 23, p.HouseType},
		A.X{`=`, 24, p.Images},
		A.X{`=`, 25, p.Bedroom},
		A.X{`=`, 26, p.Bathroom},
		A.X{`=`, 27, p.AgencyFeePercent},
		A.X{`=`, 28, p.FloorList},
		A.X{`=`, 29, p.Version},
		A.X{`=`, 30, p.YearBuilt},
		A.X{`=`, 31, p.YearRenovated},
		A.X{`=`, 32, p.TotalSqft},
		A.X{`=`, 33, p.CountyName},
		A.X{`=`, 34, p.Street},
		A.X{`=`, 35, p.City},
		A.X{`=`, 36, p.State},
		A.X{`=`, 37, p.Zip},
		A.X{`=`, 38, p.PropertyLastUpdatedDate},
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

// IdxPriceHistoriesSell return name of the index
func (p *Property) IdxPriceHistoriesSell() int { //nolint:dupl false positive
	return 20
}

// SqlPriceHistoriesSell return name of the column being indexed
func (p *Property) SqlPriceHistoriesSell() string { //nolint:dupl false positive
	return `"priceHistoriesSell"`
}

// IdxPriceHistoriesRent return name of the index
func (p *Property) IdxPriceHistoriesRent() int { //nolint:dupl false positive
	return 21
}

// SqlPriceHistoriesRent return name of the column being indexed
func (p *Property) SqlPriceHistoriesRent() string { //nolint:dupl false positive
	return `"priceHistoriesRent"`
}

// IdxPurpose return name of the index
func (p *Property) IdxPurpose() int { //nolint:dupl false positive
	return 22
}

// SqlPurpose return name of the column being indexed
func (p *Property) SqlPurpose() string { //nolint:dupl false positive
	return `"purpose"`
}

// IdxHouseType return name of the index
func (p *Property) IdxHouseType() int { //nolint:dupl false positive
	return 23
}

// SqlHouseType return name of the column being indexed
func (p *Property) SqlHouseType() string { //nolint:dupl false positive
	return `"houseType"`
}

// IdxImages return name of the index
func (p *Property) IdxImages() int { //nolint:dupl false positive
	return 24
}

// SqlImages return name of the column being indexed
func (p *Property) SqlImages() string { //nolint:dupl false positive
	return `"images"`
}

// IdxBedroom return name of the index
func (p *Property) IdxBedroom() int { //nolint:dupl false positive
	return 25
}

// SqlBedroom return name of the column being indexed
func (p *Property) SqlBedroom() string { //nolint:dupl false positive
	return `"bedroom"`
}

// IdxBathroom return name of the index
func (p *Property) IdxBathroom() int { //nolint:dupl false positive
	return 26
}

// SqlBathroom return name of the column being indexed
func (p *Property) SqlBathroom() string { //nolint:dupl false positive
	return `"bathroom"`
}

// IdxAgencyFeePercent return name of the index
func (p *Property) IdxAgencyFeePercent() int { //nolint:dupl false positive
	return 27
}

// SqlAgencyFeePercent return name of the column being indexed
func (p *Property) SqlAgencyFeePercent() string { //nolint:dupl false positive
	return `"agencyFeePercent"`
}

// IdxFloorList return name of the index
func (p *Property) IdxFloorList() int { //nolint:dupl false positive
	return 28
}

// SqlFloorList return name of the column being indexed
func (p *Property) SqlFloorList() string { //nolint:dupl false positive
	return `"floorList"`
}

// IdxVersion return name of the index
func (p *Property) IdxVersion() int { //nolint:dupl false positive
	return 29
}

// SqlVersion return name of the column being indexed
func (p *Property) SqlVersion() string { //nolint:dupl false positive
	return `"version"`
}

// IdxYearBuilt return name of the index
func (p *Property) IdxYearBuilt() int { //nolint:dupl false positive
	return 30
}

// SqlYearBuilt return name of the column being indexed
func (p *Property) SqlYearBuilt() string { //nolint:dupl false positive
	return `"yearBuilt"`
}

// IdxYearRenovated return name of the index
func (p *Property) IdxYearRenovated() int { //nolint:dupl false positive
	return 31
}

// SqlYearRenovated return name of the column being indexed
func (p *Property) SqlYearRenovated() string { //nolint:dupl false positive
	return `"yearRenovated"`
}

// IdxTotalSqft return name of the index
func (p *Property) IdxTotalSqft() int { //nolint:dupl false positive
	return 32
}

// SqlTotalSqft return name of the column being indexed
func (p *Property) SqlTotalSqft() string { //nolint:dupl false positive
	return `"totalSqft"`
}

// IdxCountyName return name of the index
func (p *Property) IdxCountyName() int { //nolint:dupl false positive
	return 33
}

// SqlCountyName return name of the column being indexed
func (p *Property) SqlCountyName() string { //nolint:dupl false positive
	return `"countyName"`
}

// IdxStreet return name of the index
func (p *Property) IdxStreet() int { //nolint:dupl false positive
	return 34
}

// SqlStreet return name of the column being indexed
func (p *Property) SqlStreet() string { //nolint:dupl false positive
	return `"street"`
}

// IdxCity return name of the index
func (p *Property) IdxCity() int { //nolint:dupl false positive
	return 35
}

// SqlCity return name of the column being indexed
func (p *Property) SqlCity() string { //nolint:dupl false positive
	return `"city"`
}

// IdxState return name of the index
func (p *Property) IdxState() int { //nolint:dupl false positive
	return 36
}

// SqlState return name of the column being indexed
func (p *Property) SqlState() string { //nolint:dupl false positive
	return `"state"`
}

// IdxZip return name of the index
func (p *Property) IdxZip() int { //nolint:dupl false positive
	return 37
}

// SqlZip return name of the column being indexed
func (p *Property) SqlZip() string { //nolint:dupl false positive
	return `"zip"`
}

// IdxPropertyLastUpdatedDate return name of the index
func (p *Property) IdxPropertyLastUpdatedDate() int { //nolint:dupl false positive
	return 38
}

// SqlPropertyLastUpdatedDate return name of the column being indexed
func (p *Property) SqlPropertyLastUpdatedDate() string { //nolint:dupl false positive
	return `"propertyLastUpdatedDate"`
}

// ToArray receiver fields to slice
func (p *Property) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.UniqPropKey,             // 1
		p.SerialNumber,            // 2
		p.SizeM2,                  // 3
		p.MainUse,                 // 4
		p.MainBuildingMaterial,    // 5
		p.ConstructCompletedDate,  // 6
		p.NumberOfFloors,          // 7
		p.BuildingLamination,      // 8
		p.Address,                 // 9
		p.District,                // 10
		p.Note,                    // 11
		p.Coord,                   // 12
		p.CreatedAt,               // 13
		p.CreatedBy,               // 14
		p.UpdatedAt,               // 15
		p.UpdatedBy,               // 16
		p.DeletedAt,               // 17
		p.FormattedAddress,        // 18
		p.LastPrice,               // 19
		p.PriceHistoriesSell,      // 20
		p.PriceHistoriesRent,      // 21
		p.Purpose,                 // 22
		p.HouseType,               // 23
		p.Images,                  // 24
		p.Bedroom,                 // 25
		p.Bathroom,                // 26
		p.AgencyFeePercent,        // 27
		p.FloorList,               // 28
		p.Version,                 // 29
		p.YearBuilt,               // 30
		p.YearRenovated,           // 31
		p.TotalSqft,               // 32
		p.CountyName,              // 33
		p.Street,                  // 34
		p.City,                    // 35
		p.State,                   // 36
		p.Zip,                     // 37
		p.PropertyLastUpdatedDate, // 38
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
	p.PriceHistoriesSell = X.ToArr(a[20])
	p.PriceHistoriesRent = X.ToArr(a[21])
	p.Purpose = X.ToS(a[22])
	p.HouseType = X.ToS(a[23])
	p.Images = X.ToArr(a[24])
	p.Bedroom = X.ToI(a[25])
	p.Bathroom = X.ToI(a[26])
	p.AgencyFeePercent = X.ToF(a[27])
	p.FloorList = X.ToArr(a[28])
	p.Version = X.ToS(a[29])
	p.YearBuilt = X.ToI(a[30])
	p.YearRenovated = X.ToI(a[31])
	p.TotalSqft = X.ToF(a[32])
	p.CountyName = X.ToS(a[33])
	p.Street = X.ToS(a[34])
	p.City = X.ToS(a[35])
	p.State = X.ToS(a[36])
	p.Zip = X.ToS(a[37])
	p.PropertyLastUpdatedDate = X.ToI(a[38])
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
	p.PriceHistoriesSell = X.ToArr(a[20])
	p.PriceHistoriesRent = X.ToArr(a[21])
	p.Purpose = X.ToS(a[22])
	p.HouseType = X.ToS(a[23])
	p.Images = X.ToArr(a[24])
	p.Bedroom = X.ToI(a[25])
	p.Bathroom = X.ToI(a[26])
	p.AgencyFeePercent = X.ToF(a[27])
	p.FloorList = X.ToArr(a[28])
	p.Version = X.ToS(a[29])
	p.YearBuilt = X.ToI(a[30])
	p.YearRenovated = X.ToI(a[31])
	p.TotalSqft = X.ToF(a[32])
	p.CountyName = X.ToS(a[33])
	p.Street = X.ToS(a[34])
	p.City = X.ToS(a[35])
	p.State = X.ToS(a[36])
	p.Zip = X.ToS(a[37])
	p.PropertyLastUpdatedDate = X.ToI(a[38])
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
	`id`:                      Tt.Unsigned,
	`uniqPropKey`:             Tt.String,
	`serialNumber`:            Tt.String,
	`sizeM2`:                  Tt.String,
	`mainUse`:                 Tt.String,
	`mainBuildingMaterial`:    Tt.String,
	`constructCompletedDate`:  Tt.String,
	`numberOfFloors`:          Tt.String,
	`buildingLamination`:      Tt.String,
	`address`:                 Tt.String,
	`district`:                Tt.String,
	`note`:                    Tt.String,
	`coord`:                   Tt.Array,
	`createdAt`:               Tt.Integer,
	`createdBy`:               Tt.Unsigned,
	`updatedAt`:               Tt.Integer,
	`updatedBy`:               Tt.Unsigned,
	`deletedAt`:               Tt.Integer,
	`formattedAddress`:        Tt.String,
	`lastPrice`:               Tt.String,
	`priceHistoriesSell`:      Tt.Array,
	`priceHistoriesRent`:      Tt.Array,
	`purpose`:                 Tt.String,
	`houseType`:               Tt.String,
	`images`:                  Tt.Array,
	`bedroom`:                 Tt.Integer,
	`bathroom`:                Tt.Integer,
	`agencyFeePercent`:        Tt.Double,
	`floorList`:               Tt.Array,
	`version`:                 Tt.String,
	`yearBuilt`:               Tt.Integer,
	`yearRenovated`:           Tt.Integer,
	`totalSqft`:               Tt.Double,
	`countyName`:              Tt.String,
	`street`:                  Tt.String,
	`city`:                    Tt.String,
	`state`:                   Tt.String,
	`zip`:                     Tt.String,
	`propertyLastUpdatedDate`: Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyExtraUS DAO reader/query struct
type PropertyExtraUS struct {
	Adapter            *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                 uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	PropertyKey        string      `json:"propertyKey" form:"propertyKey" query:"propertyKey" long:"propertyKey" msg:"propertyKey"`
	CountyUrl          string      `json:"countyUrl" form:"countyUrl" query:"countyUrl" long:"countyUrl" msg:"countyUrl"`
	CountyIsActive     bool        `json:"countyIsActive" form:"countyIsActive" query:"countyIsActive" long:"countyIsActive" msg:"countyIsActive"`
	ZoneDataInfo       string      `json:"zoneDataInfo" form:"zoneDataInfo" query:"zoneDataInfo" long:"zoneDataInfo" msg:"zoneDataInfo"`
	TaxInfo            string      `json:"taxInfo" form:"taxInfo" query:"taxInfo" long:"taxInfo" msg:"taxInfo"`
	HistoryTaxInfo     string      `json:"historyTaxInfo" form:"historyTaxInfo" query:"historyTaxInfo" long:"historyTaxInfo" msg:"historyTaxInfo"`
	AmenitySuperGroups string      `json:"amenitySuperGroups" form:"amenitySuperGroups" query:"amenitySuperGroups" long:"amenitySuperGroups" msg:"amenitySuperGroups"`
	MlsDisclaimerInfo  string      `json:"mlsDisclaimerInfo" form:"mlsDisclaimerInfo" query:"mlsDisclaimerInfo" long:"mlsDisclaimerInfo" msg:"mlsDisclaimerInfo"`
	FacilityInfo       string      `json:"facilityInfo" form:"facilityInfo" query:"facilityInfo" long:"facilityInfo" msg:"facilityInfo"`
	RiskInfo           string      `json:"riskInfo" form:"riskInfo" query:"riskInfo" long:"riskInfo" msg:"riskInfo"`
	MediaSourceJson    string      `json:"mediaSourceJson" form:"mediaSourceJson" query:"mediaSourceJson" long:"mediaSourceJson" msg:"mediaSourceJson"`
	TaxNote            string      `json:"taxNote" form:"taxNote" query:"taxNote" long:"taxNote" msg:"taxNote"`
}

// NewPropertyExtraUS create new ORM reader/query object
func NewPropertyExtraUS(adapter *Tt.Adapter) *PropertyExtraUS {
	return &PropertyExtraUS{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *PropertyExtraUS) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TablePropertyExtraUS) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *PropertyExtraUS) SqlTableName() string { //nolint:dupl false positive
	return `"propertyExtraUS"`
}

func (p *PropertyExtraUS) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *PropertyExtraUS) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{p.Id})
	if L.IsError(err, `PropertyExtraUS.FindById failed: `+p.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		p.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexPropertyKey return unique index name
func (p *PropertyExtraUS) UniqueIndexPropertyKey() string { //nolint:dupl false positive
	return `propertyKey`
}

// FindByPropertyKey Find one by PropertyKey
func (p *PropertyExtraUS) FindByPropertyKey() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexPropertyKey(), 0, 1, tarantool.IterEq, A.X{p.PropertyKey})
	if L.IsError(err, `PropertyExtraUS.FindByPropertyKey failed: `+p.SpaceName()) {
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
func (p *PropertyExtraUS) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "propertyKey"
	, "countyUrl"
	, "countyIsActive"
	, "zoneDataInfo"
	, "taxInfo"
	, "historyTaxInfo"
	, "amenitySuperGroups"
	, "mlsDisclaimerInfo"
	, "facilityInfo"
	, "riskInfo"
	, "mediaSourceJson"
	, "taxNote"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *PropertyExtraUS) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "propertyKey"
	, "countyUrl"
	, "countyIsActive"
	, "zoneDataInfo"
	, "taxInfo"
	, "historyTaxInfo"
	, "amenitySuperGroups"
	, "mlsDisclaimerInfo"
	, "facilityInfo"
	, "riskInfo"
	, "mediaSourceJson"
	, "taxNote"
	`
}

// ToUpdateArray generate slice of update command
func (p *PropertyExtraUS) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, p.Id},
		A.X{`=`, 1, p.PropertyKey},
		A.X{`=`, 2, p.CountyUrl},
		A.X{`=`, 3, p.CountyIsActive},
		A.X{`=`, 4, p.ZoneDataInfo},
		A.X{`=`, 5, p.TaxInfo},
		A.X{`=`, 6, p.HistoryTaxInfo},
		A.X{`=`, 7, p.AmenitySuperGroups},
		A.X{`=`, 8, p.MlsDisclaimerInfo},
		A.X{`=`, 9, p.FacilityInfo},
		A.X{`=`, 10, p.RiskInfo},
		A.X{`=`, 11, p.MediaSourceJson},
		A.X{`=`, 12, p.TaxNote},
	}
}

// IdxId return name of the index
func (p *PropertyExtraUS) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *PropertyExtraUS) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxPropertyKey return name of the index
func (p *PropertyExtraUS) IdxPropertyKey() int { //nolint:dupl false positive
	return 1
}

// SqlPropertyKey return name of the column being indexed
func (p *PropertyExtraUS) SqlPropertyKey() string { //nolint:dupl false positive
	return `"propertyKey"`
}

// IdxCountyUrl return name of the index
func (p *PropertyExtraUS) IdxCountyUrl() int { //nolint:dupl false positive
	return 2
}

// SqlCountyUrl return name of the column being indexed
func (p *PropertyExtraUS) SqlCountyUrl() string { //nolint:dupl false positive
	return `"countyUrl"`
}

// IdxCountyIsActive return name of the index
func (p *PropertyExtraUS) IdxCountyIsActive() int { //nolint:dupl false positive
	return 3
}

// SqlCountyIsActive return name of the column being indexed
func (p *PropertyExtraUS) SqlCountyIsActive() string { //nolint:dupl false positive
	return `"countyIsActive"`
}

// IdxZoneDataInfo return name of the index
func (p *PropertyExtraUS) IdxZoneDataInfo() int { //nolint:dupl false positive
	return 4
}

// SqlZoneDataInfo return name of the column being indexed
func (p *PropertyExtraUS) SqlZoneDataInfo() string { //nolint:dupl false positive
	return `"zoneDataInfo"`
}

// IdxTaxInfo return name of the index
func (p *PropertyExtraUS) IdxTaxInfo() int { //nolint:dupl false positive
	return 5
}

// SqlTaxInfo return name of the column being indexed
func (p *PropertyExtraUS) SqlTaxInfo() string { //nolint:dupl false positive
	return `"taxInfo"`
}

// IdxHistoryTaxInfo return name of the index
func (p *PropertyExtraUS) IdxHistoryTaxInfo() int { //nolint:dupl false positive
	return 6
}

// SqlHistoryTaxInfo return name of the column being indexed
func (p *PropertyExtraUS) SqlHistoryTaxInfo() string { //nolint:dupl false positive
	return `"historyTaxInfo"`
}

// IdxAmenitySuperGroups return name of the index
func (p *PropertyExtraUS) IdxAmenitySuperGroups() int { //nolint:dupl false positive
	return 7
}

// SqlAmenitySuperGroups return name of the column being indexed
func (p *PropertyExtraUS) SqlAmenitySuperGroups() string { //nolint:dupl false positive
	return `"amenitySuperGroups"`
}

// IdxMlsDisclaimerInfo return name of the index
func (p *PropertyExtraUS) IdxMlsDisclaimerInfo() int { //nolint:dupl false positive
	return 8
}

// SqlMlsDisclaimerInfo return name of the column being indexed
func (p *PropertyExtraUS) SqlMlsDisclaimerInfo() string { //nolint:dupl false positive
	return `"mlsDisclaimerInfo"`
}

// IdxFacilityInfo return name of the index
func (p *PropertyExtraUS) IdxFacilityInfo() int { //nolint:dupl false positive
	return 9
}

// SqlFacilityInfo return name of the column being indexed
func (p *PropertyExtraUS) SqlFacilityInfo() string { //nolint:dupl false positive
	return `"facilityInfo"`
}

// IdxRiskInfo return name of the index
func (p *PropertyExtraUS) IdxRiskInfo() int { //nolint:dupl false positive
	return 10
}

// SqlRiskInfo return name of the column being indexed
func (p *PropertyExtraUS) SqlRiskInfo() string { //nolint:dupl false positive
	return `"riskInfo"`
}

// IdxMediaSourceJson return name of the index
func (p *PropertyExtraUS) IdxMediaSourceJson() int { //nolint:dupl false positive
	return 11
}

// SqlMediaSourceJson return name of the column being indexed
func (p *PropertyExtraUS) SqlMediaSourceJson() string { //nolint:dupl false positive
	return `"mediaSourceJson"`
}

// IdxTaxNote return name of the index
func (p *PropertyExtraUS) IdxTaxNote() int { //nolint:dupl false positive
	return 12
}

// SqlTaxNote return name of the column being indexed
func (p *PropertyExtraUS) SqlTaxNote() string { //nolint:dupl false positive
	return `"taxNote"`
}

// ToArray receiver fields to slice
func (p *PropertyExtraUS) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.PropertyKey,        // 1
		p.CountyUrl,          // 2
		p.CountyIsActive,     // 3
		p.ZoneDataInfo,       // 4
		p.TaxInfo,            // 5
		p.HistoryTaxInfo,     // 6
		p.AmenitySuperGroups, // 7
		p.MlsDisclaimerInfo,  // 8
		p.FacilityInfo,       // 9
		p.RiskInfo,           // 10
		p.MediaSourceJson,    // 11
		p.TaxNote,            // 12
	}
}

// FromArray convert slice to receiver fields
func (p *PropertyExtraUS) FromArray(a A.X) *PropertyExtraUS { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PropertyKey = X.ToS(a[1])
	p.CountyUrl = X.ToS(a[2])
	p.CountyIsActive = X.ToBool(a[3])
	p.ZoneDataInfo = X.ToS(a[4])
	p.TaxInfo = X.ToS(a[5])
	p.HistoryTaxInfo = X.ToS(a[6])
	p.AmenitySuperGroups = X.ToS(a[7])
	p.MlsDisclaimerInfo = X.ToS(a[8])
	p.FacilityInfo = X.ToS(a[9])
	p.RiskInfo = X.ToS(a[10])
	p.MediaSourceJson = X.ToS(a[11])
	p.TaxNote = X.ToS(a[12])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *PropertyExtraUS) FromUncensoredArray(a A.X) *PropertyExtraUS { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PropertyKey = X.ToS(a[1])
	p.CountyUrl = X.ToS(a[2])
	p.CountyIsActive = X.ToBool(a[3])
	p.ZoneDataInfo = X.ToS(a[4])
	p.TaxInfo = X.ToS(a[5])
	p.HistoryTaxInfo = X.ToS(a[6])
	p.AmenitySuperGroups = X.ToS(a[7])
	p.MlsDisclaimerInfo = X.ToS(a[8])
	p.FacilityInfo = X.ToS(a[9])
	p.RiskInfo = X.ToS(a[10])
	p.MediaSourceJson = X.ToS(a[11])
	p.TaxNote = X.ToS(a[12])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *PropertyExtraUS) FindOffsetLimit(offset, limit uint32, idx string) []PropertyExtraUS { //nolint:dupl false positive
	var rows []PropertyExtraUS
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropertyExtraUS.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := PropertyExtraUS{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *PropertyExtraUS) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropertyExtraUS.FindOffsetLimit failed: `+p.SpaceName()) {
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
func (p *PropertyExtraUS) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PropertyExtraUSFieldTypeMap returns key value of field name and key
var PropertyExtraUSFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                 Tt.Unsigned,
	`propertyKey`:        Tt.String,
	`countyUrl`:          Tt.String,
	`countyIsActive`:     Tt.Boolean,
	`zoneDataInfo`:       Tt.String,
	`taxInfo`:            Tt.String,
	`historyTaxInfo`:     Tt.String,
	`amenitySuperGroups`: Tt.String,
	`mlsDisclaimerInfo`:  Tt.String,
	`facilityInfo`:       Tt.String,
	`riskInfo`:           Tt.String,
	`mediaSourceJson`:    Tt.String,
	`taxNote`:            Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyUS DAO reader/query struct
type PropertyUS struct {
	Adapter                 *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                      uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	UniqPropKey             string      `json:"uniqPropKey" form:"uniqPropKey" query:"uniqPropKey" long:"uniqPropKey" msg:"uniqPropKey"`
	SerialNumber            string      `json:"serialNumber" form:"serialNumber" query:"serialNumber" long:"serialNumber" msg:"serialNumber"`
	SizeM2                  string      `json:"sizeM2" form:"sizeM2" query:"sizeM2" long:"sizeM2" msg:"sizeM2"`
	MainUse                 string      `json:"mainUse" form:"mainUse" query:"mainUse" long:"mainUse" msg:"mainUse"`
	MainBuildingMaterial    string      `json:"mainBuildingMaterial" form:"mainBuildingMaterial" query:"mainBuildingMaterial" long:"mainBuildingMaterial" msg:"mainBuildingMaterial"`
	ConstructCompletedDate  string      `json:"constructCompletedDate" form:"constructCompletedDate" query:"constructCompletedDate" long:"constructCompletedDate" msg:"constructCompletedDate"`
	NumberOfFloors          string      `json:"numberOfFloors" form:"numberOfFloors" query:"numberOfFloors" long:"numberOfFloors" msg:"numberOfFloors"`
	BuildingLamination      string      `json:"buildingLamination" form:"buildingLamination" query:"buildingLamination" long:"buildingLamination" msg:"buildingLamination"`
	Address                 string      `json:"address" form:"address" query:"address" long:"address" msg:"address"`
	District                string      `json:"district" form:"district" query:"district" long:"district" msg:"district"`
	Note                    string      `json:"note" form:"note" query:"note" long:"note" msg:"note"`
	Coord                   []any       `json:"coord" form:"coord" query:"coord" long:"coord" msg:"coord"`
	CreatedAt               int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy               uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt               int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy               uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt               int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	FormattedAddress        string      `json:"formattedAddress" form:"formattedAddress" query:"formattedAddress" long:"formattedAddress" msg:"formattedAddress"`
	LastPrice               string      `json:"lastPrice" form:"lastPrice" query:"lastPrice" long:"lastPrice" msg:"lastPrice"`
	PriceHistoriesSell      []any       `json:"priceHistoriesSell" form:"priceHistoriesSell" query:"priceHistoriesSell" long:"priceHistoriesSell" msg:"priceHistoriesSell"`
	PriceHistoriesRent      []any       `json:"priceHistoriesRent" form:"priceHistoriesRent" query:"priceHistoriesRent" long:"priceHistoriesRent" msg:"priceHistoriesRent"`
	Purpose                 string      `json:"purpose" form:"purpose" query:"purpose" long:"purpose" msg:"purpose"`
	HouseType               string      `json:"houseType" form:"houseType" query:"houseType" long:"houseType" msg:"houseType"`
	Images                  []any       `json:"images" form:"images" query:"images" long:"images" msg:"images"`
	Bedroom                 int64       `json:"bedroom" form:"bedroom" query:"bedroom" long:"bedroom" msg:"bedroom"`
	Bathroom                int64       `json:"bathroom" form:"bathroom" query:"bathroom" long:"bathroom" msg:"bathroom"`
	AgencyFeePercent        float64     `json:"agencyFeePercent" form:"agencyFeePercent" query:"agencyFeePercent" long:"agencyFeePercent" msg:"agencyFeePercent"`
	FloorList               []any       `json:"floorList" form:"floorList" query:"floorList" long:"floorList" msg:"floorList"`
	Version                 string      `json:"version" form:"version" query:"version" long:"version" msg:"version"`
	YearBuilt               int64       `json:"yearBuilt" form:"yearBuilt" query:"yearBuilt" long:"yearBuilt" msg:"yearBuilt"`
	YearRenovated           int64       `json:"yearRenovated" form:"yearRenovated" query:"yearRenovated" long:"yearRenovated" msg:"yearRenovated"`
	TotalSqft               float64     `json:"totalSqft" form:"totalSqft" query:"totalSqft" long:"totalSqft" msg:"totalSqft"`
	CountyName              string      `json:"countyName" form:"countyName" query:"countyName" long:"countyName" msg:"countyName"`
	Street                  string      `json:"street" form:"street" query:"street" long:"street" msg:"street"`
	City                    string      `json:"city" form:"city" query:"city" long:"city" msg:"city"`
	State                   string      `json:"state" form:"state" query:"state" long:"state" msg:"state"`
	Zip                     string      `json:"zip" form:"zip" query:"zip" long:"zip" msg:"zip"`
	PropertyLastUpdatedDate int64       `json:"propertyLastUpdatedDate" form:"propertyLastUpdatedDate" query:"propertyLastUpdatedDate" long:"propertyLastUpdatedDate" msg:"propertyLastUpdatedDate"`
}

// NewPropertyUS create new ORM reader/query object
func NewPropertyUS(adapter *Tt.Adapter) *PropertyUS {
	return &PropertyUS{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *PropertyUS) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TablePropertyUS) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *PropertyUS) SqlTableName() string { //nolint:dupl false positive
	return `"propertyUS"`
}

func (p *PropertyUS) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *PropertyUS) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{p.Id})
	if L.IsError(err, `PropertyUS.FindById failed: `+p.SpaceName()) {
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
func (p *PropertyUS) SpatialIndexCoord() string { //nolint:dupl false positive
	return `coord`
}

// UniqueIndexUniqPropKey return unique index name
func (p *PropertyUS) UniqueIndexUniqPropKey() string { //nolint:dupl false positive
	return `uniqPropKey`
}

// FindByUniqPropKey Find one by UniqPropKey
func (p *PropertyUS) FindByUniqPropKey() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexUniqPropKey(), 0, 1, tarantool.IterEq, A.X{p.UniqPropKey})
	if L.IsError(err, `PropertyUS.FindByUniqPropKey failed: `+p.SpaceName()) {
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
func (p *PropertyUS) SqlSelectAllFields() string { //nolint:dupl false positive
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
	, "priceHistoriesSell"
	, "priceHistoriesRent"
	, "purpose"
	, "houseType"
	, "images"
	, "bedroom"
	, "bathroom"
	, "agencyFeePercent"
	, "floorList"
	, "version"
	, "yearBuilt"
	, "yearRenovated"
	, "totalSqft"
	, "countyName"
	, "street"
	, "city"
	, "state"
	, "zip"
	, "propertyLastUpdatedDate"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *PropertyUS) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
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
	, "priceHistoriesSell"
	, "priceHistoriesRent"
	, "purpose"
	, "houseType"
	, "images"
	, "bedroom"
	, "bathroom"
	, "agencyFeePercent"
	, "floorList"
	, "version"
	, "yearBuilt"
	, "yearRenovated"
	, "totalSqft"
	, "countyName"
	, "street"
	, "city"
	, "state"
	, "zip"
	, "propertyLastUpdatedDate"
	`
}

// ToUpdateArray generate slice of update command
func (p *PropertyUS) ToUpdateArray() A.X { //nolint:dupl false positive
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
		A.X{`=`, 20, p.PriceHistoriesSell},
		A.X{`=`, 21, p.PriceHistoriesRent},
		A.X{`=`, 22, p.Purpose},
		A.X{`=`, 23, p.HouseType},
		A.X{`=`, 24, p.Images},
		A.X{`=`, 25, p.Bedroom},
		A.X{`=`, 26, p.Bathroom},
		A.X{`=`, 27, p.AgencyFeePercent},
		A.X{`=`, 28, p.FloorList},
		A.X{`=`, 29, p.Version},
		A.X{`=`, 30, p.YearBuilt},
		A.X{`=`, 31, p.YearRenovated},
		A.X{`=`, 32, p.TotalSqft},
		A.X{`=`, 33, p.CountyName},
		A.X{`=`, 34, p.Street},
		A.X{`=`, 35, p.City},
		A.X{`=`, 36, p.State},
		A.X{`=`, 37, p.Zip},
		A.X{`=`, 38, p.PropertyLastUpdatedDate},
	}
}

// IdxId return name of the index
func (p *PropertyUS) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *PropertyUS) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxUniqPropKey return name of the index
func (p *PropertyUS) IdxUniqPropKey() int { //nolint:dupl false positive
	return 1
}

// SqlUniqPropKey return name of the column being indexed
func (p *PropertyUS) SqlUniqPropKey() string { //nolint:dupl false positive
	return `"uniqPropKey"`
}

// IdxSerialNumber return name of the index
func (p *PropertyUS) IdxSerialNumber() int { //nolint:dupl false positive
	return 2
}

// SqlSerialNumber return name of the column being indexed
func (p *PropertyUS) SqlSerialNumber() string { //nolint:dupl false positive
	return `"serialNumber"`
}

// IdxSizeM2 return name of the index
func (p *PropertyUS) IdxSizeM2() int { //nolint:dupl false positive
	return 3
}

// SqlSizeM2 return name of the column being indexed
func (p *PropertyUS) SqlSizeM2() string { //nolint:dupl false positive
	return `"sizeM2"`
}

// IdxMainUse return name of the index
func (p *PropertyUS) IdxMainUse() int { //nolint:dupl false positive
	return 4
}

// SqlMainUse return name of the column being indexed
func (p *PropertyUS) SqlMainUse() string { //nolint:dupl false positive
	return `"mainUse"`
}

// IdxMainBuildingMaterial return name of the index
func (p *PropertyUS) IdxMainBuildingMaterial() int { //nolint:dupl false positive
	return 5
}

// SqlMainBuildingMaterial return name of the column being indexed
func (p *PropertyUS) SqlMainBuildingMaterial() string { //nolint:dupl false positive
	return `"mainBuildingMaterial"`
}

// IdxConstructCompletedDate return name of the index
func (p *PropertyUS) IdxConstructCompletedDate() int { //nolint:dupl false positive
	return 6
}

// SqlConstructCompletedDate return name of the column being indexed
func (p *PropertyUS) SqlConstructCompletedDate() string { //nolint:dupl false positive
	return `"constructCompletedDate"`
}

// IdxNumberOfFloors return name of the index
func (p *PropertyUS) IdxNumberOfFloors() int { //nolint:dupl false positive
	return 7
}

// SqlNumberOfFloors return name of the column being indexed
func (p *PropertyUS) SqlNumberOfFloors() string { //nolint:dupl false positive
	return `"numberOfFloors"`
}

// IdxBuildingLamination return name of the index
func (p *PropertyUS) IdxBuildingLamination() int { //nolint:dupl false positive
	return 8
}

// SqlBuildingLamination return name of the column being indexed
func (p *PropertyUS) SqlBuildingLamination() string { //nolint:dupl false positive
	return `"buildingLamination"`
}

// IdxAddress return name of the index
func (p *PropertyUS) IdxAddress() int { //nolint:dupl false positive
	return 9
}

// SqlAddress return name of the column being indexed
func (p *PropertyUS) SqlAddress() string { //nolint:dupl false positive
	return `"address"`
}

// IdxDistrict return name of the index
func (p *PropertyUS) IdxDistrict() int { //nolint:dupl false positive
	return 10
}

// SqlDistrict return name of the column being indexed
func (p *PropertyUS) SqlDistrict() string { //nolint:dupl false positive
	return `"district"`
}

// IdxNote return name of the index
func (p *PropertyUS) IdxNote() int { //nolint:dupl false positive
	return 11
}

// SqlNote return name of the column being indexed
func (p *PropertyUS) SqlNote() string { //nolint:dupl false positive
	return `"note"`
}

// IdxCoord return name of the index
func (p *PropertyUS) IdxCoord() int { //nolint:dupl false positive
	return 12
}

// SqlCoord return name of the column being indexed
func (p *PropertyUS) SqlCoord() string { //nolint:dupl false positive
	return `"coord"`
}

// IdxCreatedAt return name of the index
func (p *PropertyUS) IdxCreatedAt() int { //nolint:dupl false positive
	return 13
}

// SqlCreatedAt return name of the column being indexed
func (p *PropertyUS) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *PropertyUS) IdxCreatedBy() int { //nolint:dupl false positive
	return 14
}

// SqlCreatedBy return name of the column being indexed
func (p *PropertyUS) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *PropertyUS) IdxUpdatedAt() int { //nolint:dupl false positive
	return 15
}

// SqlUpdatedAt return name of the column being indexed
func (p *PropertyUS) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *PropertyUS) IdxUpdatedBy() int { //nolint:dupl false positive
	return 16
}

// SqlUpdatedBy return name of the column being indexed
func (p *PropertyUS) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *PropertyUS) IdxDeletedAt() int { //nolint:dupl false positive
	return 17
}

// SqlDeletedAt return name of the column being indexed
func (p *PropertyUS) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxFormattedAddress return name of the index
func (p *PropertyUS) IdxFormattedAddress() int { //nolint:dupl false positive
	return 18
}

// SqlFormattedAddress return name of the column being indexed
func (p *PropertyUS) SqlFormattedAddress() string { //nolint:dupl false positive
	return `"formattedAddress"`
}

// IdxLastPrice return name of the index
func (p *PropertyUS) IdxLastPrice() int { //nolint:dupl false positive
	return 19
}

// SqlLastPrice return name of the column being indexed
func (p *PropertyUS) SqlLastPrice() string { //nolint:dupl false positive
	return `"lastPrice"`
}

// IdxPriceHistoriesSell return name of the index
func (p *PropertyUS) IdxPriceHistoriesSell() int { //nolint:dupl false positive
	return 20
}

// SqlPriceHistoriesSell return name of the column being indexed
func (p *PropertyUS) SqlPriceHistoriesSell() string { //nolint:dupl false positive
	return `"priceHistoriesSell"`
}

// IdxPriceHistoriesRent return name of the index
func (p *PropertyUS) IdxPriceHistoriesRent() int { //nolint:dupl false positive
	return 21
}

// SqlPriceHistoriesRent return name of the column being indexed
func (p *PropertyUS) SqlPriceHistoriesRent() string { //nolint:dupl false positive
	return `"priceHistoriesRent"`
}

// IdxPurpose return name of the index
func (p *PropertyUS) IdxPurpose() int { //nolint:dupl false positive
	return 22
}

// SqlPurpose return name of the column being indexed
func (p *PropertyUS) SqlPurpose() string { //nolint:dupl false positive
	return `"purpose"`
}

// IdxHouseType return name of the index
func (p *PropertyUS) IdxHouseType() int { //nolint:dupl false positive
	return 23
}

// SqlHouseType return name of the column being indexed
func (p *PropertyUS) SqlHouseType() string { //nolint:dupl false positive
	return `"houseType"`
}

// IdxImages return name of the index
func (p *PropertyUS) IdxImages() int { //nolint:dupl false positive
	return 24
}

// SqlImages return name of the column being indexed
func (p *PropertyUS) SqlImages() string { //nolint:dupl false positive
	return `"images"`
}

// IdxBedroom return name of the index
func (p *PropertyUS) IdxBedroom() int { //nolint:dupl false positive
	return 25
}

// SqlBedroom return name of the column being indexed
func (p *PropertyUS) SqlBedroom() string { //nolint:dupl false positive
	return `"bedroom"`
}

// IdxBathroom return name of the index
func (p *PropertyUS) IdxBathroom() int { //nolint:dupl false positive
	return 26
}

// SqlBathroom return name of the column being indexed
func (p *PropertyUS) SqlBathroom() string { //nolint:dupl false positive
	return `"bathroom"`
}

// IdxAgencyFeePercent return name of the index
func (p *PropertyUS) IdxAgencyFeePercent() int { //nolint:dupl false positive
	return 27
}

// SqlAgencyFeePercent return name of the column being indexed
func (p *PropertyUS) SqlAgencyFeePercent() string { //nolint:dupl false positive
	return `"agencyFeePercent"`
}

// IdxFloorList return name of the index
func (p *PropertyUS) IdxFloorList() int { //nolint:dupl false positive
	return 28
}

// SqlFloorList return name of the column being indexed
func (p *PropertyUS) SqlFloorList() string { //nolint:dupl false positive
	return `"floorList"`
}

// IdxVersion return name of the index
func (p *PropertyUS) IdxVersion() int { //nolint:dupl false positive
	return 29
}

// SqlVersion return name of the column being indexed
func (p *PropertyUS) SqlVersion() string { //nolint:dupl false positive
	return `"version"`
}

// IdxYearBuilt return name of the index
func (p *PropertyUS) IdxYearBuilt() int { //nolint:dupl false positive
	return 30
}

// SqlYearBuilt return name of the column being indexed
func (p *PropertyUS) SqlYearBuilt() string { //nolint:dupl false positive
	return `"yearBuilt"`
}

// IdxYearRenovated return name of the index
func (p *PropertyUS) IdxYearRenovated() int { //nolint:dupl false positive
	return 31
}

// SqlYearRenovated return name of the column being indexed
func (p *PropertyUS) SqlYearRenovated() string { //nolint:dupl false positive
	return `"yearRenovated"`
}

// IdxTotalSqft return name of the index
func (p *PropertyUS) IdxTotalSqft() int { //nolint:dupl false positive
	return 32
}

// SqlTotalSqft return name of the column being indexed
func (p *PropertyUS) SqlTotalSqft() string { //nolint:dupl false positive
	return `"totalSqft"`
}

// IdxCountyName return name of the index
func (p *PropertyUS) IdxCountyName() int { //nolint:dupl false positive
	return 33
}

// SqlCountyName return name of the column being indexed
func (p *PropertyUS) SqlCountyName() string { //nolint:dupl false positive
	return `"countyName"`
}

// IdxStreet return name of the index
func (p *PropertyUS) IdxStreet() int { //nolint:dupl false positive
	return 34
}

// SqlStreet return name of the column being indexed
func (p *PropertyUS) SqlStreet() string { //nolint:dupl false positive
	return `"street"`
}

// IdxCity return name of the index
func (p *PropertyUS) IdxCity() int { //nolint:dupl false positive
	return 35
}

// SqlCity return name of the column being indexed
func (p *PropertyUS) SqlCity() string { //nolint:dupl false positive
	return `"city"`
}

// IdxState return name of the index
func (p *PropertyUS) IdxState() int { //nolint:dupl false positive
	return 36
}

// SqlState return name of the column being indexed
func (p *PropertyUS) SqlState() string { //nolint:dupl false positive
	return `"state"`
}

// IdxZip return name of the index
func (p *PropertyUS) IdxZip() int { //nolint:dupl false positive
	return 37
}

// SqlZip return name of the column being indexed
func (p *PropertyUS) SqlZip() string { //nolint:dupl false positive
	return `"zip"`
}

// IdxPropertyLastUpdatedDate return name of the index
func (p *PropertyUS) IdxPropertyLastUpdatedDate() int { //nolint:dupl false positive
	return 38
}

// SqlPropertyLastUpdatedDate return name of the column being indexed
func (p *PropertyUS) SqlPropertyLastUpdatedDate() string { //nolint:dupl false positive
	return `"propertyLastUpdatedDate"`
}

// ToArray receiver fields to slice
func (p *PropertyUS) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.UniqPropKey,             // 1
		p.SerialNumber,            // 2
		p.SizeM2,                  // 3
		p.MainUse,                 // 4
		p.MainBuildingMaterial,    // 5
		p.ConstructCompletedDate,  // 6
		p.NumberOfFloors,          // 7
		p.BuildingLamination,      // 8
		p.Address,                 // 9
		p.District,                // 10
		p.Note,                    // 11
		p.Coord,                   // 12
		p.CreatedAt,               // 13
		p.CreatedBy,               // 14
		p.UpdatedAt,               // 15
		p.UpdatedBy,               // 16
		p.DeletedAt,               // 17
		p.FormattedAddress,        // 18
		p.LastPrice,               // 19
		p.PriceHistoriesSell,      // 20
		p.PriceHistoriesRent,      // 21
		p.Purpose,                 // 22
		p.HouseType,               // 23
		p.Images,                  // 24
		p.Bedroom,                 // 25
		p.Bathroom,                // 26
		p.AgencyFeePercent,        // 27
		p.FloorList,               // 28
		p.Version,                 // 29
		p.YearBuilt,               // 30
		p.YearRenovated,           // 31
		p.TotalSqft,               // 32
		p.CountyName,              // 33
		p.Street,                  // 34
		p.City,                    // 35
		p.State,                   // 36
		p.Zip,                     // 37
		p.PropertyLastUpdatedDate, // 38
	}
}

// FromArray convert slice to receiver fields
func (p *PropertyUS) FromArray(a A.X) *PropertyUS { //nolint:dupl false positive
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
	p.PriceHistoriesSell = X.ToArr(a[20])
	p.PriceHistoriesRent = X.ToArr(a[21])
	p.Purpose = X.ToS(a[22])
	p.HouseType = X.ToS(a[23])
	p.Images = X.ToArr(a[24])
	p.Bedroom = X.ToI(a[25])
	p.Bathroom = X.ToI(a[26])
	p.AgencyFeePercent = X.ToF(a[27])
	p.FloorList = X.ToArr(a[28])
	p.Version = X.ToS(a[29])
	p.YearBuilt = X.ToI(a[30])
	p.YearRenovated = X.ToI(a[31])
	p.TotalSqft = X.ToF(a[32])
	p.CountyName = X.ToS(a[33])
	p.Street = X.ToS(a[34])
	p.City = X.ToS(a[35])
	p.State = X.ToS(a[36])
	p.Zip = X.ToS(a[37])
	p.PropertyLastUpdatedDate = X.ToI(a[38])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *PropertyUS) FromUncensoredArray(a A.X) *PropertyUS { //nolint:dupl false positive
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
	p.PriceHistoriesSell = X.ToArr(a[20])
	p.PriceHistoriesRent = X.ToArr(a[21])
	p.Purpose = X.ToS(a[22])
	p.HouseType = X.ToS(a[23])
	p.Images = X.ToArr(a[24])
	p.Bedroom = X.ToI(a[25])
	p.Bathroom = X.ToI(a[26])
	p.AgencyFeePercent = X.ToF(a[27])
	p.FloorList = X.ToArr(a[28])
	p.Version = X.ToS(a[29])
	p.YearBuilt = X.ToI(a[30])
	p.YearRenovated = X.ToI(a[31])
	p.TotalSqft = X.ToF(a[32])
	p.CountyName = X.ToS(a[33])
	p.Street = X.ToS(a[34])
	p.City = X.ToS(a[35])
	p.State = X.ToS(a[36])
	p.Zip = X.ToS(a[37])
	p.PropertyLastUpdatedDate = X.ToI(a[38])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *PropertyUS) FindOffsetLimit(offset, limit uint32, idx string) []PropertyUS { //nolint:dupl false positive
	var rows []PropertyUS
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropertyUS.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := PropertyUS{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *PropertyUS) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `PropertyUS.FindOffsetLimit failed: `+p.SpaceName()) {
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
func (p *PropertyUS) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PropertyUSFieldTypeMap returns key value of field name and key
var PropertyUSFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                      Tt.Unsigned,
	`uniqPropKey`:             Tt.String,
	`serialNumber`:            Tt.String,
	`sizeM2`:                  Tt.String,
	`mainUse`:                 Tt.String,
	`mainBuildingMaterial`:    Tt.String,
	`constructCompletedDate`:  Tt.String,
	`numberOfFloors`:          Tt.String,
	`buildingLamination`:      Tt.String,
	`address`:                 Tt.String,
	`district`:                Tt.String,
	`note`:                    Tt.String,
	`coord`:                   Tt.Array,
	`createdAt`:               Tt.Integer,
	`createdBy`:               Tt.Unsigned,
	`updatedAt`:               Tt.Integer,
	`updatedBy`:               Tt.Unsigned,
	`deletedAt`:               Tt.Integer,
	`formattedAddress`:        Tt.String,
	`lastPrice`:               Tt.String,
	`priceHistoriesSell`:      Tt.Array,
	`priceHistoriesRent`:      Tt.Array,
	`purpose`:                 Tt.String,
	`houseType`:               Tt.String,
	`images`:                  Tt.Array,
	`bedroom`:                 Tt.Integer,
	`bathroom`:                Tt.Integer,
	`agencyFeePercent`:        Tt.Double,
	`floorList`:               Tt.Array,
	`version`:                 Tt.String,
	`yearBuilt`:               Tt.Integer,
	`yearRenovated`:           Tt.Integer,
	`totalSqft`:               Tt.Double,
	`countyName`:              Tt.String,
	`street`:                  Tt.String,
	`city`:                    Tt.String,
	`state`:                   Tt.String,
	`zip`:                     Tt.String,
	`propertyLastUpdatedDate`: Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyHistory DAO reader/query struct
type PropertyHistory struct {
	Adapter                *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                     uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	PropertyKey            string      `json:"propertyKey" form:"propertyKey" query:"propertyKey" long:"propertyKey" msg:"propertyKey"`
	TransactionKey         string      `json:"transactionKey" form:"transactionKey" query:"transactionKey" long:"transactionKey" msg:"transactionKey"`
	TransactionType        string      `json:"transactionType" form:"transactionType" query:"transactionType" long:"transactionType" msg:"transactionType"`
	TransactionSign        string      `json:"transactionSign" form:"transactionSign" query:"transactionSign" long:"transactionSign" msg:"transactionSign"`
	TransactionTime        string      `json:"transactionTime" form:"transactionTime" query:"transactionTime" long:"transactionTime" msg:"transactionTime"`
	TransactionDateNormal  string      `json:"transactionDateNormal" form:"transactionDateNormal" query:"transactionDateNormal" long:"transactionDateNormal" msg:"transactionDateNormal"`
	TransactionNumber      string      `json:"transactionNumber" form:"transactionNumber" query:"transactionNumber" long:"transactionNumber" msg:"transactionNumber"`
	PriceNtd               int64       `json:"priceNtd" form:"priceNtd" query:"priceNtd" long:"priceNtd" msg:"priceNtd"`
	PricePerUnit           int64       `json:"pricePerUnit" form:"pricePerUnit" query:"pricePerUnit" long:"pricePerUnit" msg:"pricePerUnit"`
	Price                  int64       `json:"price" form:"price" query:"price" long:"price" msg:"price"`
	Address                string      `json:"address" form:"address" query:"address" long:"address" msg:"address"`
	District               string      `json:"district" form:"district" query:"district" long:"district" msg:"district"`
	Note                   string      `json:"note" form:"note" query:"note" long:"note" msg:"note"`
	CreatedAt              int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy              uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt              int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy              uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt              int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	SerialNumber           string      `json:"serialNumber" form:"serialNumber" query:"serialNumber" long:"serialNumber" msg:"serialNumber"`
	TransactionDescription string      `json:"transactionDescription" form:"transactionDescription" query:"transactionDescription" long:"transactionDescription" msg:"transactionDescription"`
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
	, "transactionDescription"
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
	, "transactionDescription"
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
		A.X{`=`, 20, p.TransactionDescription},
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

// IdxTransactionDescription return name of the index
func (p *PropertyHistory) IdxTransactionDescription() int { //nolint:dupl false positive
	return 20
}

// SqlTransactionDescription return name of the column being indexed
func (p *PropertyHistory) SqlTransactionDescription() string { //nolint:dupl false positive
	return `"transactionDescription"`
}

// ToArray receiver fields to slice
func (p *PropertyHistory) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.PropertyKey,            // 1
		p.TransactionKey,         // 2
		p.TransactionType,        // 3
		p.TransactionSign,        // 4
		p.TransactionTime,        // 5
		p.TransactionDateNormal,  // 6
		p.TransactionNumber,      // 7
		p.PriceNtd,               // 8
		p.PricePerUnit,           // 9
		p.Price,                  // 10
		p.Address,                // 11
		p.District,               // 12
		p.Note,                   // 13
		p.CreatedAt,              // 14
		p.CreatedBy,              // 15
		p.UpdatedAt,              // 16
		p.UpdatedBy,              // 17
		p.DeletedAt,              // 18
		p.SerialNumber,           // 19
		p.TransactionDescription, // 20
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
	p.TransactionDescription = X.ToS(a[20])
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
	p.TransactionDescription = X.ToS(a[20])
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
	`id`:                     Tt.Unsigned,
	`propertyKey`:            Tt.String,
	`transactionKey`:         Tt.String,
	`transactionType`:        Tt.String,
	`transactionSign`:        Tt.String,
	`transactionTime`:        Tt.String,
	`transactionDateNormal`:  Tt.String,
	`transactionNumber`:      Tt.String,
	`priceNtd`:               Tt.Integer,
	`pricePerUnit`:           Tt.Integer,
	`price`:                  Tt.Integer,
	`address`:                Tt.String,
	`district`:               Tt.String,
	`note`:                   Tt.String,
	`createdAt`:              Tt.Integer,
	`createdBy`:              Tt.Unsigned,
	`updatedAt`:              Tt.Integer,
	`updatedBy`:              Tt.Unsigned,
	`deletedAt`:              Tt.Integer,
	`serialNumber`:           Tt.String,
	`transactionDescription`: Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// UserPropLikes DAO reader/query struct
type UserPropLikes struct {
	Adapter   *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	PropId    uint64      `json:"propId,string" form:"propId" query:"propId" long:"propId" msg:"propId"`
	UserId    uint64      `json:"userId,string" form:"userId" query:"userId" long:"userId" msg:"userId"`
	CreatedAt int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
}

// NewUserPropLikes create new ORM reader/query object
func NewUserPropLikes(adapter *Tt.Adapter) *UserPropLikes {
	return &UserPropLikes{Adapter: adapter}
}

// SpaceName returns full package and table name
func (u *UserPropLikes) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableUserPropLikes) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (u *UserPropLikes) SqlTableName() string { //nolint:dupl false positive
	return `"userPropLikes"`
}

// UniqueIndexUserIdPropId return unique index name
func (u *UserPropLikes) UniqueIndexUserIdPropId() string { //nolint:dupl false positive
	return `userId__propId`
}

// FindByUserIdPropId Find one by UserIdPropId
func (u *UserPropLikes) FindByUserIdPropId() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexUserIdPropId(), 0, 1, tarantool.IterEq, A.X{u.UserId, u.PropId})
	if L.IsError(err, `UserPropLikes.FindByUserIdPropId failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (u *UserPropLikes) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "propId"
	, "userId"
	, "createdAt"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (u *UserPropLikes) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "propId"
	, "userId"
	, "createdAt"
	`
}

// ToUpdateArray generate slice of update command
func (u *UserPropLikes) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, u.PropId},
		A.X{`=`, 1, u.UserId},
		A.X{`=`, 2, u.CreatedAt},
	}
}

// IdxPropId return name of the index
func (u *UserPropLikes) IdxPropId() int { //nolint:dupl false positive
	return 0
}

// SqlPropId return name of the column being indexed
func (u *UserPropLikes) SqlPropId() string { //nolint:dupl false positive
	return `"propId"`
}

// IdxUserId return name of the index
func (u *UserPropLikes) IdxUserId() int { //nolint:dupl false positive
	return 1
}

// SqlUserId return name of the column being indexed
func (u *UserPropLikes) SqlUserId() string { //nolint:dupl false positive
	return `"userId"`
}

// IdxCreatedAt return name of the index
func (u *UserPropLikes) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (u *UserPropLikes) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// ToArray receiver fields to slice
func (u *UserPropLikes) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		u.PropId,    // 0
		u.UserId,    // 1
		u.CreatedAt, // 2
	}
}

// FromArray convert slice to receiver fields
func (u *UserPropLikes) FromArray(a A.X) *UserPropLikes { //nolint:dupl false positive
	u.PropId = X.ToU(a[0])
	u.UserId = X.ToU(a[1])
	u.CreatedAt = X.ToI(a[2])
	return u
}

// FromUncensoredArray convert slice to receiver fields
func (u *UserPropLikes) FromUncensoredArray(a A.X) *UserPropLikes { //nolint:dupl false positive
	u.PropId = X.ToU(a[0])
	u.UserId = X.ToU(a[1])
	u.CreatedAt = X.ToI(a[2])
	return u
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (u *UserPropLikes) FindOffsetLimit(offset, limit uint32, idx string) []UserPropLikes { //nolint:dupl false positive
	var rows []UserPropLikes
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `UserPropLikes.FindOffsetLimit failed: `+u.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := UserPropLikes{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (u *UserPropLikes) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `UserPropLikes.FindOffsetLimit failed: `+u.SpaceName()) {
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
func (u *UserPropLikes) Total() int64 { //nolint:dupl false positive
	rows := u.Adapter.CallBoxSpace(u.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// UserPropLikesFieldTypeMap returns key value of field name and key
var UserPropLikesFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`propId`:    Tt.Unsigned,
	`userId`:    Tt.Unsigned,
	`createdAt`: Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
