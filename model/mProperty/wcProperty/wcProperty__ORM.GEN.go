package wcProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/X"
)

// PropLikeCountMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcProperty__ORM.GEN.go
type PropLikeCountMutator struct {
	rqProperty.PropLikeCount
	mutations []A.X
	logs      []A.X
}

// NewPropLikeCountMutator create new ORM writer/command object
func NewPropLikeCountMutator(adapter *Tt.Adapter) (res *PropLikeCountMutator) {
	res = &PropLikeCountMutator{PropLikeCount: rqProperty.PropLikeCount{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (p *PropLikeCountMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PropLikeCountMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PropLikeCountMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
	p.logs = []A.X{}
}

// func (p *PropLikeCountMutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := p.Adapter.Upsert(p.SpaceName(), p.ToArray(), A.X{
//		A.X{`=`, 0, p.PropId},
//		A.X{`=`, 1, p.Count},
//	})
//	return !L.IsError(err, `PropLikeCount.DoUpsert failed: `+p.SpaceName())
// }

// DoOverwriteByPropId update all columns, error if not exists, not using mutations/Set*
func (p *PropLikeCountMutator) DoOverwriteByPropId() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexPropId(), A.X{p.PropId}, p.ToUpdateArray())
	return !L.IsError(err, `PropLikeCount.DoOverwriteByPropId failed: `+p.SpaceName())
}

// DoUpdateByPropId update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PropLikeCountMutator) DoUpdateByPropId() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexPropId(), A.X{p.PropId}, p.mutations)
	return !L.IsError(err, `PropLikeCount.DoUpdateByPropId failed: `+p.SpaceName())
}

// DoDeletePermanentByPropId permanent delete
func (p *PropLikeCountMutator) DoDeletePermanentByPropId() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexPropId(), A.X{p.PropId})
	return !L.IsError(err, `PropLikeCount.DoDeletePermanentByPropId failed: `+p.SpaceName())
}

