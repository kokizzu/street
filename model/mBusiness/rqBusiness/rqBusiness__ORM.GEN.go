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

// Revenue DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqBusiness__ORM.GEN.go
type Revenue struct {
	Adapter         *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id              uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	RealtorId       uint64      `json:"realtorId,string" form:"realtorId" query:"realtorId" long:"realtorId" msg:"realtorId"`
	PropertyId      uint64      `json:"propertyId,string" form:"propertyId" query:"propertyId" long:"propertyId" msg:"propertyId"`
	PropertyBought  int64       `json:"propertyBought" form:"propertyBought" query:"propertyBought" long:"propertyBought" msg:"propertyBought"`
	PropertyCountry string      `json:"propertyCountry" form:"propertyCountry" query:"propertyCountry" long:"propertyCountry" msg:"propertyCountry"`
	BuyerEmail      string      `json:"buyerEmail" form:"buyerEmail" query:"buyerEmail" long:"buyerEmail" msg:"buyerEmail"`
	CreatedAt       int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy       uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt       int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy       uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt       int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
}

// NewRevenue create new ORM reader/query object
func NewRevenue(adapter *Tt.Adapter) *Revenue {
	return &Revenue{Adapter: adapter}
}

// SpaceName returns full package and table name
func (r *Revenue) SpaceName() string { //nolint:dupl false positive
	return string(mBusiness.TableRevenue) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (r *Revenue) SqlTableName() string { //nolint:dupl false positive
	return `"revenue"`
}

// UniqueIndexRealtorIdPropertyId return unique index name
func (r *Revenue) UniqueIndexRealtorIdPropertyId() string { //nolint:dupl false positive
	return `realtorId__propertyId`
}

// FindByRealtorIdPropertyId Find one by RealtorIdPropertyId
func (r *Revenue) FindByRealtorIdPropertyId() bool { //nolint:dupl false positive
	res, err := r.Adapter.Select(r.SpaceName(), r.UniqueIndexRealtorIdPropertyId(), 0, 1, tarantool.IterEq, A.X{r.RealtorId, r.PropertyId})
	if L.IsError(err, `Revenue.FindByRealtorIdPropertyId failed: `+r.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		r.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (r *Revenue) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "realtorId"
	, "propertyId"
	, "propertyBought"
	, "propertyCountry"
	, "buyerEmail"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (r *Revenue) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "realtorId"
	, "propertyId"
	, "propertyBought"
	, "propertyCountry"
	, "buyerEmail"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// ToUpdateArray generate slice of update command
func (r *Revenue) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, r.Id},
		A.X{`=`, 1, r.RealtorId},
		A.X{`=`, 2, r.PropertyId},
		A.X{`=`, 3, r.PropertyBought},
		A.X{`=`, 4, r.PropertyCountry},
		A.X{`=`, 5, r.BuyerEmail},
		A.X{`=`, 6, r.CreatedAt},
		A.X{`=`, 7, r.CreatedBy},
		A.X{`=`, 8, r.UpdatedAt},
		A.X{`=`, 9, r.UpdatedBy},
		A.X{`=`, 10, r.DeletedAt},
	}
}

// IdxId return name of the index
func (r *Revenue) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (r *Revenue) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxRealtorId return name of the index
func (r *Revenue) IdxRealtorId() int { //nolint:dupl false positive
	return 1
}

// SqlRealtorId return name of the column being indexed
func (r *Revenue) SqlRealtorId() string { //nolint:dupl false positive
	return `"realtorId"`
}

// IdxPropertyId return name of the index
func (r *Revenue) IdxPropertyId() int { //nolint:dupl false positive
	return 2
}

// SqlPropertyId return name of the column being indexed
func (r *Revenue) SqlPropertyId() string { //nolint:dupl false positive
	return `"propertyId"`
}

// IdxPropertyBought return name of the index
func (r *Revenue) IdxPropertyBought() int { //nolint:dupl false positive
	return 3
}

// SqlPropertyBought return name of the column being indexed
func (r *Revenue) SqlPropertyBought() string { //nolint:dupl false positive
	return `"propertyBought"`
}

