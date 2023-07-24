package wcStorage

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mStorage/rqStorage"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// FilesMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcStorage__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcStorage__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcStorage__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcStorage__ORM.GEN.go
type FilesMutator struct {
	rqStorage.Files
	mutations []A.X
	logs      []A.X
}

// NewFilesMutator create new ORM writer/command object
func NewFilesMutator(adapter *Tt.Adapter) *FilesMutator {
	return &FilesMutator{Files: rqStorage.Files{Adapter: adapter}}
}

// Logs get array of logs [field, old, new]
func (f *FilesMutator) Logs() []A.X { //nolint:dupl false positive
	return f.logs
}

// HaveMutation check whether Set* methods ever called
func (f *FilesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(f.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (f *FilesMutator) ClearMutations() { //nolint:dupl false positive
	f.mutations = []A.X{}
	f.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (f *FilesMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.ToUpdateArray())
	return !L.IsError(err, `Files.DoOverwriteById failed: `+f.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (f *FilesMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !f.HaveMutation() {
		return true
	}
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.mutations)
	return !L.IsError(err, `Files.DoUpdateById failed: `+f.SpaceName())
}

// DoDeletePermanentById permanent delete
func (f *FilesMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Delete(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id})
	return !L.IsError(err, `Files.DoDeletePermanentById failed: `+f.SpaceName())
}

// func (f *FilesMutator) DoUpsert() bool { //nolint:dupl false positive
//	_, err := f.Adapter.Upsert(f.SpaceName(), f.ToArray(), A.X{
//		A.X{`=`, 0, f.Id},
//		A.X{`=`, 1, f.CreatedAt},
//		A.X{`=`, 2, f.CreatedBy},
//		A.X{`=`, 3, f.Mime},
//		A.X{`=`, 4, f.Purpose},
//		A.X{`=`, 5, f.RefId},
//		A.X{`=`, 6, f.AccessCount},
//		A.X{`=`, 7, f.LastAccessAt},
//		A.X{`=`, 8, f.OriginalPath},
//		A.X{`=`, 9, f.OriginalSize},
//		A.X{`=`, 10, f.ResizedPath},
//		A.X{`=`, 11, f.ResizedSize},
//	})
//	return !L.IsError(err, `Files.DoUpsert failed: `+f.SpaceName())
// }

// DoOverwriteByOriginalPath update all columns, error if not exists, not using mutations/Set*
func (f *FilesMutator) DoOverwriteByOriginalPath() bool { //nolint:dupl false positive
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexOriginalPath(), A.X{f.OriginalPath}, f.ToUpdateArray())
	return !L.IsError(err, `Files.DoOverwriteByOriginalPath failed: `+f.SpaceName())
}

// DoUpdateByOriginalPath update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (f *FilesMutator) DoUpdateByOriginalPath() bool { //nolint:dupl false positive
	if !f.HaveMutation() {
		return true
	}
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexOriginalPath(), A.X{f.OriginalPath}, f.mutations)
	return !L.IsError(err, `Files.DoUpdateByOriginalPath failed: `+f.SpaceName())
}

// DoDeletePermanentByOriginalPath permanent delete
func (f *FilesMutator) DoDeletePermanentByOriginalPath() bool { //nolint:dupl false positive
	_, err := f.Adapter.Delete(f.SpaceName(), f.UniqueIndexOriginalPath(), A.X{f.OriginalPath})
	return !L.IsError(err, `Files.DoDeletePermanentByOriginalPath failed: `+f.SpaceName())
}

