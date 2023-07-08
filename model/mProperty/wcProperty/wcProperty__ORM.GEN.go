package wcProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// PropertyMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcProperty__ORM.GEN.go
type PropertyMutator struct {
	rqProperty.Property
	mutations []A.X
	logs      []A.X
}

// NewPropertyMutator create new ORM writer/command object
func NewPropertyMutator(adapter *Tt.Adapter) *PropertyMutator {
	return &PropertyMutator{Property: rqProperty.Property{Adapter: adapter}}
}

// Logs get array of logs [field, old, new]
func (p *PropertyMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PropertyMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PropertyMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *PropertyMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.ToUpdateArray())
	return !L.IsError(err, `Property.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PropertyMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.mutations)
	return !L.IsError(err, `Property.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *PropertyMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id})
	return !L.IsError(err, `Property.DoDeletePermanentById failed: `+p.SpaceName())
}

// func (p *PropertyMutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := p.Adapter.Upsert(p.SpaceName(), p.ToArray(), A.X{
//		A.X{`=`, 0, p.Id},
//		A.X{`=`, 1, p.UniqPropKey},
//		A.X{`=`, 2, p.SerialNumber},
//		A.X{`=`, 3, p.SizeM2},
//		A.X{`=`, 4, p.MainUse},
//		A.X{`=`, 5, p.MainBuildingMaterial},
//		A.X{`=`, 6, p.ConstructCompletedDate},
//		A.X{`=`, 7, p.NumberOfFloors},
//		A.X{`=`, 8, p.BuildingLamination},
//		A.X{`=`, 9, p.Address},
//		A.X{`=`, 10, p.District},
//		A.X{`=`, 11, p.Note},
//		A.X{`=`, 12, p.Coord},
//		A.X{`=`, 13, p.CreatedAt},
//		A.X{`=`, 14, p.CreatedBy},
//		A.X{`=`, 15, p.UpdatedAt},
//		A.X{`=`, 16, p.UpdatedBy},
//		A.X{`=`, 17, p.DeletedAt},
//		A.X{`=`, 18, p.FormattedAddress},
//	})
//	return !L.IsError(err, `Property.DoUpsert failed: `+p.SpaceName())
// }

// DoOverwriteByUniqPropKey update all columns, error if not exists, not using mutations/Set*
func (p *PropertyMutator) DoOverwriteByUniqPropKey() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexUniqPropKey(), A.X{p.UniqPropKey}, p.ToUpdateArray())
	return !L.IsError(err, `Property.DoOverwriteByUniqPropKey failed: `+p.SpaceName())
}

// DoUpdateByUniqPropKey update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PropertyMutator) DoUpdateByUniqPropKey() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexUniqPropKey(), A.X{p.UniqPropKey}, p.mutations)
	return !L.IsError(err, `Property.DoUpdateByUniqPropKey failed: `+p.SpaceName())
}

// DoDeletePermanentByUniqPropKey permanent delete
func (p *PropertyMutator) DoDeletePermanentByUniqPropKey() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexUniqPropKey(), A.X{p.UniqPropKey})
	return !L.IsError(err, `Property.DoDeletePermanentByUniqPropKey failed: `+p.SpaceName())
}