// IdxPropertyCountry return name of the index
func (r *Revenue) IdxPropertyCountry() int { //nolint:dupl false positive
	return 4
}

// SqlPropertyCountry return name of the column being indexed
func (r *Revenue) SqlPropertyCountry() string { //nolint:dupl false positive
	return `"propertyCountry"`
}

// IdxBuyerEmail return name of the index
func (r *Revenue) IdxBuyerEmail() int { //nolint:dupl false positive
	return 5
}

// SqlBuyerEmail return name of the column being indexed
func (r *Revenue) SqlBuyerEmail() string { //nolint:dupl false positive
	return `"buyerEmail"`
}

// IdxCreatedAt return name of the index
func (r *Revenue) IdxCreatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedAt return name of the column being indexed
func (r *Revenue) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (r *Revenue) IdxCreatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlCreatedBy return name of the column being indexed
func (r *Revenue) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (r *Revenue) IdxUpdatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedAt return name of the column being indexed
func (r *Revenue) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (r *Revenue) IdxUpdatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlUpdatedBy return name of the column being indexed
func (r *Revenue) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (r *Revenue) IdxDeletedAt() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedAt return name of the column being indexed
func (r *Revenue) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// ToArray receiver fields to slice
func (r *Revenue) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		r.Id,              // 0
		r.RealtorId,       // 1
		r.PropertyId,      // 2
		r.PropertyBought,  // 3
		r.PropertyCountry, // 4
		r.BuyerEmail,      // 5
		r.CreatedAt,       // 6
		r.CreatedBy,       // 7
		r.UpdatedAt,       // 8
		r.UpdatedBy,       // 9
		r.DeletedAt,       // 10
	}
}

// FromArray convert slice to receiver fields
func (r *Revenue) FromArray(a A.X) *Revenue { //nolint:dupl false positive
	r.Id = X.ToU(a[0])
	r.RealtorId = X.ToU(a[1])
	r.PropertyId = X.ToU(a[2])
	r.PropertyBought = X.ToI(a[3])
	r.PropertyCountry = X.ToS(a[4])
	r.BuyerEmail = X.ToS(a[5])
	r.CreatedAt = X.ToI(a[6])
	r.CreatedBy = X.ToU(a[7])
	r.UpdatedAt = X.ToI(a[8])
	r.UpdatedBy = X.ToU(a[9])
	r.DeletedAt = X.ToI(a[10])
	return r
}