// DoInsert insert, error if already exists
func (p *PropLikeCountMutator) DoInsert() bool { //nolint:dupl false positive
	_, err := p.Adapter.Insert(p.SpaceName(), p.ToArray())
	return !L.IsError(err, `PropLikeCount.DoInsert failed: `+p.SpaceName())
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (p *PropLikeCountMutator) DoUpsert() bool { //nolint:dupl false positive
	_, err := p.Adapter.Replace(p.SpaceName(), p.ToArray())
	return !L.IsError(err, `PropLikeCount.DoUpsert failed: `+p.SpaceName())
}

// SetPropId create mutations, should not duplicate
func (p *PropLikeCountMutator) SetPropId(val uint64) bool { //nolint:dupl false positive
	if val != p.PropId {
		p.mutations = append(p.mutations, A.X{`=`, 0, val})
		p.logs = append(p.logs, A.X{`propId`, p.PropId, val})
		p.PropId = val
		return true
	}
	return false
}

// SetCount create mutations, should not duplicate
func (p *PropLikeCountMutator) SetCount(val int64) bool { //nolint:dupl false positive
	if val != p.Count {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.logs = append(p.logs, A.X{`count`, p.Count, val})
		p.Count = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *PropLikeCountMutator) SetAll(from rqProperty.PropLikeCount, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`propId`] && (forceMap[`propId`] || from.PropId != 0) {
		p.PropId = from.PropId
		changed = true
	}
	if !excludeMap[`count`] && (forceMap[`count`] || from.Count != 0) {
		p.Count = from.Count
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyMutator DAO writer/command struct
type PropertyMutator struct {
	rqProperty.Property
	mutations []A.X
	logs      []A.X
}

// NewPropertyMutator create new ORM writer/command object
func NewPropertyMutator(adapter *Tt.Adapter) (res *PropertyMutator) {
	res = &PropertyMutator{Property: rqProperty.Property{Adapter: adapter}}
	res.Coord = []any{}
	res.PriceHistoriesSell = []any{}
	res.PriceHistoriesRent = []any{}
	res.Images = []any{}
	res.FloorList = []any{}
	return
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
//		A.X{`=`, 19, p.LastPrice},
//		A.X{`=`, 20, p.PriceHistoriesSell},
//		A.X{`=`, 21, p.PriceHistoriesRent},
//		A.X{`=`, 22, p.Purpose},
//		A.X{`=`, 23, p.HouseType},
//		A.X{`=`, 24, p.Images},
//		A.X{`=`, 25, p.Bedroom},
//		A.X{`=`, 26, p.Bathroom},
//		A.X{`=`, 27, p.AgencyFeePercent},
//		A.X{`=`, 28, p.FloorList},
//		A.X{`=`, 29, p.Country},
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
	row, err := p.Adapter.Replace(p.SpaceName(), p.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
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

// SetLastPrice create mutations, should not duplicate
func (p *PropertyMutator) SetLastPrice(val string) bool { //nolint:dupl false positive
	if val != p.LastPrice {
		p.mutations = append(p.mutations, A.X{`=`, 19, val})
		p.logs = append(p.logs, A.X{`lastPrice`, p.LastPrice, val})
		p.LastPrice = val
		return true
	}
	return false
}

// SetPriceHistoriesSell create mutations, should not duplicate
func (p *PropertyMutator) SetPriceHistoriesSell(val []any) bool { //nolint:dupl false positive
	p.mutations = append(p.mutations, A.X{`=`, 20, val})
	p.logs = append(p.logs, A.X{`priceHistoriesSell`, p.PriceHistoriesSell, val})
	p.PriceHistoriesSell = val
	return true
}

// SetPriceHistoriesRent create mutations, should not duplicate
func (p *PropertyMutator) SetPriceHistoriesRent(val []any) bool { //nolint:dupl false positive
	p.mutations = append(p.mutations, A.X{`=`, 21, val})
	p.logs = append(p.logs, A.X{`priceHistoriesRent`, p.PriceHistoriesRent, val})
	p.PriceHistoriesRent = val
	return true
}

// SetPurpose create mutations, should not duplicate
func (p *PropertyMutator) SetPurpose(val string) bool { //nolint:dupl false positive
	if val != p.Purpose {
		p.mutations = append(p.mutations, A.X{`=`, 22, val})
		p.logs = append(p.logs, A.X{`purpose`, p.Purpose, val})
		p.Purpose = val
		return true
	}
	return false
}

// SetHouseType create mutations, should not duplicate
func (p *PropertyMutator) SetHouseType(val string) bool { //nolint:dupl false positive
	if val != p.HouseType {
		p.mutations = append(p.mutations, A.X{`=`, 23, val})
		p.logs = append(p.logs, A.X{`houseType`, p.HouseType, val})
		p.HouseType = val
		return true
	}
	return false
}

// SetImages create mutations, should not duplicate
func (p *PropertyMutator) SetImages(val []any) bool { //nolint:dupl false positive
	p.mutations = append(p.mutations, A.X{`=`, 24, val})
	p.logs = append(p.logs, A.X{`images`, p.Images, val})
	p.Images = val
	return true
}

// SetBedroom create mutations, should not duplicate
func (p *PropertyMutator) SetBedroom(val int64) bool { //nolint:dupl false positive
	if val != p.Bedroom {
		p.mutations = append(p.mutations, A.X{`=`, 25, val})
		p.logs = append(p.logs, A.X{`bedroom`, p.Bedroom, val})
		p.Bedroom = val
		return true
	}
	return false
}

// SetBathroom create mutations, should not duplicate
func (p *PropertyMutator) SetBathroom(val int64) bool { //nolint:dupl false positive
	if val != p.Bathroom {
		p.mutations = append(p.mutations, A.X{`=`, 26, val})
		p.logs = append(p.logs, A.X{`bathroom`, p.Bathroom, val})
		p.Bathroom = val
		return true
	}
	return false
}

// SetAgencyFeePercent create mutations, should not duplicate
func (p *PropertyMutator) SetAgencyFeePercent(val float64) bool { //nolint:dupl false positive
	if val != p.AgencyFeePercent {
		p.mutations = append(p.mutations, A.X{`=`, 27, val})
		p.logs = append(p.logs, A.X{`agencyFeePercent`, p.AgencyFeePercent, val})
		p.AgencyFeePercent = val
		return true
	}
	return false
}

// SetFloorList create mutations, should not duplicate
func (p *PropertyMutator) SetFloorList(val []any) bool { //nolint:dupl false positive
	p.mutations = append(p.mutations, A.X{`=`, 28, val})
	p.logs = append(p.logs, A.X{`floorList`, p.FloorList, val})
	p.FloorList = val
	return true
}

// SetCountry create mutations, should not duplicate
func (p *PropertyMutator) SetCountry(val string) bool { //nolint:dupl false positive
	if val != p.Country {
		p.mutations = append(p.mutations, A.X{`=`, 29, val})
		p.logs = append(p.logs, A.X{`country`, p.Country, val})
		p.Country = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *PropertyMutator) SetAll(from rqProperty.Property, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		p.Id = from.Id
		changed = true
	}
	if !excludeMap[`uniqPropKey`] && (forceMap[`uniqPropKey`] || from.UniqPropKey != ``) {
		p.UniqPropKey = from.UniqPropKey
		changed = true
	}
	if !excludeMap[`serialNumber`] && (forceMap[`serialNumber`] || from.SerialNumber != ``) {
		p.SerialNumber = from.SerialNumber
		changed = true
	}
	if !excludeMap[`sizeM2`] && (forceMap[`sizeM2`] || from.SizeM2 != ``) {
		p.SizeM2 = from.SizeM2
		changed = true
	}
	if !excludeMap[`mainUse`] && (forceMap[`mainUse`] || from.MainUse != ``) {
		p.MainUse = from.MainUse
		changed = true
	}
	if !excludeMap[`mainBuildingMaterial`] && (forceMap[`mainBuildingMaterial`] || from.MainBuildingMaterial != ``) {
		p.MainBuildingMaterial = from.MainBuildingMaterial
		changed = true
	}
	if !excludeMap[`constructCompletedDate`] && (forceMap[`constructCompletedDate`] || from.ConstructCompletedDate != ``) {
		p.ConstructCompletedDate = from.ConstructCompletedDate
		changed = true
	}
	if !excludeMap[`numberOfFloors`] && (forceMap[`numberOfFloors`] || from.NumberOfFloors != ``) {
		p.NumberOfFloors = from.NumberOfFloors
		changed = true
	}
	if !excludeMap[`buildingLamination`] && (forceMap[`buildingLamination`] || from.BuildingLamination != ``) {
		p.BuildingLamination = from.BuildingLamination
		changed = true
	}
	if !excludeMap[`address`] && (forceMap[`address`] || from.Address != ``) {
		p.Address = from.Address
		changed = true
	}
	if !excludeMap[`district`] && (forceMap[`district`] || from.District != ``) {
		p.District = from.District
		changed = true
	}
	if !excludeMap[`note`] && (forceMap[`note`] || from.Note != ``) {
		p.Note = from.Note
		changed = true
	}
	if !excludeMap[`coord`] && (forceMap[`coord`] || from.Coord != nil) {
		p.Coord = from.Coord
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		p.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		p.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		p.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		p.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		p.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`formattedAddress`] && (forceMap[`formattedAddress`] || from.FormattedAddress != ``) {
		p.FormattedAddress = from.FormattedAddress
		changed = true
	}
	if !excludeMap[`lastPrice`] && (forceMap[`lastPrice`] || from.LastPrice != ``) {
		p.LastPrice = from.LastPrice
		changed = true
	}
	if !excludeMap[`priceHistoriesSell`] && (forceMap[`priceHistoriesSell`] || from.PriceHistoriesSell != nil) {
		p.PriceHistoriesSell = from.PriceHistoriesSell
		changed = true
	}
	if !excludeMap[`priceHistoriesRent`] && (forceMap[`priceHistoriesRent`] || from.PriceHistoriesRent != nil) {
		p.PriceHistoriesRent = from.PriceHistoriesRent
		changed = true
	}
	if !excludeMap[`purpose`] && (forceMap[`purpose`] || from.Purpose != ``) {
		p.Purpose = from.Purpose
		changed = true
	}
	if !excludeMap[`houseType`] && (forceMap[`houseType`] || from.HouseType != ``) {
		p.HouseType = from.HouseType
		changed = true
	}
	if !excludeMap[`images`] && (forceMap[`images`] || from.Images != nil) {
		p.Images = from.Images
		changed = true
	}
	if !excludeMap[`bedroom`] && (forceMap[`bedroom`] || from.Bedroom != 0) {
		p.Bedroom = from.Bedroom
		changed = true
	}
	if !excludeMap[`bathroom`] && (forceMap[`bathroom`] || from.Bathroom != 0) {
		p.Bathroom = from.Bathroom
		changed = true
	}
	if !excludeMap[`agencyFeePercent`] && (forceMap[`agencyFeePercent`] || from.AgencyFeePercent != 0) {
		p.AgencyFeePercent = from.AgencyFeePercent
		changed = true
	}
	if !excludeMap[`floorList`] && (forceMap[`floorList`] || from.FloorList != nil) {
		p.FloorList = from.FloorList
		changed = true
	}
	if !excludeMap[`country`] && (forceMap[`country`] || from.Country != ``) {
		p.Country = from.Country
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyHistoryMutator DAO writer/command struct
type PropertyHistoryMutator struct {
	rqProperty.PropertyHistory
	mutations []A.X
	logs      []A.X
}

// NewPropertyHistoryMutator create new ORM writer/command object
func NewPropertyHistoryMutator(adapter *Tt.Adapter) (res *PropertyHistoryMutator) {
	res = &PropertyHistoryMutator{PropertyHistory: rqProperty.PropertyHistory{Adapter: adapter}}
	return
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
//		A.X{`=`, 19, p.SerialNumber},
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
	row, err := p.Adapter.Replace(p.SpaceName(), p.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
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

// SetSerialNumber create mutations, should not duplicate
func (p *PropertyHistoryMutator) SetSerialNumber(val string) bool { //nolint:dupl false positive
	if val != p.SerialNumber {
		p.mutations = append(p.mutations, A.X{`=`, 19, val})
		p.logs = append(p.logs, A.X{`serialNumber`, p.SerialNumber, val})
		p.SerialNumber = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *PropertyHistoryMutator) SetAll(from rqProperty.PropertyHistory, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		p.Id = from.Id
		changed = true
	}
	if !excludeMap[`propertyKey`] && (forceMap[`propertyKey`] || from.PropertyKey != ``) {
		p.PropertyKey = from.PropertyKey
		changed = true
	}
	if !excludeMap[`transactionKey`] && (forceMap[`transactionKey`] || from.TransactionKey != ``) {
		p.TransactionKey = from.TransactionKey
		changed = true
	}
	if !excludeMap[`transactionType`] && (forceMap[`transactionType`] || from.TransactionType != ``) {
		p.TransactionType = from.TransactionType
		changed = true
	}
	if !excludeMap[`transactionSign`] && (forceMap[`transactionSign`] || from.TransactionSign != ``) {
		p.TransactionSign = from.TransactionSign
		changed = true
	}
	if !excludeMap[`transactionTime`] && (forceMap[`transactionTime`] || from.TransactionTime != ``) {
		p.TransactionTime = from.TransactionTime
		changed = true
	}
	if !excludeMap[`transactionDateNormal`] && (forceMap[`transactionDateNormal`] || from.TransactionDateNormal != ``) {
		p.TransactionDateNormal = from.TransactionDateNormal
		changed = true
	}
	if !excludeMap[`transactionNumber`] && (forceMap[`transactionNumber`] || from.TransactionNumber != ``) {
		p.TransactionNumber = from.TransactionNumber
		changed = true
	}
	if !excludeMap[`priceNtd`] && (forceMap[`priceNtd`] || from.PriceNtd != 0) {
		p.PriceNtd = from.PriceNtd
		changed = true
	}
	if !excludeMap[`pricePerUnit`] && (forceMap[`pricePerUnit`] || from.PricePerUnit != 0) {
		p.PricePerUnit = from.PricePerUnit
		changed = true
	}
	if !excludeMap[`price`] && (forceMap[`price`] || from.Price != 0) {
		p.Price = from.Price
		changed = true
	}
	if !excludeMap[`address`] && (forceMap[`address`] || from.Address != ``) {
		p.Address = from.Address
		changed = true
	}
	if !excludeMap[`district`] && (forceMap[`district`] || from.District != ``) {
		p.District = from.District
		changed = true
	}
	if !excludeMap[`note`] && (forceMap[`note`] || from.Note != ``) {
		p.Note = from.Note
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		p.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		p.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		p.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		p.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		p.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`serialNumber`] && (forceMap[`serialNumber`] || from.SerialNumber != ``) {
		p.SerialNumber = from.SerialNumber
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PropertyUsaMutator DAO writer/command struct
type PropertyUsaMutator struct {
	rqProperty.PropertyUsa
	mutations []A.X
	logs      []A.X
}

// NewPropertyUsaMutator create new ORM writer/command object
func NewPropertyUsaMutator(adapter *Tt.Adapter) (res *PropertyUsaMutator) {
	res = &PropertyUsaMutator{PropertyUsa: rqProperty.PropertyUsa{Adapter: adapter}}
	res.Coord = []any{}
	res.MediaSource = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (p *PropertyUsaMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PropertyUsaMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PropertyUsaMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *PropertyUsaMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.ToUpdateArray())
	return !L.IsError(err, `PropertyUsa.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PropertyUsaMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.mutations)
	return !L.IsError(err, `PropertyUsa.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *PropertyUsaMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id})
	return !L.IsError(err, `PropertyUsa.DoDeletePermanentById failed: `+p.SpaceName())
}

// func (p *PropertyUsaMutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := p.Adapter.Upsert(p.SpaceName(), p.ToArray(), A.X{
//		A.X{`=`, 0, p.Id},
//		A.X{`=`, 1, p.Version},
//		A.X{`=`, 2, p.PropertyId},
//		A.X{`=`, 3, p.Street},
//		A.X{`=`, 4, p.City},
//		A.X{`=`, 5, p.State},
//		A.X{`=`, 6, p.Zip},
//		A.X{`=`, 7, p.CountryCode},
//		A.X{`=`, 8, p.PropertyTypeName},
//		A.X{`=`, 9, p.YearBuilt},
//		A.X{`=`, 10, p.YearRenovated},
//		A.X{`=`, 11, p.TotalSqft},
//		A.X{`=`, 12, p.Apn},
//		A.X{`=`, 13, p.PropertyLastUpdatedDate},
//		A.X{`=`, 14, p.DisplayTimeZone},
//		A.X{`=`, 15, p.Coord},
//		A.X{`=`, 16, p.Bedroom},
//		A.X{`=`, 17, p.Bathroom},
//		A.X{`=`, 18, p.TaxableLandValue},
//		A.X{`=`, 19, p.TaxableImprovementValue},
//		A.X{`=`, 20, p.RollYear},
//		A.X{`=`, 21, p.TaxesDue},
//		A.X{`=`, 22, p.AmenitySuperGroups},
//		A.X{`=`, 23, p.CountyUrl},
//		A.X{`=`, 24, p.CountyName},
//		A.X{`=`, 25, p.CountyIsActive},
//		A.X{`=`, 26, p.MediaSource},
//		A.X{`=`, 27, p.ZoneName},
//		A.X{`=`, 28, p.ZoneType},
//		A.X{`=`, 29, p.ZoneSubType},
//		A.X{`=`, 30, p.ZoneDisplay},
//		A.X{`=`, 31, p.ZoneCode},
//		A.X{`=`, 32, p.PermittedLandUse},
//		A.X{`=`, 33, p.NotPermittedLandUse},
//		A.X{`=`, 34, p.Note},
//		A.X{`=`, 35, p.TaxNote},
//		A.X{`=`, 36, p.CreatedAt},
//		A.X{`=`, 37, p.CreatedBy},
//		A.X{`=`, 38, p.UpdatedAt},
//		A.X{`=`, 39, p.UpdatedBy},
//		A.X{`=`, 40, p.DeletedAt},
//		A.X{`=`, 41, p.ListingBrokerName},
//		A.X{`=`, 42, p.ListingBrokerNumber},
//		A.X{`=`, 43, p.ListingAgentName},
//		A.X{`=`, 44, p.ListingAgentNumber},
//		A.X{`=`, 45, p.AgencyFeePercent},
//		A.X{`=`, 46, p.Country},
//	})
//	return !L.IsError(err, `PropertyUsa.DoUpsert failed: `+p.SpaceName())
// }

// DoOverwriteByPropertyId update all columns, error if not exists, not using mutations/Set*
func (p *PropertyUsaMutator) DoOverwriteByPropertyId() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexPropertyId(), A.X{p.PropertyId}, p.ToUpdateArray())
	return !L.IsError(err, `PropertyUsa.DoOverwriteByPropertyId failed: `+p.SpaceName())
}

// DoUpdateByPropertyId update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PropertyUsaMutator) DoUpdateByPropertyId() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexPropertyId(), A.X{p.PropertyId}, p.mutations)
	return !L.IsError(err, `PropertyUsa.DoUpdateByPropertyId failed: `+p.SpaceName())
}

// DoDeletePermanentByPropertyId permanent delete
func (p *PropertyUsaMutator) DoDeletePermanentByPropertyId() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexPropertyId(), A.X{p.PropertyId})
	return !L.IsError(err, `PropertyUsa.DoDeletePermanentByPropertyId failed: `+p.SpaceName())
}

// DoInsert insert, error if already exists
func (p *PropertyUsaMutator) DoInsert() bool { //nolint:dupl false positive
	row, err := p.Adapter.Insert(p.SpaceName(), p.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `PropertyUsa.DoInsert failed: `+p.SpaceName())
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (p *PropertyUsaMutator) DoUpsert() bool { //nolint:dupl false positive
	row, err := p.Adapter.Replace(p.SpaceName(), p.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `PropertyUsa.DoUpsert failed: `+p.SpaceName())
}

// SetId create mutations, should not duplicate
func (p *PropertyUsaMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations = append(p.mutations, A.X{`=`, 0, val})
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetVersion create mutations, should not duplicate
func (p *PropertyUsaMutator) SetVersion(val float64) bool { //nolint:dupl false positive
	if val != p.Version {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.logs = append(p.logs, A.X{`version`, p.Version, val})
		p.Version = val
		return true
	}
	return false
}

// SetPropertyId create mutations, should not duplicate
func (p *PropertyUsaMutator) SetPropertyId(val uint64) bool { //nolint:dupl false positive
	if val != p.PropertyId {
		p.mutations = append(p.mutations, A.X{`=`, 2, val})
		p.logs = append(p.logs, A.X{`propertyId`, p.PropertyId, val})
		p.PropertyId = val
		return true
	}
	return false
}

// SetStreet create mutations, should not duplicate
func (p *PropertyUsaMutator) SetStreet(val string) bool { //nolint:dupl false positive
	if val != p.Street {
		p.mutations = append(p.mutations, A.X{`=`, 3, val})
		p.logs = append(p.logs, A.X{`street`, p.Street, val})
		p.Street = val
		return true
	}
	return false
}

// SetCity create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCity(val string) bool { //nolint:dupl false positive
	if val != p.City {
		p.mutations = append(p.mutations, A.X{`=`, 4, val})
		p.logs = append(p.logs, A.X{`city`, p.City, val})
		p.City = val
		return true
	}
	return false
}

// SetState create mutations, should not duplicate
func (p *PropertyUsaMutator) SetState(val string) bool { //nolint:dupl false positive
	if val != p.State {
		p.mutations = append(p.mutations, A.X{`=`, 5, val})
		p.logs = append(p.logs, A.X{`state`, p.State, val})
		p.State = val
		return true
	}
	return false
}

// SetZip create mutations, should not duplicate
func (p *PropertyUsaMutator) SetZip(val string) bool { //nolint:dupl false positive
	if val != p.Zip {
		p.mutations = append(p.mutations, A.X{`=`, 6, val})
		p.logs = append(p.logs, A.X{`zip`, p.Zip, val})
		p.Zip = val
		return true
	}
	return false
}

// SetCountryCode create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCountryCode(val string) bool { //nolint:dupl false positive
	if val != p.CountryCode {
		p.mutations = append(p.mutations, A.X{`=`, 7, val})
		p.logs = append(p.logs, A.X{`countryCode`, p.CountryCode, val})
		p.CountryCode = val
		return true
	}
	return false
}

// SetPropertyTypeName create mutations, should not duplicate
func (p *PropertyUsaMutator) SetPropertyTypeName(val string) bool { //nolint:dupl false positive
	if val != p.PropertyTypeName {
		p.mutations = append(p.mutations, A.X{`=`, 8, val})
		p.logs = append(p.logs, A.X{`propertyTypeName`, p.PropertyTypeName, val})
		p.PropertyTypeName = val
		return true
	}
	return false
}

// SetYearBuilt create mutations, should not duplicate
func (p *PropertyUsaMutator) SetYearBuilt(val int64) bool { //nolint:dupl false positive
	if val != p.YearBuilt {
		p.mutations = append(p.mutations, A.X{`=`, 9, val})
		p.logs = append(p.logs, A.X{`yearBuilt`, p.YearBuilt, val})
		p.YearBuilt = val
		return true
	}
	return false
}

// SetYearRenovated create mutations, should not duplicate
func (p *PropertyUsaMutator) SetYearRenovated(val int64) bool { //nolint:dupl false positive
	if val != p.YearRenovated {
		p.mutations = append(p.mutations, A.X{`=`, 10, val})
		p.logs = append(p.logs, A.X{`yearRenovated`, p.YearRenovated, val})
		p.YearRenovated = val
		return true
	}
	return false
}

// SetTotalSqft create mutations, should not duplicate
func (p *PropertyUsaMutator) SetTotalSqft(val float64) bool { //nolint:dupl false positive
	if val != p.TotalSqft {
		p.mutations = append(p.mutations, A.X{`=`, 11, val})
		p.logs = append(p.logs, A.X{`totalSqft`, p.TotalSqft, val})
		p.TotalSqft = val
		return true
	}
	return false
}

// SetApn create mutations, should not duplicate
func (p *PropertyUsaMutator) SetApn(val string) bool { //nolint:dupl false positive
	if val != p.Apn {
		p.mutations = append(p.mutations, A.X{`=`, 12, val})
		p.logs = append(p.logs, A.X{`apn`, p.Apn, val})
		p.Apn = val
		return true
	}
	return false
}

// SetPropertyLastUpdatedDate create mutations, should not duplicate
func (p *PropertyUsaMutator) SetPropertyLastUpdatedDate(val int64) bool { //nolint:dupl false positive
	if val != p.PropertyLastUpdatedDate {
		p.mutations = append(p.mutations, A.X{`=`, 13, val})
		p.logs = append(p.logs, A.X{`propertyLastUpdatedDate`, p.PropertyLastUpdatedDate, val})
		p.PropertyLastUpdatedDate = val
		return true
	}
	return false
}

// SetDisplayTimeZone create mutations, should not duplicate
func (p *PropertyUsaMutator) SetDisplayTimeZone(val string) bool { //nolint:dupl false positive
	if val != p.DisplayTimeZone {
		p.mutations = append(p.mutations, A.X{`=`, 14, val})
		p.logs = append(p.logs, A.X{`displayTimeZone`, p.DisplayTimeZone, val})
		p.DisplayTimeZone = val
		return true
	}
	return false
}

// SetCoord create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCoord(val []any) bool { //nolint:dupl false positive
	p.mutations = append(p.mutations, A.X{`=`, 15, val})
	p.logs = append(p.logs, A.X{`coord`, p.Coord, val})
	p.Coord = val
	return true
}

// SetBedroom create mutations, should not duplicate
func (p *PropertyUsaMutator) SetBedroom(val int64) bool { //nolint:dupl false positive
	if val != p.Bedroom {
		p.mutations = append(p.mutations, A.X{`=`, 16, val})
		p.logs = append(p.logs, A.X{`bedroom`, p.Bedroom, val})
		p.Bedroom = val
		return true
	}
	return false
}

// SetBathroom create mutations, should not duplicate
func (p *PropertyUsaMutator) SetBathroom(val int64) bool { //nolint:dupl false positive
	if val != p.Bathroom {
		p.mutations = append(p.mutations, A.X{`=`, 17, val})
		p.logs = append(p.logs, A.X{`bathroom`, p.Bathroom, val})
		p.Bathroom = val
		return true
	}
	return false
}

// SetTaxableLandValue create mutations, should not duplicate
func (p *PropertyUsaMutator) SetTaxableLandValue(val int64) bool { //nolint:dupl false positive
	if val != p.TaxableLandValue {
		p.mutations = append(p.mutations, A.X{`=`, 18, val})
		p.logs = append(p.logs, A.X{`taxableLandValue`, p.TaxableLandValue, val})
		p.TaxableLandValue = val
		return true
	}
	return false
}

// SetTaxableImprovementValue create mutations, should not duplicate
func (p *PropertyUsaMutator) SetTaxableImprovementValue(val int64) bool { //nolint:dupl false positive
	if val != p.TaxableImprovementValue {
		p.mutations = append(p.mutations, A.X{`=`, 19, val})
		p.logs = append(p.logs, A.X{`taxableImprovementValue`, p.TaxableImprovementValue, val})
		p.TaxableImprovementValue = val
		return true
	}
	return false
}

// SetRollYear create mutations, should not duplicate
func (p *PropertyUsaMutator) SetRollYear(val int64) bool { //nolint:dupl false positive
	if val != p.RollYear {
		p.mutations = append(p.mutations, A.X{`=`, 20, val})
		p.logs = append(p.logs, A.X{`rollYear`, p.RollYear, val})
		p.RollYear = val
		return true
	}
	return false
}

// SetTaxesDue create mutations, should not duplicate
func (p *PropertyUsaMutator) SetTaxesDue(val float64) bool { //nolint:dupl false positive
	if val != p.TaxesDue {
		p.mutations = append(p.mutations, A.X{`=`, 21, val})
		p.logs = append(p.logs, A.X{`taxesDue`, p.TaxesDue, val})
		p.TaxesDue = val
		return true
	}
	return false
}

// SetAmenitySuperGroups create mutations, should not duplicate
func (p *PropertyUsaMutator) SetAmenitySuperGroups(val string) bool { //nolint:dupl false positive
	if val != p.AmenitySuperGroups {
		p.mutations = append(p.mutations, A.X{`=`, 22, val})
		p.logs = append(p.logs, A.X{`amenitySuperGroups`, p.AmenitySuperGroups, val})
		p.AmenitySuperGroups = val
		return true
	}
	return false
}

// SetCountyUrl create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCountyUrl(val string) bool { //nolint:dupl false positive
	if val != p.CountyUrl {
		p.mutations = append(p.mutations, A.X{`=`, 23, val})
		p.logs = append(p.logs, A.X{`countyUrl`, p.CountyUrl, val})
		p.CountyUrl = val
		return true
	}
	return false
}

// SetCountyName create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCountyName(val string) bool { //nolint:dupl false positive
	if val != p.CountyName {
		p.mutations = append(p.mutations, A.X{`=`, 24, val})
		p.logs = append(p.logs, A.X{`countyName`, p.CountyName, val})
		p.CountyName = val
		return true
	}
	return false
}

// SetCountyIsActive create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCountyIsActive(val bool) bool { //nolint:dupl false positive
	if val != p.CountyIsActive {
		p.mutations = append(p.mutations, A.X{`=`, 25, val})
		p.logs = append(p.logs, A.X{`countyIsActive`, p.CountyIsActive, val})
		p.CountyIsActive = val
		return true
	}
	return false
}

// SetMediaSource create mutations, should not duplicate
func (p *PropertyUsaMutator) SetMediaSource(val []any) bool { //nolint:dupl false positive
	p.mutations = append(p.mutations, A.X{`=`, 26, val})
	p.logs = append(p.logs, A.X{`mediaSource`, p.MediaSource, val})
	p.MediaSource = val
	return true
}

// SetZoneName create mutations, should not duplicate
func (p *PropertyUsaMutator) SetZoneName(val string) bool { //nolint:dupl false positive
	if val != p.ZoneName {
		p.mutations = append(p.mutations, A.X{`=`, 27, val})
		p.logs = append(p.logs, A.X{`zoneName`, p.ZoneName, val})
		p.ZoneName = val
		return true
	}
	return false
}

// SetZoneType create mutations, should not duplicate
func (p *PropertyUsaMutator) SetZoneType(val string) bool { //nolint:dupl false positive
	if val != p.ZoneType {
		p.mutations = append(p.mutations, A.X{`=`, 28, val})
		p.logs = append(p.logs, A.X{`zoneType`, p.ZoneType, val})
		p.ZoneType = val
		return true
	}
	return false
}

// SetZoneSubType create mutations, should not duplicate
func (p *PropertyUsaMutator) SetZoneSubType(val string) bool { //nolint:dupl false positive
	if val != p.ZoneSubType {
		p.mutations = append(p.mutations, A.X{`=`, 29, val})
		p.logs = append(p.logs, A.X{`zoneSubType`, p.ZoneSubType, val})
		p.ZoneSubType = val
		return true
	}
	return false
}

// SetZoneDisplay create mutations, should not duplicate
func (p *PropertyUsaMutator) SetZoneDisplay(val string) bool { //nolint:dupl false positive
	if val != p.ZoneDisplay {
		p.mutations = append(p.mutations, A.X{`=`, 30, val})
		p.logs = append(p.logs, A.X{`zoneDisplay`, p.ZoneDisplay, val})
		p.ZoneDisplay = val
		return true
	}
	return false
}

// SetZoneCode create mutations, should not duplicate
func (p *PropertyUsaMutator) SetZoneCode(val string) bool { //nolint:dupl false positive
	if val != p.ZoneCode {
		p.mutations = append(p.mutations, A.X{`=`, 31, val})
		p.logs = append(p.logs, A.X{`zoneCode`, p.ZoneCode, val})
		p.ZoneCode = val
		return true
	}
	return false
}

// SetPermittedLandUse create mutations, should not duplicate
func (p *PropertyUsaMutator) SetPermittedLandUse(val string) bool { //nolint:dupl false positive
	if val != p.PermittedLandUse {
		p.mutations = append(p.mutations, A.X{`=`, 32, val})
		p.logs = append(p.logs, A.X{`permittedLandUse`, p.PermittedLandUse, val})
		p.PermittedLandUse = val
		return true
	}
	return false
}

// SetNotPermittedLandUse create mutations, should not duplicate
func (p *PropertyUsaMutator) SetNotPermittedLandUse(val string) bool { //nolint:dupl false positive
	if val != p.NotPermittedLandUse {
		p.mutations = append(p.mutations, A.X{`=`, 33, val})
		p.logs = append(p.logs, A.X{`notPermittedLandUse`, p.NotPermittedLandUse, val})
		p.NotPermittedLandUse = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (p *PropertyUsaMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != p.Note {
		p.mutations = append(p.mutations, A.X{`=`, 34, val})
		p.logs = append(p.logs, A.X{`note`, p.Note, val})
		p.Note = val
		return true
	}
	return false
}

// SetTaxNote create mutations, should not duplicate
func (p *PropertyUsaMutator) SetTaxNote(val string) bool { //nolint:dupl false positive
	if val != p.TaxNote {
		p.mutations = append(p.mutations, A.X{`=`, 35, val})
		p.logs = append(p.logs, A.X{`taxNote`, p.TaxNote, val})
		p.TaxNote = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 36, val})
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 37, val})
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PropertyUsaMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 38, val})
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PropertyUsaMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 39, val})
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PropertyUsaMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations = append(p.mutations, A.X{`=`, 40, val})
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetListingBrokerName create mutations, should not duplicate
func (p *PropertyUsaMutator) SetListingBrokerName(val string) bool { //nolint:dupl false positive
	if val != p.ListingBrokerName {
		p.mutations = append(p.mutations, A.X{`=`, 41, val})
		p.logs = append(p.logs, A.X{`listingBrokerName`, p.ListingBrokerName, val})
		p.ListingBrokerName = val
		return true
	}
	return false
}

