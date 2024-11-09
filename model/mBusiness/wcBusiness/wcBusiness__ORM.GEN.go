package wcBusiness

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mBusiness/rqBusiness"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
)

// RevenueMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBusiness__ORM.GEN.go
type RevenueMutator struct {
	rqBusiness.Revenue
	mutations []A.X
	logs      []A.X
}

// NewRevenueMutator create new ORM writer/command object
func NewRevenueMutator(adapter *Tt.Adapter) (res *RevenueMutator) {
	res = &RevenueMutator{Revenue: rqBusiness.Revenue{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (r *RevenueMutator) Logs() []A.X { //nolint:dupl false positive
	return r.logs
}

// HaveMutation check whether Set* methods ever called
func (r *RevenueMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(r.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (r *RevenueMutator) ClearMutations() { //nolint:dupl false positive
	r.mutations = []A.X{}
	r.logs = []A.X{}
}

// func (r *RevenueMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := r.ToArray()
//	_, err := r.Adapter.Upsert(r.SpaceName(), arr, A.X{
//		A.X{`=`, 0, r.Id},
//		A.X{`=`, 1, r.RealtorId},
//		A.X{`=`, 2, r.PropertyId},
//		A.X{`=`, 3, r.PropertyBought},
//		A.X{`=`, 4, r.PropertyCountry},
//		A.X{`=`, 5, r.BuyerEmail},
//		A.X{`=`, 6, r.CreatedAt},
//		A.X{`=`, 7, r.CreatedBy},
//		A.X{`=`, 8, r.UpdatedAt},
//		A.X{`=`, 9, r.UpdatedBy},
//		A.X{`=`, 10, r.DeletedAt},
//	})
//	return !L.IsError(err, `Revenue.DoUpsert failed: `+r.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByRealtorIdPropertyId update all columns, error if not exists, not using mutations/Set*
func (r *RevenueMutator) DoOverwriteByRealtorIdPropertyId() bool { //nolint:dupl false positive
	_, err := r.Adapter.Update(r.SpaceName(), r.UniqueIndexRealtorIdPropertyId(), A.X{r.RealtorId, r.PropertyId}, r.ToUpdateArray())
	return !L.IsError(err, `Revenue.DoOverwriteByRealtorIdPropertyId failed: `+r.SpaceName())
}

// DoUpdateByRealtorIdPropertyId update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (r *RevenueMutator) DoUpdateByRealtorIdPropertyId() bool { //nolint:dupl false positive
	if !r.HaveMutation() {
		return true
	}
	_, err := r.Adapter.Update(r.SpaceName(), r.UniqueIndexRealtorIdPropertyId(), A.X{r.RealtorId, r.PropertyId}, r.mutations)
	return !L.IsError(err, `Revenue.DoUpdateByRealtorIdPropertyId failed: `+r.SpaceName())
}

// DoDeletePermanentByRealtorIdPropertyId permanent delete
func (r *RevenueMutator) DoDeletePermanentByRealtorIdPropertyId() bool { //nolint:dupl false positive
	_, err := r.Adapter.Delete(r.SpaceName(), r.UniqueIndexRealtorIdPropertyId(), A.X{r.RealtorId, r.PropertyId})
	return !L.IsError(err, `Revenue.DoDeletePermanentByRealtorIdPropertyId failed: `+r.SpaceName())
}

// DoInsert insert, error if already exists
func (r *RevenueMutator) DoInsert() bool { //nolint:dupl false positive
	arr := r.ToArray()
	_, err := r.Adapter.Insert(r.SpaceName(), arr)
	return !L.IsError(err, `Revenue.DoInsert failed: `+r.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (r *RevenueMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := r.ToArray()
	_, err := r.Adapter.Replace(r.SpaceName(), arr)
	return !L.IsError(err, `Revenue.DoUpsert failed: `+r.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (r *RevenueMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != r.Id {
		r.mutations = append(r.mutations, A.X{`=`, 0, val})
		r.logs = append(r.logs, A.X{`id`, r.Id, val})
		r.Id = val
		return true
	}
	return false
}

// SetRealtorId create mutations, should not duplicate
func (r *RevenueMutator) SetRealtorId(val uint64) bool { //nolint:dupl false positive
	if val != r.RealtorId {
		r.mutations = append(r.mutations, A.X{`=`, 1, val})
		r.logs = append(r.logs, A.X{`realtorId`, r.RealtorId, val})
		r.RealtorId = val
		return true
	}
	return false
}

// SetPropertyId create mutations, should not duplicate
func (r *RevenueMutator) SetPropertyId(val uint64) bool { //nolint:dupl false positive
	if val != r.PropertyId {
		r.mutations = append(r.mutations, A.X{`=`, 2, val})
		r.logs = append(r.logs, A.X{`propertyId`, r.PropertyId, val})
		r.PropertyId = val
		return true
	}
	return false
}

// SetPropertyBought create mutations, should not duplicate
func (r *RevenueMutator) SetPropertyBought(val int64) bool { //nolint:dupl false positive
	if val != r.PropertyBought {
		r.mutations = append(r.mutations, A.X{`=`, 3, val})
		r.logs = append(r.logs, A.X{`propertyBought`, r.PropertyBought, val})
		r.PropertyBought = val
		return true
	}
	return false
}

// SetPropertyCountry create mutations, should not duplicate
func (r *RevenueMutator) SetPropertyCountry(val string) bool { //nolint:dupl false positive
	if val != r.PropertyCountry {
		r.mutations = append(r.mutations, A.X{`=`, 4, val})
		r.logs = append(r.logs, A.X{`propertyCountry`, r.PropertyCountry, val})
		r.PropertyCountry = val
		return true
	}
	return false
}

// SetBuyerEmail create mutations, should not duplicate
func (r *RevenueMutator) SetBuyerEmail(val string) bool { //nolint:dupl false positive
	if val != r.BuyerEmail {
		r.mutations = append(r.mutations, A.X{`=`, 5, val})
		r.logs = append(r.logs, A.X{`buyerEmail`, r.BuyerEmail, val})
		r.BuyerEmail = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (r *RevenueMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != r.CreatedAt {
		r.mutations = append(r.mutations, A.X{`=`, 6, val})
		r.logs = append(r.logs, A.X{`createdAt`, r.CreatedAt, val})
		r.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (r *RevenueMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != r.CreatedBy {
		r.mutations = append(r.mutations, A.X{`=`, 7, val})
		r.logs = append(r.logs, A.X{`createdBy`, r.CreatedBy, val})
		r.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (r *RevenueMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != r.UpdatedAt {
		r.mutations = append(r.mutations, A.X{`=`, 8, val})
		r.logs = append(r.logs, A.X{`updatedAt`, r.UpdatedAt, val})
		r.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (r *RevenueMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != r.UpdatedBy {
		r.mutations = append(r.mutations, A.X{`=`, 9, val})
		r.logs = append(r.logs, A.X{`updatedBy`, r.UpdatedBy, val})
		r.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (r *RevenueMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != r.DeletedAt {
		r.mutations = append(r.mutations, A.X{`=`, 10, val})
		r.logs = append(r.logs, A.X{`deletedAt`, r.DeletedAt, val})
		r.DeletedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (r *RevenueMutator) SetAll(from rqBusiness.Revenue, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		r.Id = from.Id
		changed = true
	}
	if !excludeMap[`realtorId`] && (forceMap[`realtorId`] || from.RealtorId != 0) {
		r.RealtorId = from.RealtorId
		changed = true
	}
	if !excludeMap[`propertyId`] && (forceMap[`propertyId`] || from.PropertyId != 0) {
		r.PropertyId = from.PropertyId
		changed = true
	}
	if !excludeMap[`propertyBought`] && (forceMap[`propertyBought`] || from.PropertyBought != 0) {
		r.PropertyBought = from.PropertyBought
		changed = true
	}
	if !excludeMap[`propertyCountry`] && (forceMap[`propertyCountry`] || from.PropertyCountry != ``) {
		r.PropertyCountry = S.Trim(from.PropertyCountry)
		changed = true
	}
	if !excludeMap[`buyerEmail`] && (forceMap[`buyerEmail`] || from.BuyerEmail != ``) {
		r.BuyerEmail = S.Trim(from.BuyerEmail)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		r.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		r.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		r.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		r.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		r.DeletedAt = from.DeletedAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// SalesMutator DAO writer/command struct
type SalesMutator struct {
	rqBusiness.Sales
	mutations []A.X
	logs      []A.X
}

// NewSalesMutator create new ORM writer/command object
func NewSalesMutator(adapter *Tt.Adapter) (res *SalesMutator) {
	res = &SalesMutator{Sales: rqBusiness.Sales{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (s *SalesMutator) Logs() []A.X { //nolint:dupl false positive
	return s.logs
}

// HaveMutation check whether Set* methods ever called
func (s *SalesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(s.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (s *SalesMutator) ClearMutations() { //nolint:dupl false positive
	s.mutations = []A.X{}
	s.logs = []A.X{}
}

// func (s *SalesMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := s.ToArray()
//	_, err := s.Adapter.Upsert(s.SpaceName(), arr, A.X{
//		A.X{`=`, 0, s.Id},
//		A.X{`=`, 1, s.PropertyId},
//		A.X{`=`, 2, s.RealtorId},
//		A.X{`=`, 3, s.PropertyCountry},
//		A.X{`=`, 4, s.BuyerId},
//		A.X{`=`, 5, s.Price},
//		A.X{`=`, 6, s.BuyerEmail},
//		A.X{`=`, 7, s.SalesDate},
//		A.X{`=`, 8, s.CreatedAt},
//		A.X{`=`, 9, s.CreatedBy},
//		A.X{`=`, 10, s.UpdatedAt},
//		A.X{`=`, 11, s.UpdatedBy},
//		A.X{`=`, 12, s.DeletedAt},
//	})
//	return !L.IsError(err, `Sales.DoUpsert failed: `+s.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByPropertyIdSalesDate update all columns, error if not exists, not using mutations/Set*
func (s *SalesMutator) DoOverwriteByPropertyIdSalesDate() bool { //nolint:dupl false positive
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexPropertyIdSalesDate(), A.X{s.PropertyId, s.SalesDate}, s.ToUpdateArray())
	return !L.IsError(err, `Sales.DoOverwriteByPropertyIdSalesDate failed: `+s.SpaceName())
}

// DoUpdateByPropertyIdSalesDate update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (s *SalesMutator) DoUpdateByPropertyIdSalesDate() bool { //nolint:dupl false positive
	if !s.HaveMutation() {
		return true
	}
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexPropertyIdSalesDate(), A.X{s.PropertyId, s.SalesDate}, s.mutations)
	return !L.IsError(err, `Sales.DoUpdateByPropertyIdSalesDate failed: `+s.SpaceName())
}

// DoDeletePermanentByPropertyIdSalesDate permanent delete
func (s *SalesMutator) DoDeletePermanentByPropertyIdSalesDate() bool { //nolint:dupl false positive
	_, err := s.Adapter.Delete(s.SpaceName(), s.UniqueIndexPropertyIdSalesDate(), A.X{s.PropertyId, s.SalesDate})
	return !L.IsError(err, `Sales.DoDeletePermanentByPropertyIdSalesDate failed: `+s.SpaceName())
}

// DoInsert insert, error if already exists
func (s *SalesMutator) DoInsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Insert(s.SpaceName(), arr)
	return !L.IsError(err, `Sales.DoInsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (s *SalesMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Replace(s.SpaceName(), arr)
	return !L.IsError(err, `Sales.DoUpsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (s *SalesMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != s.Id {
		s.mutations = append(s.mutations, A.X{`=`, 0, val})
		s.logs = append(s.logs, A.X{`id`, s.Id, val})
		s.Id = val
		return true
	}
	return false
}

// SetPropertyId create mutations, should not duplicate
func (s *SalesMutator) SetPropertyId(val uint64) bool { //nolint:dupl false positive
	if val != s.PropertyId {
		s.mutations = append(s.mutations, A.X{`=`, 1, val})
		s.logs = append(s.logs, A.X{`propertyId`, s.PropertyId, val})
		s.PropertyId = val
		return true
	}
	return false
}

// SetRealtorId create mutations, should not duplicate
func (s *SalesMutator) SetRealtorId(val uint64) bool { //nolint:dupl false positive
	if val != s.RealtorId {
		s.mutations = append(s.mutations, A.X{`=`, 2, val})
		s.logs = append(s.logs, A.X{`realtorId`, s.RealtorId, val})
		s.RealtorId = val
		return true
	}
	return false
}

// SetPropertyCountry create mutations, should not duplicate
func (s *SalesMutator) SetPropertyCountry(val string) bool { //nolint:dupl false positive
	if val != s.PropertyCountry {
		s.mutations = append(s.mutations, A.X{`=`, 3, val})
		s.logs = append(s.logs, A.X{`propertyCountry`, s.PropertyCountry, val})
		s.PropertyCountry = val
		return true
	}
	return false
}

// SetBuyerId create mutations, should not duplicate
func (s *SalesMutator) SetBuyerId(val uint64) bool { //nolint:dupl false positive
	if val != s.BuyerId {
		s.mutations = append(s.mutations, A.X{`=`, 4, val})
		s.logs = append(s.logs, A.X{`buyerId`, s.BuyerId, val})
		s.BuyerId = val
		return true
	}
	return false
}

// SetPrice create mutations, should not duplicate
func (s *SalesMutator) SetPrice(val string) bool { //nolint:dupl false positive
	if val != s.Price {
		s.mutations = append(s.mutations, A.X{`=`, 5, val})
		s.logs = append(s.logs, A.X{`price`, s.Price, val})
		s.Price = val
		return true
	}
	return false
}

// SetBuyerEmail create mutations, should not duplicate
func (s *SalesMutator) SetBuyerEmail(val string) bool { //nolint:dupl false positive
	if val != s.BuyerEmail {
		s.mutations = append(s.mutations, A.X{`=`, 6, val})
		s.logs = append(s.logs, A.X{`buyerEmail`, s.BuyerEmail, val})
		s.BuyerEmail = val
		return true
	}
	return false
}

// SetSalesDate create mutations, should not duplicate
func (s *SalesMutator) SetSalesDate(val string) bool { //nolint:dupl false positive
	if val != s.SalesDate {
		s.mutations = append(s.mutations, A.X{`=`, 7, val})
		s.logs = append(s.logs, A.X{`salesDate`, s.SalesDate, val})
		s.SalesDate = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (s *SalesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.CreatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 8, val})
		s.logs = append(s.logs, A.X{`createdAt`, s.CreatedAt, val})
		s.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (s *SalesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.CreatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 9, val})
		s.logs = append(s.logs, A.X{`createdBy`, s.CreatedBy, val})
		s.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.UpdatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 10, val})
		s.logs = append(s.logs, A.X{`updatedAt`, s.UpdatedAt, val})
		s.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.UpdatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 11, val})
		s.logs = append(s.logs, A.X{`updatedBy`, s.UpdatedBy, val})
		s.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (s *SalesMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != s.DeletedAt {
		s.mutations = append(s.mutations, A.X{`=`, 12, val})
		s.logs = append(s.logs, A.X{`deletedAt`, s.DeletedAt, val})
		s.DeletedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (s *SalesMutator) SetAll(from rqBusiness.Sales, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		s.Id = from.Id
		changed = true
	}
	if !excludeMap[`propertyId`] && (forceMap[`propertyId`] || from.PropertyId != 0) {
		s.PropertyId = from.PropertyId
		changed = true
	}
	if !excludeMap[`realtorId`] && (forceMap[`realtorId`] || from.RealtorId != 0) {
		s.RealtorId = from.RealtorId
		changed = true
	}
	if !excludeMap[`propertyCountry`] && (forceMap[`propertyCountry`] || from.PropertyCountry != ``) {
		s.PropertyCountry = S.Trim(from.PropertyCountry)
		changed = true
	}
	if !excludeMap[`buyerId`] && (forceMap[`buyerId`] || from.BuyerId != 0) {
		s.BuyerId = from.BuyerId
		changed = true
	}
	if !excludeMap[`price`] && (forceMap[`price`] || from.Price != ``) {
		s.Price = S.Trim(from.Price)
		changed = true
	}
	if !excludeMap[`buyerEmail`] && (forceMap[`buyerEmail`] || from.BuyerEmail != ``) {
		s.BuyerEmail = S.Trim(from.BuyerEmail)
		changed = true
	}
	if !excludeMap[`salesDate`] && (forceMap[`salesDate`] || from.SalesDate != ``) {
		s.SalesDate = S.Trim(from.SalesDate)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		s.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		s.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		s.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		s.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		s.DeletedAt = from.DeletedAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
