package saProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	`database/sql`
	`street/model/mProperty`
	`time`

	_ `github.com/ClickHouse/clickhouse-go/v2`
	chBuffer `github.com/kokizzu/ch-timed-buffer`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Ch`
	`github.com/kokizzu/gotro/L`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type saProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type saProperty__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type saProperty__ORM.GEN.go
// go:generate msgp -tests=false -file saProperty__ORM.GEN.go -o saProperty__MSG.GEN.go

var tablePropertyLogsDummy = TablePropertyLogs{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mProperty.TableTablePropertyLogs: func(tx *sql.Tx) *sql.Stmt {
		query := tablePropertyLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}
type TablePropertyLogs struct {
	Adapter *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	CreatedAt time.Time
	RequestId string
	Error     string
	ActorId   uint64
	IpAddr4   string
	IpAddr6   string
	UserAgent string
}

func NewTablePropertyLogs(adapter *Ch.Adapter) *TablePropertyLogs {
	return &TablePropertyLogs{Adapter: adapter}
}

func (t TablePropertyLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableTablePropertyLogs
}

func (t *TablePropertyLogs) SqlTableName() string { //nolint:dupl false positive
	return `"tablePropertyLogs"`
}

// insert, error if exists
func (t *TablePropertyLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + t.SqlTableName() + `(` + t.SqlAllFields() + `) VALUES (?,?,?,?,?,?,?)`
}

func (t *TablePropertyLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + t.SqlTableName()
}

func (t *TablePropertyLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, requestId
	, error
	, actorId
	, ipAddr4
	, ipAddr6
	, userAgent
	`
}

func (t *TablePropertyLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, requestId, error, actorId, ipAddr4, ipAddr6, userAgent`
}

func (t TablePropertyLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		t.CreatedAt, // 0 
		t.RequestId, // 1 
		t.Error, // 2 
		t.ActorId, // 3 
		t.IpAddr4, // 4 
		t.IpAddr6, // 5 
		t.UserAgent, // 6 
	}
}

func (t *TablePropertyLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (t *TablePropertyLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (t *TablePropertyLogs) IdxRequestId() int { //nolint:dupl false positive
	return 1
}

func (t *TablePropertyLogs) SqlRequestId() string { //nolint:dupl false positive
	return `requestId`
}

func (t *TablePropertyLogs) IdxError() int { //nolint:dupl false positive
	return 2
}

func (t *TablePropertyLogs) SqlError() string { //nolint:dupl false positive
	return `error`
}

func (t *TablePropertyLogs) IdxActorId() int { //nolint:dupl false positive
	return 3
}

func (t *TablePropertyLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (t *TablePropertyLogs) IdxIpAddr4() int { //nolint:dupl false positive
	return 4
}

func (t *TablePropertyLogs) SqlIpAddr4() string { //nolint:dupl false positive
	return `ipAddr4`
}

func (t *TablePropertyLogs) IdxIpAddr6() int { //nolint:dupl false positive
	return 5
}

func (t *TablePropertyLogs) SqlIpAddr6() string { //nolint:dupl false positive
	return `ipAddr6`
}

func (t *TablePropertyLogs) IdxUserAgent() int { //nolint:dupl false positive
	return 6
}

func (t *TablePropertyLogs) SqlUserAgent() string { //nolint:dupl false positive
	return `userAgent`
}

func (t *TablePropertyLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		t.CreatedAt, // 0
		t.RequestId, // 1
		t.Error,     // 2
		t.ActorId,   // 3
		t.IpAddr4,   // 4
		t.IpAddr6,   // 5
		t.UserAgent, // 6
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

