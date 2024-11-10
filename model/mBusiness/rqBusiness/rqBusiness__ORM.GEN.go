package rqBusiness

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mBusiness"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Sales DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqBusiness__ORM.GEN.go
type Sales struct {
	Adapter         *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id              uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	PropertyId      uint64      `json:"propertyId,string" form:"propertyId" query:"propertyId" long:"propertyId" msg:"propertyId"`
	RealtorId       uint64      `json:"realtorId,string" form:"realtorId" query:"realtorId" long:"realtorId" msg:"realtorId"`
	PropertyCountry string      `json:"propertyCountry" form:"propertyCountry" query:"propertyCountry" long:"propertyCountry" msg:"propertyCountry"`
	BuyerId         uint64      `json:"buyerId,string" form:"buyerId" query:"buyerId" long:"buyerId" msg:"buyerId"`
	Price           string      `json:"price" form:"price" query:"price" long:"price" msg:"price"`
	BuyerEmail      string      `json:"buyerEmail" form:"buyerEmail" query:"buyerEmail" long:"buyerEmail" msg:"buyerEmail"`
	EmailNotFound   string      `json:"emailNotFound" form:"emailNotFound" query:"emailNotFound" long:"emailNotFound" msg:"emailNotFound"`
	SalesDate       string      `json:"salesDate" form:"salesDate" query:"salesDate" long:"salesDate" msg:"salesDate"`
	CreatedAt       int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy       uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt       int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy       uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt       int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
}

// NewSales create new ORM reader/query object
func NewSales(adapter *Tt.Adapter) *Sales {
	return &Sales{Adapter: adapter}
}