// SetListingBrokerNumber create mutations, should not duplicate
func (p *PropertyUsaMutator) SetListingBrokerNumber(val string) bool { //nolint:dupl false positive
	if val != p.ListingBrokerNumber {
		p.mutations = append(p.mutations, A.X{`=`, 42, val})
		p.logs = append(p.logs, A.X{`listingBrokerNumber`, p.ListingBrokerNumber, val})
		p.ListingBrokerNumber = val
		return true
	}
	return false
}

// SetListingAgentName create mutations, should not duplicate
func (p *PropertyUsaMutator) SetListingAgentName(val string) bool { //nolint:dupl false positive
	if val != p.ListingAgentName {
		p.mutations = append(p.mutations, A.X{`=`, 43, val})
		p.logs = append(p.logs, A.X{`listingAgentName`, p.ListingAgentName, val})
		p.ListingAgentName = val
		return true
	}
	return false
}

// SetListingAgentNumber create mutations, should not duplicate
func (p *PropertyUsaMutator) SetListingAgentNumber(val string) bool { //nolint:dupl false positive
	if val != p.ListingAgentNumber {
		p.mutations = append(p.mutations, A.X{`=`, 44, val})
		p.logs = append(p.logs, A.X{`listingAgentNumber`, p.ListingAgentNumber, val})
		p.ListingAgentNumber = val
		return true
	}
	return false
}

