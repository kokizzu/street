package mStorage

import (
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TableFiles Tt.TableName = `files`

	Id           = `id`
	CreatedAt    = `createdAt`
	CreatedBy    = `createdBy`
	Mime         = `mime`
	Purpose      = `purpose` // purpose of this upload, eg. property image, floor image
	RefId        = `refId`   // reference table id
	AccessCount  = `accessCount`
	LastAccessAt = `lastAccessAt`

	OriginalPath = `originalPath` // eg. original.png
	OriginalSize = `originalSize`
	ResizedPath  = `resizedPath` // eg. small.png
	ResizedSize  = `resizedSize`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableFiles: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{Mime, Tt.String},
			{Purpose, Tt.String},
			{RefId, Tt.Unsigned},
			{AccessCount, Tt.Unsigned},
			{LastAccessAt, Tt.Integer},

			{OriginalPath, Tt.String},
			{OriginalSize, Tt.Unsigned},
			{ResizedPath, Tt.String},
			{ResizedSize, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Unique1:         OriginalPath,
		Engine:          Tt.Memtx, // memtx for now until it's overcapacity
	},
}