// SpaceName returns full package and table name
func (s *Sales) SpaceName() string { //nolint:dupl false positive
	return string(mBusiness.TableSales) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (s *Sales) SqlTableName() string { //nolint:dupl false positive
	return `"sales"`
}

func (s *Sales) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (s *Sales) FindById() bool { //nolint:dupl false positive
	res, err := s.Adapter.Select(s.SpaceName(), s.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{s.Id})
	if L.IsError(err, `Sales.FindById failed: `+s.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		s.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexPropertyIdSalesDate return unique index name
func (s *Sales) UniqueIndexPropertyIdSalesDate() string { //nolint:dupl false positive
	return `propertyId__salesDate`
}

// FindByPropertyIdSalesDate Find one by PropertyIdSalesDate
func (s *Sales) FindByPropertyIdSalesDate() bool { //nolint:dupl false positive
	res, err := s.Adapter.Select(s.SpaceName(), s.UniqueIndexPropertyIdSalesDate(), 0, 1, tarantool.IterEq, A.X{s.PropertyId, s.SalesDate})
	if L.IsError(err, `Sales.FindByPropertyIdSalesDate failed: `+s.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		s.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (s *Sales) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "propertyId"
	, "realtorId"
	, "propertyCountry"
	, "buyerId"
	, "price"
	, "buyerEmail"
	, "emailNotFound"
	, "salesDate"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (s *Sales) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "propertyId"
	, "realtorId"
	, "propertyCountry"
	, "buyerId"
	, "price"
	, "buyerEmail"
	, "emailNotFound"
	, "salesDate"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// ToUpdateArray generate slice of update command
func (s *Sales) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, s.Id},
		A.X{`=`, 1, s.PropertyId},
		A.X{`=`, 2, s.RealtorId},
		A.X{`=`, 3, s.PropertyCountry},
		A.X{`=`, 4, s.BuyerId},
		A.X{`=`, 5, s.Price},
		A.X{`=`, 6, s.BuyerEmail},
		A.X{`=`, 7, s.EmailNotFound},
		A.X{`=`, 8, s.SalesDate},
		A.X{`=`, 9, s.CreatedAt},
		A.X{`=`, 10, s.CreatedBy},
		A.X{`=`, 11, s.UpdatedAt},
		A.X{`=`, 12, s.UpdatedBy},
		A.X{`=`, 13, s.DeletedAt},
	}
}

// IdxId return name of the index
func (s *Sales) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (s *Sales) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxPropertyId return name of the index
func (s *Sales) IdxPropertyId() int { //nolint:dupl false positive
	return 1
}

// SqlPropertyId return name of the column being indexed
func (s *Sales) SqlPropertyId() string { //nolint:dupl false positive
	return `"propertyId"`
}

// IdxRealtorId return name of the index
func (s *Sales) IdxRealtorId() int { //nolint:dupl false positive
	return 2
}

// SqlRealtorId return name of the column being indexed
func (s *Sales) SqlRealtorId() string { //nolint:dupl false positive
	return `"realtorId"`
}

// IdxPropertyCountry return name of the index
func (s *Sales) IdxPropertyCountry() int { //nolint:dupl false positive
	return 3
}

// SqlPropertyCountry return name of the column being indexed
func (s *Sales) SqlPropertyCountry() string { //nolint:dupl false positive
	return `"propertyCountry"`
}

// IdxBuyerId return name of the index
func (s *Sales) IdxBuyerId() int { //nolint:dupl false positive
	return 4
}

// SqlBuyerId return name of the column being indexed
func (s *Sales) SqlBuyerId() string { //nolint:dupl false positive
	return `"buyerId"`
}

// IdxPrice return name of the index
func (s *Sales) IdxPrice() int { //nolint:dupl false positive
	return 5
}

// SqlPrice return name of the column being indexed
func (s *Sales) SqlPrice() string { //nolint:dupl false positive
	return `"price"`
}

// IdxBuyerEmail return name of the index
func (s *Sales) IdxBuyerEmail() int { //nolint:dupl false positive
	return 6
}

// SqlBuyerEmail return name of the column being indexed
func (s *Sales) SqlBuyerEmail() string { //nolint:dupl false positive
	return `"buyerEmail"`
}

// IdxEmailNotFound return name of the index
func (s *Sales) IdxEmailNotFound() int { //nolint:dupl false positive
	return 7
}

// SqlEmailNotFound return name of the column being indexed
func (s *Sales) SqlEmailNotFound() string { //nolint:dupl false positive
	return `"emailNotFound"`
}

// IdxSalesDate return name of the index
func (s *Sales) IdxSalesDate() int { //nolint:dupl false positive
	return 8
}

// SqlSalesDate return name of the column being indexed
func (s *Sales) SqlSalesDate() string { //nolint:dupl false positive
	return `"salesDate"`
}

// IdxCreatedAt return name of the index
func (s *Sales) IdxCreatedAt() int { //nolint:dupl false positive
	return 9
}

// SqlCreatedAt return name of the column being indexed
func (s *Sales) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (s *Sales) IdxCreatedBy() int { //nolint:dupl false positive
	return 10
}

// SqlCreatedBy return name of the column being indexed
func (s *Sales) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (s *Sales) IdxUpdatedAt() int { //nolint:dupl false positive
	return 11
}

// SqlUpdatedAt return name of the column being indexed
func (s *Sales) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (s *Sales) IdxUpdatedBy() int { //nolint:dupl false positive
	return 12
}

// SqlUpdatedBy return name of the column being indexed
func (s *Sales) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (s *Sales) IdxDeletedAt() int { //nolint:dupl false positive
	return 13
}

// SqlDeletedAt return name of the column being indexed
func (s *Sales) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// ToArray receiver fields to slice
func (s *Sales) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if s.Id != 0 {
		id = s.Id
	}
	return A.X{
		id,
		s.PropertyId,      // 1
		s.RealtorId,       // 2
		s.PropertyCountry, // 3
		s.BuyerId,         // 4
		s.Price,           // 5
		s.BuyerEmail,      // 6
		s.EmailNotFound,   // 7
		s.SalesDate,       // 8
		s.CreatedAt,       // 9
		s.CreatedBy,       // 10
		s.UpdatedAt,       // 11
		s.UpdatedBy,       // 12
		s.DeletedAt,       // 13
	}
}

// FromArray convert slice to receiver fields
func (s *Sales) FromArray(a A.X) *Sales { //nolint:dupl false positive
	s.Id = X.ToU(a[0])
	s.PropertyId = X.ToU(a[1])
	s.RealtorId = X.ToU(a[2])
	s.PropertyCountry = X.ToS(a[3])
	s.BuyerId = X.ToU(a[4])
	s.Price = X.ToS(a[5])
	s.BuyerEmail = X.ToS(a[6])
	s.EmailNotFound = X.ToS(a[7])
	s.SalesDate = X.ToS(a[8])
	s.CreatedAt = X.ToI(a[9])
	s.CreatedBy = X.ToU(a[10])
	s.UpdatedAt = X.ToI(a[11])
	s.UpdatedBy = X.ToU(a[12])
	s.DeletedAt = X.ToI(a[13])
	return s
}

// FromUncensoredArray convert slice to receiver fields
func (s *Sales) FromUncensoredArray(a A.X) *Sales { //nolint:dupl false positive
	s.Id = X.ToU(a[0])
	s.PropertyId = X.ToU(a[1])
	s.RealtorId = X.ToU(a[2])
	s.PropertyCountry = X.ToS(a[3])
	s.BuyerId = X.ToU(a[4])
	s.Price = X.ToS(a[5])
	s.BuyerEmail = X.ToS(a[6])
	s.EmailNotFound = X.ToS(a[7])
	s.SalesDate = X.ToS(a[8])
	s.CreatedAt = X.ToI(a[9])
	s.CreatedBy = X.ToU(a[10])
	s.UpdatedAt = X.ToI(a[11])
	s.UpdatedBy = X.ToU(a[12])
	s.DeletedAt = X.ToI(a[13])
	return s
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (s *Sales) FindOffsetLimit(offset, limit uint32, idx string) []Sales { //nolint:dupl false positive
	var rows []Sales
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sales.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Sales{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (s *Sales) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sales.FindOffsetLimit failed: `+s.SpaceName()) {
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
func (s *Sales) Total() int64 { //nolint:dupl false positive
	rows := s.Adapter.CallBoxSpace(s.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// SalesFieldTypeMap returns key value of field name and key
var SalesFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:              Tt.Unsigned,
	`propertyId`:      Tt.Unsigned,
	`realtorId`:       Tt.Unsigned,
	`propertyCountry`: Tt.String,
	`buyerId`:         Tt.Unsigned,
	`price`:           Tt.String,
	`buyerEmail`:      Tt.String,
	`emailNotFound`:   Tt.String,
	`salesDate`:       Tt.String,
	`createdAt`:       Tt.Integer,
	`createdBy`:       Tt.Unsigned,
	`updatedAt`:       Tt.Integer,
	`updatedBy`:       Tt.Unsigned,
	`deletedAt`:       Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