// FromUncensoredArray convert slice to receiver fields
func (r *Revenue) FromUncensoredArray(a A.X) *Revenue { //nolint:dupl false positive
	r.Id = X.ToU(a[0])
	r.RealtorId = X.ToU(a[1])
	r.PropertyId = X.ToU(a[2])
	r.PropertyBought = X.ToI(a[3])
	r.PropertyCountry = X.ToS(a[4])
	r.BuyerEmail = X.ToS(a[5])
	r.CreatedAt = X.ToI(a[6])
	r.CreatedBy = X.ToU(a[7])
	r.UpdatedAt = X.ToI(a[8])
	r.UpdatedBy = X.ToU(a[9])
	r.DeletedAt = X.ToI(a[10])
	return r
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (r *Revenue) FindOffsetLimit(offset, limit uint32, idx string) []Revenue { //nolint:dupl false positive
	var rows []Revenue
	res, err := r.Adapter.Select(r.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Revenue.FindOffsetLimit failed: `+r.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Revenue{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (r *Revenue) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := r.Adapter.Select(r.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Revenue.FindOffsetLimit failed: `+r.SpaceName()) {
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
func (r *Revenue) Total() int64 { //nolint:dupl false positive
	rows := r.Adapter.CallBoxSpace(r.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// RevenueFieldTypeMap returns key value of field name and key
var RevenueFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:              Tt.Unsigned,
	`realtorId`:       Tt.Unsigned,
	`propertyId`:      Tt.Unsigned,
	`propertyBought`:  Tt.Integer,
	`propertyCountry`: Tt.String,
	`buyerEmail`:      Tt.String,
	`createdAt`:       Tt.Integer,
	`createdBy`:       Tt.Unsigned,
	`updatedAt`:       Tt.Integer,
	`updatedBy`:       Tt.Unsigned,
	`deletedAt`:       Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Sales DAO reader/query struct
type Sales struct {
	Adapter         *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id              uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	PropertyId      uint64      `json:"propertyId,string" form:"propertyId" query:"propertyId" long:"propertyId" msg:"propertyId"`
	RealtorId       uint64      `json:"realtorId,string" form:"realtorId" query:"realtorId" long:"realtorId" msg:"realtorId"`
	PropertyCountry string      `json:"propertyCountry" form:"propertyCountry" query:"propertyCountry" long:"propertyCountry" msg:"propertyCountry"`
	BuyerId         uint64      `json:"buyerId,string" form:"buyerId" query:"buyerId" long:"buyerId" msg:"buyerId"`
	Price           string      `json:"price" form:"price" query:"price" long:"price" msg:"price"`
	BuyerEmail      string      `json:"buyerEmail" form:"buyerEmail" query:"buyerEmail" long:"buyerEmail" msg:"buyerEmail"`
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
		A.X{`=`, 7, s.SalesDate},
		A.X{`=`, 8, s.CreatedAt},
		A.X{`=`, 9, s.CreatedBy},
		A.X{`=`, 10, s.UpdatedAt},
		A.X{`=`, 11, s.UpdatedBy},
		A.X{`=`, 12, s.DeletedAt},
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

// IdxSalesDate return name of the index
func (s *Sales) IdxSalesDate() int { //nolint:dupl false positive
	return 7
}

// SqlSalesDate return name of the column being indexed
func (s *Sales) SqlSalesDate() string { //nolint:dupl false positive
	return `"salesDate"`
}

// IdxCreatedAt return name of the index
func (s *Sales) IdxCreatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlCreatedAt return name of the column being indexed
func (s *Sales) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (s *Sales) IdxCreatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlCreatedBy return name of the column being indexed
func (s *Sales) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (s *Sales) IdxUpdatedAt() int { //nolint:dupl false positive
	return 10
}

// SqlUpdatedAt return name of the column being indexed
func (s *Sales) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (s *Sales) IdxUpdatedBy() int { //nolint:dupl false positive
	return 11
}

// SqlUpdatedBy return name of the column being indexed
func (s *Sales) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (s *Sales) IdxDeletedAt() int { //nolint:dupl false positive
	return 12
}

// SqlDeletedAt return name of the column being indexed
func (s *Sales) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// ToArray receiver fields to slice
func (s *Sales) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.Id,              // 0
		s.PropertyId,      // 1
		s.RealtorId,       // 2
		s.PropertyCountry, // 3
		s.BuyerId,         // 4
		s.Price,           // 5
		s.BuyerEmail,      // 6
		s.SalesDate,       // 7
		s.CreatedAt,       // 8
		s.CreatedBy,       // 9
		s.UpdatedAt,       // 10
		s.UpdatedBy,       // 11
		s.DeletedAt,       // 12
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
	s.SalesDate = X.ToS(a[7])
	s.CreatedAt = X.ToI(a[8])
	s.CreatedBy = X.ToU(a[9])
	s.UpdatedAt = X.ToI(a[10])
	s.UpdatedBy = X.ToU(a[11])
	s.DeletedAt = X.ToI(a[12])
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
	s.SalesDate = X.ToS(a[7])
	s.CreatedAt = X.ToI(a[8])
	s.CreatedBy = X.ToU(a[9])
	s.UpdatedAt = X.ToI(a[10])
	s.UpdatedBy = X.ToU(a[11])
	s.DeletedAt = X.ToI(a[12])
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
	`salesDate`:       Tt.String,
	`createdAt`:       Tt.Integer,
	`createdBy`:       Tt.Unsigned,
	`updatedAt`:       Tt.Integer,
	`updatedBy`:       Tt.Unsigned,
	`deletedAt`:       Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