// DoInsert insert, error if already exists
func (p *PropertyMutator) DoInsert() bool { //nolint:dupl false positive
	row, err := p.Adapter.Insert(p.SpaceName(), p.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Property.DoInsert failed: `+p.SpaceName())
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (p *PropertyMutator) DoUpsert() bool { //nolint:dupl false positive
	_, err := p.Adapter.Replace(p.SpaceName(), p.ToArray())
	return !L.IsError(err, `Property.DoUpsert failed: `+p.SpaceName())
}

// SetId create mutations, should not duplicate
func (p *PropertyMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations = append(p.mutations, A.X{`=`, 0, val})
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetUniqPropKey create mutations, should not duplicate
func (p *PropertyMutator) SetUniqPropKey(val string) bool { //nolint:dupl false positive
	if val != p.UniqPropKey {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.logs = append(p.logs, A.X{`uniqPropKey`, p.UniqPropKey, val})
		p.UniqPropKey = val
		return true
	}
	return false
}

// SetSerialNumber create mutations, should not duplicate
func (p *PropertyMutator) SetSerialNumber(val string) bool { //nolint:dupl false positive
	if val != p.SerialNumber {
		p.mutations = append(p.mutations, A.X{`=`, 2, val})
		p.logs = append(p.logs, A.X{`serialNumber`, p.SerialNumber, val})
		p.SerialNumber = val
		return true
	}
	return false
}

// SetSizeM2 create mutations, should not duplicate
func (p *PropertyMutator) SetSizeM2(val string) bool { //nolint:dupl false positive
	if val != p.SizeM2 {
		p.mutations = append(p.mutations, A.X{`=`, 3, val})
		p.logs = append(p.logs, A.X{`sizeM2`, p.SizeM2, val})
		p.SizeM2 = val
		return true
	}
	return false
}

// SetMainUse create mutations, should not duplicate
func (p *PropertyMutator) SetMainUse(val string) bool { //nolint:dupl false positive
	if val != p.MainUse {
		p.mutations = append(p.mutations, A.X{`=`, 4, val})
		p.logs = append(p.logs, A.X{`mainUse`, p.MainUse, val})
		p.MainUse = val
		return true
	}
	return false
}

// SetMainBuildingMaterial create mutations, should not duplicate
func (p *PropertyMutator) SetMainBuildingMaterial(val string) bool { //nolint:dupl false positive
	if val != p.MainBuildingMaterial {
		p.mutations = append(p.mutations, A.X{`=`, 5, val})
		p.logs = append(p.logs, A.X{`mainBuildingMaterial`, p.MainBuildingMaterial, val})
		p.MainBuildingMaterial = val
		return true
	}
	return false
}

// SetConstructCompletedDate create mutations, should not duplicate
func (p *PropertyMutator) SetConstructCompletedDate(val string) bool { //nolint:dupl false positive
	if val != p.ConstructCompletedDate {
		p.mutations = append(p.mutations, A.X{`=`, 6, val})
		p.logs = append(p.logs, A.X{`constructCompletedDate`, p.ConstructCompletedDate, val})
		p.ConstructCompletedDate = val
		return true
	}
	return false
}

// SetNumberOfFloors create mutations, should not duplicate
func (p *PropertyMutator) SetNumberOfFloors(val string) bool { //nolint:dupl false positive
	if val != p.NumberOfFloors {
		p.mutations = append(p.mutations, A.X{`=`, 7, val})
		p.logs = append(p.logs, A.X{`numberOfFloors`, p.NumberOfFloors, val})
		p.NumberOfFloors = val
		return true
	}
	return false
}

// SetBuildingLamination create mutations, should not duplicate
func (p *PropertyMutator) SetBuildingLamination(val string) bool { //nolint:dupl false positive
	if val != p.BuildingLamination {
		p.mutations = append(p.mutations, A.X{`=`, 8, val})
		p.logs = append(p.logs, A.X{`buildingLamination`, p.BuildingLamination, val})
		p.BuildingLamination = val
		return true
	}
	return false
}

// SetAddress create mutations, should not duplicate
func (p *PropertyMutator) SetAddress(val string) bool { //nolint:dupl false positive
	if val != p.Address {
		p.mutations = append(p.mutations, A.X{`=`, 9, val})
		p.logs = append(p.logs, A.X{`address`, p.Address, val})
		p.Address = val
		return true
	}
	return false
}

// SetDistrict create mutations, should not duplicate
func (p *PropertyMutator) SetDistrict(val string) bool { //nolint:dupl false positive
	if val != p.District {
		p.mutations = append(p.mutations, A.X{`=`, 10, val})
		p.logs = append(p.logs, A.X{`district`, p.District, val})
		p.District = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (p *PropertyMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != p.Note {
		p.mutations = append(p.mutations, A.X{`=`, 11, val})
		p.logs = append(p.logs, A.X{`note`, p.Note, val})
		p.Note = val
		return true
	}
	return false
}

// SetCoord create mutations, should not duplicate
func (p *PropertyMutator) SetCoord(val []any) bool { //nolint:dupl false positive
	p.mutations = append(p.mutations, A.X{`=`, 12, val})
	p.logs = append(p.logs, A.X{`coord`, p.Coord, val})
	p.Coord = val
	return true
}

// SetCreatedAt create mutations, should not duplicate
func (p *PropertyMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 13, val})
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PropertyMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 14, val})
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PropertyMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 15, val})
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PropertyMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 16, val})
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PropertyMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations = append(p.mutations, A.X{`=`, 17, val})
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetFormattedAddress create mutations, should not duplicate
func (p *PropertyMutator) SetFormattedAddress(val string) bool { //nolint:dupl false positive
	if val != p.FormattedAddress {
		p.mutations = append(p.mutations, A.X{`=`, 18, val})
		p.logs = append(p.logs, A.X{`formattedAddress`, p.FormattedAddress, val})
		p.FormattedAddress = val
		return true
	}
	return false
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyHistoryMutator DAO writer/command struct
type PropertyHistoryMutator struct {
	rqProperty.PropertyHistory
	mutations []A.X
	logs      []A.X
}

// NewPropertyHistoryMutator create new ORM writer/command object
func NewPropertyHistoryMutator(adapter *Tt.Adapter) *PropertyHistoryMutator {
	return &PropertyHistoryMutator{PropertyHistory: rqProperty.PropertyHistory{Adapter: adapter}}
}

// Logs get array of logs [field, old, new]
func (p *PropertyHistoryMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PropertyHistoryMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PropertyHistoryMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *PropertyHistoryMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.ToUpdateArray())
	return !L.IsError(err, `PropertyHistory.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PropertyHistoryMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.mutations)
	return !L.IsError(err, `PropertyHistory.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *PropertyHistoryMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id})
	return !L.IsError(err, `PropertyHistory.DoDeletePermanentById failed: `+p.SpaceName())
}

// func (p *PropertyHistoryMutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := p.Adapter.Upsert(p.SpaceName(), p.ToArray(), A.X{
//		A.X{`=`, 0, p.Id},
//		A.X{`=`, 1, p.PropertyKey},
//		A.X{`=`, 2, p.TransactionKey},
//		A.X{`=`, 3, p.TransactionType},
//		A.X{`=`, 4, p.TransactionSign},
//		A.X{`=`, 5, p.TransactionTime},
//		A.X{`=`, 6, p.TransactionDateNormal},
//		A.X{`=`, 7, p.TransactionNumber},
//		A.X{`=`, 8, p.PriceNtd},
//		A.X{`=`, 9, p.PricePerUnit},
//		A.X{`=`, 10, p.Price},
//		A.X{`=`, 11, p.Address},
//		A.X{`=`, 12, p.District},
//		A.X{`=`, 13, p.Note},
//		A.X{`=`, 14, p.CreatedAt},
//		A.X{`=`, 15, p.CreatedBy},
//		A.X{`=`, 16, p.UpdatedAt},
//		A.X{`=`, 17, p.UpdatedBy},
//		A.X{`=`, 18, p.DeletedAt},
//	})
//	return !L.IsError(err, `PropertyHistory.DoUpsert failed: `+p.SpaceName())
// }

// DoOverwriteByTransactionKey update all columns, error if not exists, not using mutations/Set*
func (p *PropertyHistoryMutator) DoOverwriteByTransactionKey() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexTransactionKey(), A.X{p.TransactionKey}, p.ToUpdateArray())
	return !L.IsError(err, `PropertyHistory.DoOverwriteByTransactionKey failed: `+p.SpaceName())
}

// DoUpdateByTransactionKey update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PropertyHistoryMutator) DoUpdateByTransactionKey() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexTransactionKey(), A.X{p.TransactionKey}, p.mutations)
	return !L.IsError(err, `PropertyHistory.DoUpdateByTransactionKey failed: `+p.SpaceName())
}

// DoDeletePermanentByTransactionKey permanent delete
func (p *PropertyHistoryMutator) DoDeletePermanentByTransactionKey() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexTransactionKey(), A.X{p.TransactionKey})
	return !L.IsError(err, `PropertyHistory.DoDeletePermanentByTransactionKey failed: `+p.SpaceName())
}

// DoInsert insert, error if already exists
func (p *PropertyHistoryMutator) DoInsert() bool { //nolint:dupl false positive
	row, err := p.Adapter.Insert(p.SpaceName(), p.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `PropertyHistory.DoInsert failed: `+p.SpaceName())
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (p *PropertyHistoryMutator) DoUpsert() bool { //nolint:dupl false positive
	_, err := p.Adapter.Replace(p.SpaceName(), p.ToArray())
	return !L.IsError(err, `PropertyHistory.DoUpsert failed: `+p.SpaceName())
}

// SetId create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations = append(p.mutations, A.X{`=`, 0, val})
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetPropertyKey create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetPropertyKey(val string) bool { //nolint:dupl false positive
	if val != p.PropertyKey {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.logs = append(p.logs, A.X{`propertyKey`, p.PropertyKey, val})
		p.PropertyKey = val
		return true
	}
	return false
}

// SetTransactionKey create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetTransactionKey(val string) bool { //nolint:dupl false positive
	if val != p.TransactionKey {
		p.mutations = append(p.mutations, A.X{`=`, 2, val})
		p.logs = append(p.logs, A.X{`transactionKey`, p.TransactionKey, val})
		p.TransactionKey = val
		return true
	}
	return false
}

// SetTransactionType create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetTransactionType(val string) bool { //nolint:dupl false positive
	if val != p.TransactionType {
		p.mutations = append(p.mutations, A.X{`=`, 3, val})
		p.logs = append(p.logs, A.X{`transactionType`, p.TransactionType, val})
		p.TransactionType = val
		return true
	}
	return false
}

// SetTransactionSign create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetTransactionSign(val string) bool { //nolint:dupl false positive
	if val != p.TransactionSign {
		p.mutations = append(p.mutations, A.X{`=`, 4, val})
		p.logs = append(p.logs, A.X{`transactionSign`, p.TransactionSign, val})
		p.TransactionSign = val
		return true
	}
	return false
}

// SetTransactionTime create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetTransactionTime(val string) bool { //nolint:dupl false positive
	if val != p.TransactionTime {
		p.mutations = append(p.mutations, A.X{`=`, 5, val})
		p.logs = append(p.logs, A.X{`transactionTime`, p.TransactionTime, val})
		p.TransactionTime = val
		return true
	}
	return false
}

// SetTransactionDateNormal create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetTransactionDateNormal(val string) bool { //nolint:dupl false positive
	if val != p.TransactionDateNormal {
		p.mutations = append(p.mutations, A.X{`=`, 6, val})
		p.logs = append(p.logs, A.X{`transactionDateNormal`, p.TransactionDateNormal, val})
		p.TransactionDateNormal = val
		return true
	}
	return false
}

// SetTransactionNumber create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetTransactionNumber(val string) bool { //nolint:dupl false positive
	if val != p.TransactionNumber {
		p.mutations = append(p.mutations, A.X{`=`, 7, val})
		p.logs = append(p.logs, A.X{`transactionNumber`, p.TransactionNumber, val})
		p.TransactionNumber = val
		return true
	}
	return false
}

// SetPriceNtd create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetPriceNtd(val int64) bool { //nolint:dupl false positive
	if val != p.PriceNtd {
		p.mutations = append(p.mutations, A.X{`=`, 8, val})
		p.logs = append(p.logs, A.X{`priceNtd`, p.PriceNtd, val})
		p.PriceNtd = val
		return true
	}
	return false
}

// SetPricePerUnit create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetPricePerUnit(val int64) bool { //nolint:dupl false positive
	if val != p.PricePerUnit {
		p.mutations = append(p.mutations, A.X{`=`, 9, val})
		p.logs = append(p.logs, A.X{`pricePerUnit`, p.PricePerUnit, val})
		p.PricePerUnit = val
		return true
	}
	return false
}

// SetPrice create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetPrice(val int64) bool { //nolint:dupl false positive
	if val != p.Price {
		p.mutations = append(p.mutations, A.X{`=`, 10, val})
		p.logs = append(p.logs, A.X{`price`, p.Price, val})
		p.Price = val
		return true
	}
	return false
}

// SetAddress create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetAddress(val string) bool { //nolint:dupl false positive
	if val != p.Address {
		p.mutations = append(p.mutations, A.X{`=`, 11, val})
		p.logs = append(p.logs, A.X{`address`, p.Address, val})
		p.Address = val
		return true
	}
	return false
}

// SetDistrict create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetDistrict(val string) bool { //nolint:dupl false positive
	if val != p.District {
		p.mutations = append(p.mutations, A.X{`=`, 12, val})
		p.logs = append(p.logs, A.X{`district`, p.District, val})
		p.District = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != p.Note {
		p.mutations = append(p.mutations, A.X{`=`, 13, val})
		p.logs = append(p.logs, A.X{`note`, p.Note, val})
		p.Note = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 14, val})
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 15, val})
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 16, val})
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 17, val})
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations = append(p.mutations, A.X{`=`, 18, val})
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
