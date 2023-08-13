package rqStorage

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mStorage"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Files DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqStorage__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqStorage__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqStorage__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqStorage__ORM.GEN.go
type Files struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id           uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	CreatedAt    int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy    uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	Mime         string      `json:"mime" form:"mime" query:"mime" long:"mime" msg:"mime"`
	Purpose      string      `json:"purpose" form:"purpose" query:"purpose" long:"purpose" msg:"purpose"`
	RefId        uint64      `json:"refId,string" form:"refId" query:"refId" long:"refId" msg:"refId"`
	AccessCount  uint64      `json:"accessCount" form:"accessCount" query:"accessCount" long:"accessCount" msg:"accessCount"`
	LastAccessAt int64       `json:"lastAccessAt" form:"lastAccessAt" query:"lastAccessAt" long:"lastAccessAt" msg:"lastAccessAt"`
	OriginalPath string      `json:"originalPath" form:"originalPath" query:"originalPath" long:"originalPath" msg:"originalPath"`
	OriginalSize uint64      `json:"originalSize" form:"originalSize" query:"originalSize" long:"originalSize" msg:"originalSize"`
	ResizedPath  string      `json:"resizedPath" form:"resizedPath" query:"resizedPath" long:"resizedPath" msg:"resizedPath"`
	ResizedSize  uint64      `json:"resizedSize" form:"resizedSize" query:"resizedSize" long:"resizedSize" msg:"resizedSize"`
}

// NewFiles create new ORM reader/query object
func NewFiles(adapter *Tt.Adapter) *Files {
	return &Files{Adapter: adapter}
}

