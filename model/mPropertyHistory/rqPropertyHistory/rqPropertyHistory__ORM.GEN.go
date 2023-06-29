package rqPropertyHistory

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mPropertyHistory"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqPropertyHistory__ORM.GEN.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type rqPropertyHistory__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type rqPropertyHistory__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type rqPropertyHistory__ORM.GEN.go
// go:generate msgp -tests=false -file rqPropertyHistory__ORM.GEN.go -o rqPropertyHistory__MSG.GEN.go

// PropertyHistory DAO reader/query struct
type PropertyHistory struct {
	Adapter               *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                    uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	PropertyKey           string      `json:"propertyKey" form:"propertyKey" query:"propertyKey" long:"propertyKey" msg:"propertyKey"`
	TransactionKey        string      `json:"transactionKey" form:"transactionKey" query:"transactionKey" long:"transactionKey" msg:"transactionKey"`
	TransactionType       string      `json:"transactionType" form:"transactionType" query:"transactionType" long:"transactionType" msg:"transactionType"`
	TransactionSign       string      `json:"transactionSign" form:"transactionSign" query:"transactionSign" long:"transactionSign" msg:"transactionSign"`
	TransactionTime       string      `json:"transactionTime" form:"transactionTime" query:"transactionTime" long:"transactionTime" msg:"transactionTime"`
	TransactionDateNormal string      `json:"transactionDateNormal" form:"transactionDateNormal" query:"transactionDateNormal" long:"transactionDateNormal" msg:"transactionDateNormal"`
	TransactionNumber     string      `json:"transactionNumber" form:"transactionNumber" query:"transactionNumber" long:"transactionNumber" msg:"transactionNumber"`
	PriceNtd              int64       `json:"priceNtd" form:"priceNtd" query:"priceNtd" long:"priceNtd" msg:"priceNtd"`
	PricePerUnit          int64       `json:"pricePerUnit" form:"pricePerUnit" query:"pricePerUnit" long:"pricePerUnit" msg:"pricePerUnit"`
	Price                 int64       `json:"price" form:"price" query:"price" long:"price" msg:"price"`
	Address               string      `json:"address" form:"address" query:"address" long:"address" msg:"address"`
	District              string      `json:"district" form:"district" query:"district" long:"district" msg:"district"`
	Note                  string      `json:"note" form:"note" query:"note" long:"note" msg:"note"`
	CreatedAt             int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy             uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt             int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy             uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt             int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
}

// NewPropertyHistory create new ORM reader/query object
func NewPropertyHistory(adapter *Tt.Adapter) *PropertyHistory {
	return &PropertyHistory{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *PropertyHistory) SpaceName() string { //nolint:dupl false positive
	return string(mPropertyHistory.TablePropertyHistory) // casting required to string from Tt.TableName
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
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
