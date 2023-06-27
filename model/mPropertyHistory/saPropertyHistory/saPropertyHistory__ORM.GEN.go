package saPropertyHistory

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	"database/sql"
	"street/model/mPropertyHistory"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	chBuffer "github.com/kokizzu/ch-timed-buffer"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saPropertyHistory__ORM.GEN.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type saPropertyHistory__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type saPropertyHistory__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type saPropertyHistory__ORM.GEN.go
// go:generate msgp -tests=false -file saPropertyHistory__ORM.GEN.go -o saPropertyHistory__MSG.GEN.go

var tablePropertyHistoryLogsDummy = TablePropertyHistoryLogs{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mPropertyHistory.TableTablePropertyHistoryLogs: func(tx *sql.Tx) *sql.Stmt {
		query := tablePropertyHistoryLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}

type TablePropertyHistoryLogs struct {
	Adapter   *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	RequestId string      `json:"requestId,string" form:"requestId" query:"requestId" long:"requestId" msg:"requestId"`
	Error     string      `json:"error" form:"error" query:"error" long:"error" msg:"error"`
	ActorId   uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	IpAddr4   string      `json:"ipAddr4" form:"ipAddr4" query:"ipAddr4" long:"ipAddr4" msg:"ipAddr4"`
	IpAddr6   string      `json:"ipAddr6" form:"ipAddr6" query:"ipAddr6" long:"ipAddr6" msg:"ipAddr6"`
	UserAgent string      `json:"userAgent" form:"userAgent" query:"userAgent" long:"userAgent" msg:"userAgent"`
}

func NewTablePropertyHistoryLogs(adapter *Ch.Adapter) *TablePropertyHistoryLogs {
	return &TablePropertyHistoryLogs{Adapter: adapter}
}

func (t TablePropertyHistoryLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mPropertyHistory.TableTablePropertyHistoryLogs
}

func (t *TablePropertyHistoryLogs) SqlTableName() string { //nolint:dupl false positive
	return `"tablePropertyHistoryLogs"`
}

// insert, error if exists
func (t *TablePropertyHistoryLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + t.SqlTableName() + `(` + t.SqlAllFields() + `) VALUES (?,?,?,?,?,?,?)`
}

func (t *TablePropertyHistoryLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + t.SqlTableName()
}

func (t *TablePropertyHistoryLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, requestId
	, error
	, actorId
	, ipAddr4
	, ipAddr6
	, userAgent
	`
}

func (t *TablePropertyHistoryLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, requestId, error, actorId, ipAddr4, ipAddr6, userAgent`
}

func (t TablePropertyHistoryLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		t.CreatedAt, // 0
		t.RequestId, // 1
		t.Error,     // 2
		t.ActorId,   // 3
		t.IpAddr4,   // 4
		t.IpAddr6,   // 5
		t.UserAgent, // 6
	}
}

func (t *TablePropertyHistoryLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (t *TablePropertyHistoryLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (t *TablePropertyHistoryLogs) IdxRequestId() int { //nolint:dupl false positive
	return 1
}

func (t *TablePropertyHistoryLogs) SqlRequestId() string { //nolint:dupl false positive
	return `requestId`
}

func (t *TablePropertyHistoryLogs) IdxError() int { //nolint:dupl false positive
	return 2
}

func (t *TablePropertyHistoryLogs) SqlError() string { //nolint:dupl false positive
	return `error`
}

func (t *TablePropertyHistoryLogs) IdxActorId() int { //nolint:dupl false positive
	return 3
}

func (t *TablePropertyHistoryLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (t *TablePropertyHistoryLogs) IdxIpAddr4() int { //nolint:dupl false positive
	return 4
}

func (t *TablePropertyHistoryLogs) SqlIpAddr4() string { //nolint:dupl false positive
	return `ipAddr4`
}

func (t *TablePropertyHistoryLogs) IdxIpAddr6() int { //nolint:dupl false positive
	return 5
}

func (t *TablePropertyHistoryLogs) SqlIpAddr6() string { //nolint:dupl false positive
	return `ipAddr6`
}

func (t *TablePropertyHistoryLogs) IdxUserAgent() int { //nolint:dupl false positive
	return 6
}

func (t *TablePropertyHistoryLogs) SqlUserAgent() string { //nolint:dupl false positive
	return `userAgent`
}

func (t *TablePropertyHistoryLogs) ToArray() A.X { //nolint:dupl false positive
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