// SpaceName returns full package and table name
func (f *Files) SpaceName() string { //nolint:dupl false positive
	return string(mStorage.TableFiles) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (f *Files) SqlTableName() string { //nolint:dupl false positive
	return `"files"`
}

func (f *Files) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (f *Files) FindById() bool { //nolint:dupl false positive
	res, err := f.Adapter.Select(f.SpaceName(), f.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{f.Id})
	if L.IsError(err, `Files.FindById failed: `+f.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		f.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexOriginalPath return unique index name
func (f *Files) UniqueIndexOriginalPath() string { //nolint:dupl false positive
	return `originalPath`
}

// FindByOriginalPath Find one by OriginalPath
func (f *Files) FindByOriginalPath() bool { //nolint:dupl false positive
	res, err := f.Adapter.Select(f.SpaceName(), f.UniqueIndexOriginalPath(), 0, 1, tarantool.IterEq, A.X{f.OriginalPath})
	if L.IsError(err, `Files.FindByOriginalPath failed: `+f.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		f.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (f *Files) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "createdAt"
	, "createdBy"
	, "mime"
	, "purpose"
	, "refId"
	, "accessCount"
	, "lastAccessAt"
	, "originalPath"
	, "originalSize"
	, "resizedPath"
	, "resizedSize"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (f *Files) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "createdAt"
	, "createdBy"
	, "mime"
	, "purpose"
	, "refId"
	, "accessCount"
	, "lastAccessAt"
	, "originalPath"
	, "originalSize"
	, "resizedPath"
	, "resizedSize"
	`
}

// ToUpdateArray generate slice of update command
func (f *Files) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, f.Id},
		A.X{`=`, 1, f.CreatedAt},
		A.X{`=`, 2, f.CreatedBy},
		A.X{`=`, 3, f.Mime},
		A.X{`=`, 4, f.Purpose},
		A.X{`=`, 5, f.RefId},
		A.X{`=`, 6, f.AccessCount},
		A.X{`=`, 7, f.LastAccessAt},
		A.X{`=`, 8, f.OriginalPath},
		A.X{`=`, 9, f.OriginalSize},
		A.X{`=`, 10, f.ResizedPath},
		A.X{`=`, 11, f.ResizedSize},
	}
}

// IdxId return name of the index
func (f *Files) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (f *Files) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxCreatedAt return name of the index
func (f *Files) IdxCreatedAt() int { //nolint:dupl false positive
	return 1
}

// SqlCreatedAt return name of the column being indexed
func (f *Files) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (f *Files) IdxCreatedBy() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedBy return name of the column being indexed
func (f *Files) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxMime return name of the index
func (f *Files) IdxMime() int { //nolint:dupl false positive
	return 3
}

// SqlMime return name of the column being indexed
func (f *Files) SqlMime() string { //nolint:dupl false positive
	return `"mime"`
}

// IdxPurpose return name of the index
func (f *Files) IdxPurpose() int { //nolint:dupl false positive
	return 4
}

// SqlPurpose return name of the column being indexed
func (f *Files) SqlPurpose() string { //nolint:dupl false positive
	return `"purpose"`
}

// IdxRefId return name of the index
func (f *Files) IdxRefId() int { //nolint:dupl false positive
	return 5
}

// SqlRefId return name of the column being indexed
func (f *Files) SqlRefId() string { //nolint:dupl false positive
	return `"refId"`
}

// IdxAccessCount return name of the index
func (f *Files) IdxAccessCount() int { //nolint:dupl false positive
	return 6
}

// SqlAccessCount return name of the column being indexed
func (f *Files) SqlAccessCount() string { //nolint:dupl false positive
	return `"accessCount"`
}

// IdxLastAccessAt return name of the index
func (f *Files) IdxLastAccessAt() int { //nolint:dupl false positive
	return 7
}

// SqlLastAccessAt return name of the column being indexed
func (f *Files) SqlLastAccessAt() string { //nolint:dupl false positive
	return `"lastAccessAt"`
}

// IdxOriginalPath return name of the index
func (f *Files) IdxOriginalPath() int { //nolint:dupl false positive
	return 8
}

// SqlOriginalPath return name of the column being indexed
func (f *Files) SqlOriginalPath() string { //nolint:dupl false positive
	return `"originalPath"`
}

// IdxOriginalSize return name of the index
func (f *Files) IdxOriginalSize() int { //nolint:dupl false positive
	return 9
}

// SqlOriginalSize return name of the column being indexed
func (f *Files) SqlOriginalSize() string { //nolint:dupl false positive
	return `"originalSize"`
}

// IdxResizedPath return name of the index
func (f *Files) IdxResizedPath() int { //nolint:dupl false positive
	return 10
}

// SqlResizedPath return name of the column being indexed
func (f *Files) SqlResizedPath() string { //nolint:dupl false positive
	return `"resizedPath"`
}

// IdxResizedSize return name of the index
func (f *Files) IdxResizedSize() int { //nolint:dupl false positive
	return 11
}

// SqlResizedSize return name of the column being indexed
func (f *Files) SqlResizedSize() string { //nolint:dupl false positive
	return `"resizedSize"`
}

// ToArray receiver fields to slice
func (f *Files) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if f.Id != 0 {
		id = f.Id
	}
	return A.X{
		id,
		f.CreatedAt,    // 1
		f.CreatedBy,    // 2
		f.Mime,         // 3
		f.Purpose,      // 4
		f.RefId,        // 5
		f.AccessCount,  // 6
		f.LastAccessAt, // 7
		f.OriginalPath, // 8
		f.OriginalSize, // 9
		f.ResizedPath,  // 10
		f.ResizedSize,  // 11
	}
}

// FromArray convert slice to receiver fields
func (f *Files) FromArray(a A.X) *Files { //nolint:dupl false positive
	f.Id = X.ToU(a[0])
	f.CreatedAt = X.ToI(a[1])
	f.CreatedBy = X.ToU(a[2])
	f.Mime = X.ToS(a[3])
	f.Purpose = X.ToS(a[4])
	f.RefId = X.ToU(a[5])
	f.AccessCount = X.ToU(a[6])
	f.LastAccessAt = X.ToI(a[7])
	f.OriginalPath = X.ToS(a[8])
	f.OriginalSize = X.ToU(a[9])
	f.ResizedPath = X.ToS(a[10])
	f.ResizedSize = X.ToU(a[11])
	return f
}

// FromUncensoredArray convert slice to receiver fields
func (f *Files) FromUncensoredArray(a A.X) *Files { //nolint:dupl false positive
	f.Id = X.ToU(a[0])
	f.CreatedAt = X.ToI(a[1])
	f.CreatedBy = X.ToU(a[2])
	f.Mime = X.ToS(a[3])
	f.Purpose = X.ToS(a[4])
	f.RefId = X.ToU(a[5])
	f.AccessCount = X.ToU(a[6])
	f.LastAccessAt = X.ToI(a[7])
	f.OriginalPath = X.ToS(a[8])
	f.OriginalSize = X.ToU(a[9])
	f.ResizedPath = X.ToS(a[10])
	f.ResizedSize = X.ToU(a[11])
	return f
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (f *Files) FindOffsetLimit(offset, limit uint32, idx string) []Files { //nolint:dupl false positive
	var rows []Files
	res, err := f.Adapter.Select(f.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Files.FindOffsetLimit failed: `+f.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Files{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (f *Files) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := f.Adapter.Select(f.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Files.FindOffsetLimit failed: `+f.SpaceName()) {
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
func (f *Files) Total() int64 { //nolint:dupl false positive
	rows := f.Adapter.CallBoxSpace(f.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// FilesFieldTypeMap returns key value of field name and key
var FilesFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:           Tt.Unsigned,
	`createdAt`:    Tt.Integer,
	`createdBy`:    Tt.Unsigned,
	`mime`:         Tt.String,
	`purpose`:      Tt.String,
	`refId`:        Tt.Unsigned,
	`accessCount`:  Tt.Unsigned,
	`lastAccessAt`: Tt.Integer,
	`originalPath`: Tt.String,
	`originalSize`: Tt.Unsigned,
	`resizedPath`:  Tt.String,
	`resizedSize`:  Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