// DoInsert insert, error if already exists
func (f *FilesMutator) DoInsert() bool { //nolint:dupl false positive
	row, err := f.Adapter.Insert(f.SpaceName(), f.ToArray())
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Files.DoInsert failed: `+f.SpaceName())
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (f *FilesMutator) DoUpsert() bool { //nolint:dupl false positive
	_, err := f.Adapter.Replace(f.SpaceName(), f.ToArray())
	return !L.IsError(err, `Files.DoUpsert failed: `+f.SpaceName())
}

// SetId create mutations, should not duplicate
func (f *FilesMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != f.Id {
		f.mutations = append(f.mutations, A.X{`=`, 0, val})
		f.logs = append(f.logs, A.X{`id`, f.Id, val})
		f.Id = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (f *FilesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != f.CreatedAt {
		f.mutations = append(f.mutations, A.X{`=`, 1, val})
		f.logs = append(f.logs, A.X{`createdAt`, f.CreatedAt, val})
		f.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (f *FilesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.CreatedBy {
		f.mutations = append(f.mutations, A.X{`=`, 2, val})
		f.logs = append(f.logs, A.X{`createdBy`, f.CreatedBy, val})
		f.CreatedBy = val
		return true
	}
	return false
}

// SetMime create mutations, should not duplicate
func (f *FilesMutator) SetMime(val string) bool { //nolint:dupl false positive
	if val != f.Mime {
		f.mutations = append(f.mutations, A.X{`=`, 3, val})
		f.logs = append(f.logs, A.X{`mime`, f.Mime, val})
		f.Mime = val
		return true
	}
	return false
}

// SetPurpose create mutations, should not duplicate
func (f *FilesMutator) SetPurpose(val string) bool { //nolint:dupl false positive
	if val != f.Purpose {
		f.mutations = append(f.mutations, A.X{`=`, 4, val})
		f.logs = append(f.logs, A.X{`purpose`, f.Purpose, val})
		f.Purpose = val
		return true
	}
	return false
}

// SetRefId create mutations, should not duplicate
func (f *FilesMutator) SetRefId(val uint64) bool { //nolint:dupl false positive
	if val != f.RefId {
		f.mutations = append(f.mutations, A.X{`=`, 5, val})
		f.logs = append(f.logs, A.X{`refId`, f.RefId, val})
		f.RefId = val
		return true
	}
	return false
}

// SetAccessCount create mutations, should not duplicate
func (f *FilesMutator) SetAccessCount(val uint64) bool { //nolint:dupl false positive
	if val != f.AccessCount {
		f.mutations = append(f.mutations, A.X{`=`, 6, val})
		f.logs = append(f.logs, A.X{`accessCount`, f.AccessCount, val})
		f.AccessCount = val
		return true
	}
	return false
}

// SetLastAccessAt create mutations, should not duplicate
func (f *FilesMutator) SetLastAccessAt(val int64) bool { //nolint:dupl false positive
	if val != f.LastAccessAt {
		f.mutations = append(f.mutations, A.X{`=`, 7, val})
		f.logs = append(f.logs, A.X{`lastAccessAt`, f.LastAccessAt, val})
		f.LastAccessAt = val
		return true
	}
	return false
}

// SetOriginalPath create mutations, should not duplicate
func (f *FilesMutator) SetOriginalPath(val string) bool { //nolint:dupl false positive
	if val != f.OriginalPath {
		f.mutations = append(f.mutations, A.X{`=`, 8, val})
		f.logs = append(f.logs, A.X{`originalPath`, f.OriginalPath, val})
		f.OriginalPath = val
		return true
	}
	return false
}

// SetOriginalSize create mutations, should not duplicate
func (f *FilesMutator) SetOriginalSize(val uint64) bool { //nolint:dupl false positive
	if val != f.OriginalSize {
		f.mutations = append(f.mutations, A.X{`=`, 9, val})
		f.logs = append(f.logs, A.X{`originalSize`, f.OriginalSize, val})
		f.OriginalSize = val
		return true
	}
	return false
}

// SetResizedPath create mutations, should not duplicate
func (f *FilesMutator) SetResizedPath(val string) bool { //nolint:dupl false positive
	if val != f.ResizedPath {
		f.mutations = append(f.mutations, A.X{`=`, 10, val})
		f.logs = append(f.logs, A.X{`resizedPath`, f.ResizedPath, val})
		f.ResizedPath = val
		return true
	}
	return false
}

// SetResizedSize create mutations, should not duplicate
func (f *FilesMutator) SetResizedSize(val uint64) bool { //nolint:dupl false positive
	if val != f.ResizedSize {
		f.mutations = append(f.mutations, A.X{`=`, 11, val})
		f.logs = append(f.logs, A.X{`resizedSize`, f.ResizedSize, val})
		f.ResizedSize = val
		return true
	}
	return false
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
