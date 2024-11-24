package wcStorage

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mStorage/rqStorage"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// DesignFilesMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcStorage__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcStorage__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcStorage__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcStorage__ORM.GEN.go
type DesignFilesMutator struct {
	rqStorage.DesignFiles
	mutations []A.X
	logs      []A.X
}

// NewDesignFilesMutator create new ORM writer/command object
func NewDesignFilesMutator(adapter *Tt.Adapter) (res *DesignFilesMutator) {
	res = &DesignFilesMutator{DesignFiles: rqStorage.DesignFiles{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (d *DesignFilesMutator) Logs() []A.X { //nolint:dupl false positive
	return d.logs
}

// HaveMutation check whether Set* methods ever called
func (d *DesignFilesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(d.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (d *DesignFilesMutator) ClearMutations() { //nolint:dupl false positive
	d.mutations = []A.X{}
	d.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (d *DesignFilesMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := d.Adapter.Update(d.SpaceName(), d.UniqueIndexId(), A.X{d.Id}, d.ToUpdateArray())
	return !L.IsError(err, `DesignFiles.DoOverwriteById failed: `+d.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (d *DesignFilesMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !d.HaveMutation() {
		return true
	}
	_, err := d.Adapter.Update(d.SpaceName(), d.UniqueIndexId(), A.X{d.Id}, d.mutations)
	return !L.IsError(err, `DesignFiles.DoUpdateById failed: `+d.SpaceName())
}

// DoDeletePermanentById permanent delete
func (d *DesignFilesMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := d.Adapter.Delete(d.SpaceName(), d.UniqueIndexId(), A.X{d.Id})
	return !L.IsError(err, `DesignFiles.DoDeletePermanentById failed: `+d.SpaceName())
}

// func (d *DesignFilesMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := d.ToArray()
//	_, err := d.Adapter.Upsert(d.SpaceName(), arr, A.X{
//		A.X{`=`, 0, d.Id},
//		A.X{`=`, 1, d.CountryPropId},
//		A.X{`=`, 2, d.FilePath},
//		A.X{`=`, 3, d.CreatedAt},
//		A.X{`=`, 4, d.CreatedBy},
//	})
//	return !L.IsError(err, `DesignFiles.DoUpsert failed: `+d.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByFilePath update all columns, error if not exists, not using mutations/Set*
func (d *DesignFilesMutator) DoOverwriteByFilePath() bool { //nolint:dupl false positive
	_, err := d.Adapter.Update(d.SpaceName(), d.UniqueIndexFilePath(), A.X{d.FilePath}, d.ToUpdateArray())
	return !L.IsError(err, `DesignFiles.DoOverwriteByFilePath failed: `+d.SpaceName())
}

// DoUpdateByFilePath update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (d *DesignFilesMutator) DoUpdateByFilePath() bool { //nolint:dupl false positive
	if !d.HaveMutation() {
		return true
	}
	_, err := d.Adapter.Update(d.SpaceName(), d.UniqueIndexFilePath(), A.X{d.FilePath}, d.mutations)
	return !L.IsError(err, `DesignFiles.DoUpdateByFilePath failed: `+d.SpaceName())
}

// DoDeletePermanentByFilePath permanent delete
func (d *DesignFilesMutator) DoDeletePermanentByFilePath() bool { //nolint:dupl false positive
	_, err := d.Adapter.Delete(d.SpaceName(), d.UniqueIndexFilePath(), A.X{d.FilePath})
	return !L.IsError(err, `DesignFiles.DoDeletePermanentByFilePath failed: `+d.SpaceName())
}

// DoOverwriteByCountryPropId update all columns, error if not exists, not using mutations/Set*
func (d *DesignFilesMutator) DoOverwriteByCountryPropId() bool { //nolint:dupl false positive
	_, err := d.Adapter.Update(d.SpaceName(), d.UniqueIndexCountryPropId(), A.X{d.CountryPropId}, d.ToUpdateArray())
	return !L.IsError(err, `DesignFiles.DoOverwriteByCountryPropId failed: `+d.SpaceName())
}

// DoUpdateByCountryPropId update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (d *DesignFilesMutator) DoUpdateByCountryPropId() bool { //nolint:dupl false positive
	if !d.HaveMutation() {
		return true
	}
	_, err := d.Adapter.Update(d.SpaceName(), d.UniqueIndexCountryPropId(), A.X{d.CountryPropId}, d.mutations)
	return !L.IsError(err, `DesignFiles.DoUpdateByCountryPropId failed: `+d.SpaceName())
}

// DoDeletePermanentByCountryPropId permanent delete
func (d *DesignFilesMutator) DoDeletePermanentByCountryPropId() bool { //nolint:dupl false positive
	_, err := d.Adapter.Delete(d.SpaceName(), d.UniqueIndexCountryPropId(), A.X{d.CountryPropId})
	return !L.IsError(err, `DesignFiles.DoDeletePermanentByCountryPropId failed: `+d.SpaceName())
}

// DoInsert insert, error if already exists
func (d *DesignFilesMutator) DoInsert() bool { //nolint:dupl false positive
	arr := d.ToArray()
	row, err := d.Adapter.Insert(d.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			d.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `DesignFiles.DoInsert failed: `+d.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (d *DesignFilesMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := d.ToArray()
	row, err := d.Adapter.Replace(d.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			d.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `DesignFiles.DoUpsert failed: `+d.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (d *DesignFilesMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != d.Id {
		d.mutations = append(d.mutations, A.X{`=`, 0, val})
		d.logs = append(d.logs, A.X{`id`, d.Id, val})
		d.Id = val
		return true
	}
	return false
}

// SetCountryPropId create mutations, should not duplicate
func (d *DesignFilesMutator) SetCountryPropId(val string) bool { //nolint:dupl false positive
	if val != d.CountryPropId {
		d.mutations = append(d.mutations, A.X{`=`, 1, val})
		d.logs = append(d.logs, A.X{`countryPropId`, d.CountryPropId, val})
		d.CountryPropId = val
		return true
	}
	return false
}

// SetFilePath create mutations, should not duplicate
func (d *DesignFilesMutator) SetFilePath(val string) bool { //nolint:dupl false positive
	if val != d.FilePath {
		d.mutations = append(d.mutations, A.X{`=`, 2, val})
		d.logs = append(d.logs, A.X{`filePath`, d.FilePath, val})
		d.FilePath = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (d *DesignFilesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != d.CreatedAt {
		d.mutations = append(d.mutations, A.X{`=`, 3, val})
		d.logs = append(d.logs, A.X{`createdAt`, d.CreatedAt, val})
		d.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (d *DesignFilesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != d.CreatedBy {
		d.mutations = append(d.mutations, A.X{`=`, 4, val})
		d.logs = append(d.logs, A.X{`createdBy`, d.CreatedBy, val})
		d.CreatedBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (d *DesignFilesMutator) SetAll(from rqStorage.DesignFiles, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		d.Id = from.Id
		changed = true
	}
	if !excludeMap[`countryPropId`] && (forceMap[`countryPropId`] || from.CountryPropId != ``) {
		d.CountryPropId = S.Trim(from.CountryPropId)
		changed = true
	}
	if !excludeMap[`filePath`] && (forceMap[`filePath`] || from.FilePath != ``) {
		d.FilePath = S.Trim(from.FilePath)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		d.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		d.CreatedBy = from.CreatedBy
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// FilesMutator DAO writer/command struct
type FilesMutator struct {
	rqStorage.Files
	mutations []A.X
	logs      []A.X
}

// NewFilesMutator create new ORM writer/command object
func NewFilesMutator(adapter *Tt.Adapter) (res *FilesMutator) {
	res = &FilesMutator{Files: rqStorage.Files{Adapter: adapter}}
	return
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
//	arr := f.ToArray()
//	_, err := f.Adapter.Upsert(f.SpaceName(), arr, A.X{
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
//	return !L.IsError(err, `Files.DoUpsert failed: `+f.SpaceName()+ `\n%#v`, arr)
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
	arr := f.ToArray()
	row, err := f.Adapter.Insert(f.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Files.DoInsert failed: `+f.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (f *FilesMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := f.ToArray()
	row, err := f.Adapter.Replace(f.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Files.DoUpsert failed: `+f.SpaceName()+`\n%#v`, arr)
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

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (f *FilesMutator) SetAll(from rqStorage.Files, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		f.Id = from.Id
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		f.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		f.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`mime`] && (forceMap[`mime`] || from.Mime != ``) {
		f.Mime = S.Trim(from.Mime)
		changed = true
	}
	if !excludeMap[`purpose`] && (forceMap[`purpose`] || from.Purpose != ``) {
		f.Purpose = S.Trim(from.Purpose)
		changed = true
	}
	if !excludeMap[`refId`] && (forceMap[`refId`] || from.RefId != 0) {
		f.RefId = from.RefId
		changed = true
	}
	if !excludeMap[`accessCount`] && (forceMap[`accessCount`] || from.AccessCount != 0) {
		f.AccessCount = from.AccessCount
		changed = true
	}
	if !excludeMap[`lastAccessAt`] && (forceMap[`lastAccessAt`] || from.LastAccessAt != 0) {
		f.LastAccessAt = from.LastAccessAt
		changed = true
	}
	if !excludeMap[`originalPath`] && (forceMap[`originalPath`] || from.OriginalPath != ``) {
		f.OriginalPath = S.Trim(from.OriginalPath)
		changed = true
	}
	if !excludeMap[`originalSize`] && (forceMap[`originalSize`] || from.OriginalSize != 0) {
		f.OriginalSize = from.OriginalSize
		changed = true
	}
	if !excludeMap[`resizedPath`] && (forceMap[`resizedPath`] || from.ResizedPath != ``) {
		f.ResizedPath = S.Trim(from.ResizedPath)
		changed = true
	}
	if !excludeMap[`resizedSize`] && (forceMap[`resizedSize`] || from.ResizedSize != 0) {
		f.ResizedSize = from.ResizedSize
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
