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
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type wcProperty__ORM.GEN.go
type PropertyMutator struct {
	rqProperty.Property
	mutations []A.X
}

// NewPropertyMutator create new ORM writer/command object
func NewPropertyMutator(adapter *Tt.Adapter) *PropertyMutator {
	return &PropertyMutator{Property: rqProperty.Property{Adapter: adapter}}
}

// HaveMutation check whether Set* methods ever called
func (p *PropertyMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PropertyMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
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
//		A.X{`=`, 1, p.SerialNumber},
//		A.X{`=`, 2, p.SizeM2},
//		A.X{`=`, 3, p.MainUse},
//		A.X{`=`, 4, p.MainBuildingMaterial},
//		A.X{`=`, 5, p.ConstructCompletedDate},
//		A.X{`=`, 6, p.NumberOfFloors},
//		A.X{`=`, 7, p.BuildingLamination},
//		A.X{`=`, 8, p.Note},
//		A.X{`=`, 9, p.CreatedAt},
//		A.X{`=`, 10, p.CreatedBy},
//		A.X{`=`, 11, p.UpdatedAt},
//		A.X{`=`, 12, p.UpdatedBy},
//		A.X{`=`, 13, p.DeletedAt},
//	})
//	return !L.IsError(err, `Property.DoUpsert failed: `+p.SpaceName())
// }

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
		p.Id = val
		return true
	}
	return false
}

// SetSerialNumber create mutations, should not duplicate
func (p *PropertyMutator) SetSerialNumber(val string) bool { //nolint:dupl false positive
	if val != p.SerialNumber {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.SerialNumber = val
		return true
	}
	return false
}

// SetSizeM2 create mutations, should not duplicate
func (p *PropertyMutator) SetSizeM2(val float64) bool { //nolint:dupl false positive
	if val != p.SizeM2 {
		p.mutations = append(p.mutations, A.X{`=`, 2, val})
		p.SizeM2 = val
		return true
	}
	return false
}

// SetMainUse create mutations, should not duplicate
func (p *PropertyMutator) SetMainUse(val string) bool { //nolint:dupl false positive
	if val != p.MainUse {
		p.mutations = append(p.mutations, A.X{`=`, 3, val})
		p.MainUse = val
		return true
	}
	return false
}

// SetMainBuildingMaterial create mutations, should not duplicate
func (p *PropertyMutator) SetMainBuildingMaterial(val string) bool { //nolint:dupl false positive
	if val != p.MainBuildingMaterial {
		p.mutations = append(p.mutations, A.X{`=`, 4, val})
		p.MainBuildingMaterial = val
		return true
	}
	return false
}

// SetConstructCompletedDate create mutations, should not duplicate
func (p *PropertyMutator) SetConstructCompletedDate(val int64) bool { //nolint:dupl false positive
	if val != p.ConstructCompletedDate {
		p.mutations = append(p.mutations, A.X{`=`, 5, val})
		p.ConstructCompletedDate = val
		return true
	}
	return false
}

// SetNumberOfFloors create mutations, should not duplicate
func (p *PropertyMutator) SetNumberOfFloors(val float64) bool { //nolint:dupl false positive
	if val != p.NumberOfFloors {
		p.mutations = append(p.mutations, A.X{`=`, 6, val})
		p.NumberOfFloors = val
		return true
	}
	return false
}

// SetBuildingLamination create mutations, should not duplicate
func (p *PropertyMutator) SetBuildingLamination(val string) bool { //nolint:dupl false positive
	if val != p.BuildingLamination {
		p.mutations = append(p.mutations, A.X{`=`, 7, val})
		p.BuildingLamination = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (p *PropertyMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != p.Note {
		p.mutations = append(p.mutations, A.X{`=`, 8, val})
		p.Note = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *PropertyMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 9, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PropertyMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 10, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PropertyMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 11, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PropertyMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 12, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PropertyMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations = append(p.mutations, A.X{`=`, 13, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