// SetAgencyFeePercent create mutations, should not duplicate
func (p *PropertyUsaMutator) SetAgencyFeePercent(val float64) bool { //nolint:dupl false positive
	if val != p.AgencyFeePercent {
		p.mutations = append(p.mutations, A.X{`=`, 45, val})
		p.logs = append(p.logs, A.X{`agencyFeePercent`, p.AgencyFeePercent, val})
		p.AgencyFeePercent = val
		return true
	}
	return false
}

// SetCountry create mutations, should not duplicate
func (p *PropertyUsaMutator) SetCountry(val string) bool { //nolint:dupl false positive
	if val != p.Country {
		p.mutations = append(p.mutations, A.X{`=`, 46, val})
		p.logs = append(p.logs, A.X{`country`, p.Country, val})
		p.Country = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *PropertyUsaMutator) SetAll(from rqProperty.PropertyUsa, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		p.Id = from.Id
		changed = true
	}
	if !excludeMap[`version`] && (forceMap[`version`] || from.Version != 0) {
		p.Version = from.Version
		changed = true
	}
	if !excludeMap[`propertyId`] && (forceMap[`propertyId`] || from.PropertyId != 0) {
		p.PropertyId = from.PropertyId
		changed = true
	}
	if !excludeMap[`street`] && (forceMap[`street`] || from.Street != ``) {
		p.Street = from.Street
		changed = true
	}
	if !excludeMap[`city`] && (forceMap[`city`] || from.City != ``) {
		p.City = from.City
		changed = true
	}
	if !excludeMap[`state`] && (forceMap[`state`] || from.State != ``) {
		p.State = from.State
		changed = true
	}
	if !excludeMap[`zip`] && (forceMap[`zip`] || from.Zip != ``) {
		p.Zip = from.Zip
		changed = true
	}
	if !excludeMap[`countryCode`] && (forceMap[`countryCode`] || from.CountryCode != ``) {
		p.CountryCode = from.CountryCode
		changed = true
	}
	if !excludeMap[`propertyTypeName`] && (forceMap[`propertyTypeName`] || from.PropertyTypeName != ``) {
		p.PropertyTypeName = from.PropertyTypeName
		changed = true
	}
	if !excludeMap[`yearBuilt`] && (forceMap[`yearBuilt`] || from.YearBuilt != 0) {
		p.YearBuilt = from.YearBuilt
		changed = true
	}
	if !excludeMap[`yearRenovated`] && (forceMap[`yearRenovated`] || from.YearRenovated != 0) {
		p.YearRenovated = from.YearRenovated
		changed = true
	}
	if !excludeMap[`totalSqft`] && (forceMap[`totalSqft`] || from.TotalSqft != 0) {
		p.TotalSqft = from.TotalSqft
		changed = true
	}
	if !excludeMap[`apn`] && (forceMap[`apn`] || from.Apn != ``) {
		p.Apn = from.Apn
		changed = true
	}
	if !excludeMap[`propertyLastUpdatedDate`] && (forceMap[`propertyLastUpdatedDate`] || from.PropertyLastUpdatedDate != 0) {
		p.PropertyLastUpdatedDate = from.PropertyLastUpdatedDate
		changed = true
	}
	if !excludeMap[`displayTimeZone`] && (forceMap[`displayTimeZone`] || from.DisplayTimeZone != ``) {
		p.DisplayTimeZone = from.DisplayTimeZone
		changed = true
	}
	if !excludeMap[`coord`] && (forceMap[`coord`] || from.Coord != nil) {
		p.Coord = from.Coord
		changed = true
	}
	if !excludeMap[`bedroom`] && (forceMap[`bedroom`] || from.Bedroom != 0) {
		p.Bedroom = from.Bedroom
		changed = true
	}
	if !excludeMap[`bathroom`] && (forceMap[`bathroom`] || from.Bathroom != 0) {
		p.Bathroom = from.Bathroom
		changed = true
	}
	if !excludeMap[`taxableLandValue`] && (forceMap[`taxableLandValue`] || from.TaxableLandValue != 0) {
		p.TaxableLandValue = from.TaxableLandValue
		changed = true
	}
	if !excludeMap[`taxableImprovementValue`] && (forceMap[`taxableImprovementValue`] || from.TaxableImprovementValue != 0) {
		p.TaxableImprovementValue = from.TaxableImprovementValue
		changed = true
	}
	if !excludeMap[`rollYear`] && (forceMap[`rollYear`] || from.RollYear != 0) {
		p.RollYear = from.RollYear
		changed = true
	}
	if !excludeMap[`taxesDue`] && (forceMap[`taxesDue`] || from.TaxesDue != 0) {
		p.TaxesDue = from.TaxesDue
		changed = true
	}
	if !excludeMap[`amenitySuperGroups`] && (forceMap[`amenitySuperGroups`] || from.AmenitySuperGroups != ``) {
		p.AmenitySuperGroups = from.AmenitySuperGroups
		changed = true
	}
	if !excludeMap[`countyUrl`] && (forceMap[`countyUrl`] || from.CountyUrl != ``) {
		p.CountyUrl = from.CountyUrl
		changed = true
	}
	if !excludeMap[`countyName`] && (forceMap[`countyName`] || from.CountyName != ``) {
		p.CountyName = from.CountyName
		changed = true
	}
	if !excludeMap[`countyIsActive`] && (forceMap[`countyIsActive`] || from.CountyIsActive != false) {
		p.CountyIsActive = from.CountyIsActive
		changed = true
	}
	if !excludeMap[`mediaSource`] && (forceMap[`mediaSource`] || from.MediaSource != nil) {
		p.MediaSource = from.MediaSource
		changed = true
	}
	if !excludeMap[`zoneName`] && (forceMap[`zoneName`] || from.ZoneName != ``) {
		p.ZoneName = from.ZoneName
		changed = true
	}
	if !excludeMap[`zoneType`] && (forceMap[`zoneType`] || from.ZoneType != ``) {
		p.ZoneType = from.ZoneType
		changed = true
	}
	if !excludeMap[`zoneSubType`] && (forceMap[`zoneSubType`] || from.ZoneSubType != ``) {
		p.ZoneSubType = from.ZoneSubType
		changed = true
	}
	if !excludeMap[`zoneDisplay`] && (forceMap[`zoneDisplay`] || from.ZoneDisplay != ``) {
		p.ZoneDisplay = from.ZoneDisplay
		changed = true
	}
	if !excludeMap[`zoneCode`] && (forceMap[`zoneCode`] || from.ZoneCode != ``) {
		p.ZoneCode = from.ZoneCode
		changed = true
	}
	if !excludeMap[`permittedLandUse`] && (forceMap[`permittedLandUse`] || from.PermittedLandUse != ``) {
		p.PermittedLandUse = from.PermittedLandUse
		changed = true
	}
	if !excludeMap[`notPermittedLandUse`] && (forceMap[`notPermittedLandUse`] || from.NotPermittedLandUse != ``) {
		p.NotPermittedLandUse = from.NotPermittedLandUse
		changed = true
	}
	if !excludeMap[`note`] && (forceMap[`note`] || from.Note != ``) {
		p.Note = from.Note
		changed = true
	}
	if !excludeMap[`taxNote`] && (forceMap[`taxNote`] || from.TaxNote != ``) {
		p.TaxNote = from.TaxNote
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		p.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		p.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		p.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		p.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		p.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`listingBrokerName`] && (forceMap[`listingBrokerName`] || from.ListingBrokerName != ``) {
		p.ListingBrokerName = from.ListingBrokerName
		changed = true
	}
	if !excludeMap[`listingBrokerNumber`] && (forceMap[`listingBrokerNumber`] || from.ListingBrokerNumber != ``) {
		p.ListingBrokerNumber = from.ListingBrokerNumber
		changed = true
	}
	if !excludeMap[`listingAgentName`] && (forceMap[`listingAgentName`] || from.ListingAgentName != ``) {
		p.ListingAgentName = from.ListingAgentName
		changed = true
	}
	if !excludeMap[`listingAgentNumber`] && (forceMap[`listingAgentNumber`] || from.ListingAgentNumber != ``) {
		p.ListingAgentNumber = from.ListingAgentNumber
		changed = true
	}
	if !excludeMap[`agencyFeePercent`] && (forceMap[`agencyFeePercent`] || from.AgencyFeePercent != 0) {
		p.AgencyFeePercent = from.AgencyFeePercent
		changed = true
	}
	if !excludeMap[`country`] && (forceMap[`country`] || from.Country != ``) {
		p.Country = from.Country
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// UserPropLikesMutator DAO writer/command struct
type UserPropLikesMutator struct {
	rqProperty.UserPropLikes
	mutations []A.X
	logs      []A.X
}

// NewUserPropLikesMutator create new ORM writer/command object
func NewUserPropLikesMutator(adapter *Tt.Adapter) (res *UserPropLikesMutator) {
	res = &UserPropLikesMutator{UserPropLikes: rqProperty.UserPropLikes{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (u *UserPropLikesMutator) Logs() []A.X { //nolint:dupl false positive
	return u.logs
}

// HaveMutation check whether Set* methods ever called
func (u *UserPropLikesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(u.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (u *UserPropLikesMutator) ClearMutations() { //nolint:dupl false positive
	u.mutations = []A.X{}
	u.logs = []A.X{}
}

// func (u *UserPropLikesMutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := u.Adapter.Upsert(u.SpaceName(), u.ToArray(), A.X{
//		A.X{`=`, 0, u.PropId},
//		A.X{`=`, 1, u.UserId},
//		A.X{`=`, 2, u.CreatedAt},
//	})
//	return !L.IsError(err, `UserPropLikes.DoUpsert failed: `+u.SpaceName())
// }

// DoOverwriteByUserIdPropId update all columns, error if not exists, not using mutations/Set*
func (u *UserPropLikesMutator) DoOverwriteByUserIdPropId() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexUserIdPropId(), A.X{u.UserId, u.PropId}, u.ToUpdateArray())
	return !L.IsError(err, `UserPropLikes.DoOverwriteByUserIdPropId failed: `+u.SpaceName())
}

// DoUpdateByUserIdPropId update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UserPropLikesMutator) DoUpdateByUserIdPropId() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexUserIdPropId(), A.X{u.UserId, u.PropId}, u.mutations)
	return !L.IsError(err, `UserPropLikes.DoUpdateByUserIdPropId failed: `+u.SpaceName())
}

// DoDeletePermanentByUserIdPropId permanent delete
func (u *UserPropLikesMutator) DoDeletePermanentByUserIdPropId() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexUserIdPropId(), A.X{u.UserId, u.PropId})
	return !L.IsError(err, `UserPropLikes.DoDeletePermanentByUserIdPropId failed: `+u.SpaceName())
}

// DoInsert insert, error if already exists
func (u *UserPropLikesMutator) DoInsert() bool { //nolint:dupl false positive
	_, err := u.Adapter.Insert(u.SpaceName(), u.ToArray())
	return !L.IsError(err, `UserPropLikes.DoInsert failed: `+u.SpaceName())
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (u *UserPropLikesMutator) DoUpsert() bool { //nolint:dupl false positive
	_, err := u.Adapter.Replace(u.SpaceName(), u.ToArray())
	return !L.IsError(err, `UserPropLikes.DoUpsert failed: `+u.SpaceName())
}

// SetPropId create mutations, should not duplicate
func (u *UserPropLikesMutator) SetPropId(val uint64) bool { //nolint:dupl false positive
	if val != u.PropId {
		u.mutations = append(u.mutations, A.X{`=`, 0, val})
		u.logs = append(u.logs, A.X{`propId`, u.PropId, val})
		u.PropId = val
		return true
	}
	return false
}

// SetUserId create mutations, should not duplicate
func (u *UserPropLikesMutator) SetUserId(val uint64) bool { //nolint:dupl false positive
	if val != u.UserId {
		u.mutations = append(u.mutations, A.X{`=`, 1, val})
		u.logs = append(u.logs, A.X{`userId`, u.UserId, val})
		u.UserId = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (u *UserPropLikesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != u.CreatedAt {
		u.mutations = append(u.mutations, A.X{`=`, 2, val})
		u.logs = append(u.logs, A.X{`createdAt`, u.CreatedAt, val})
		u.CreatedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (u *UserPropLikesMutator) SetAll(from rqProperty.UserPropLikes, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`propId`] && (forceMap[`propId`] || from.PropId != 0) {
		u.PropId = from.PropId
		changed = true
	}
	if !excludeMap[`userId`] && (forceMap[`userId`] || from.UserId != 0) {
		u.UserId = from.UserId
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		u.CreatedAt = from.CreatedAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
