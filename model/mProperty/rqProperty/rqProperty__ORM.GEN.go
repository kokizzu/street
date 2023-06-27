package rqProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`street/model/mProperty`

	`github.com/tarantool/go-tarantool`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Tt`
	`github.com/kokizzu/gotro/L`
	`github.com/kokizzu/gotro/X`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type rqProperty__ORM.GEN.go
// go:generate msgp -tests=false -file rqProperty__ORM.GEN.go -o rqProperty__MSG.GEN.go

// Property DAO reader/query struct
type Property struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
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

// UniqueIndexUniqPropKey return unique index name
func (p *Property) UniqueIndexUniqPropKey() string { //nolint:dupl false positive
	return `UniqPropKey`
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
	, "UniqPropKey"
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
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Property) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "UniqPropKey"
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
	return `"UniqPropKey"`
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
	rows := p.Adapter.CallBoxSpace(p.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PropertyFieldTypeMap returns key value of field name and key
var PropertyFieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
	`id`:                     Tt.Unsigned,
	`UniqPropKey`:            Tt.String,
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
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

