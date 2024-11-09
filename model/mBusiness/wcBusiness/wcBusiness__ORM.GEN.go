package wcBusiness

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mBusiness/rqBusiness"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// SalesMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBusiness__ORM.GEN.go
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

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (s *SalesMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id}, s.ToUpdateArray())
	return !L.IsError(err, `Sales.DoOverwriteById failed: `+s.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (s *SalesMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !s.HaveMutation() {
		return true
	}
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id}, s.mutations)
	return !L.IsError(err, `Sales.DoUpdateById failed: `+s.SpaceName())
}

// DoDeletePermanentById permanent delete
func (s *SalesMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := s.Adapter.Delete(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id})
	return !L.IsError(err, `Sales.DoDeletePermanentById failed: `+s.SpaceName())
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
//		A.X{`=`, 7, s.EmailNotFound},
//		A.X{`=`, 8, s.SalesDate},
//		A.X{`=`, 9, s.CreatedAt},
//		A.X{`=`, 10, s.CreatedBy},
//		A.X{`=`, 11, s.UpdatedAt},
//		A.X{`=`, 12, s.UpdatedBy},
//		A.X{`=`, 13, s.DeletedAt},
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
	row, err := s.Adapter.Insert(s.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			s.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Sales.DoInsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (s *SalesMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	row, err := s.Adapter.Replace(s.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			s.Id = X.ToU(tup[0][0])
		}
	}
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

// SetEmailNotFound create mutations, should not duplicate
func (s *SalesMutator) SetEmailNotFound(val string) bool { //nolint:dupl false positive
	if val != s.EmailNotFound {
		s.mutations = append(s.mutations, A.X{`=`, 7, val})
		s.logs = append(s.logs, A.X{`emailNotFound`, s.EmailNotFound, val})
		s.EmailNotFound = val
		return true
	}
	return false
}

// SetSalesDate create mutations, should not duplicate
func (s *SalesMutator) SetSalesDate(val string) bool { //nolint:dupl false positive
	if val != s.SalesDate {
		s.mutations = append(s.mutations, A.X{`=`, 8, val})
		s.logs = append(s.logs, A.X{`salesDate`, s.SalesDate, val})
		s.SalesDate = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (s *SalesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.CreatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 9, val})
		s.logs = append(s.logs, A.X{`createdAt`, s.CreatedAt, val})
		s.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (s *SalesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.CreatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 10, val})
		s.logs = append(s.logs, A.X{`createdBy`, s.CreatedBy, val})
		s.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.UpdatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 11, val})
		s.logs = append(s.logs, A.X{`updatedAt`, s.UpdatedAt, val})
		s.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.UpdatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 12, val})
		s.logs = append(s.logs, A.X{`updatedBy`, s.UpdatedBy, val})
		s.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (s *SalesMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != s.DeletedAt {
		s.mutations = append(s.mutations, A.X{`=`, 13, val})
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
	if !excludeMap[`emailNotFound`] && (forceMap[`emailNotFound`] || from.EmailNotFound != ``) {
		s.EmailNotFound = S.Trim(from.EmailNotFound)
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
