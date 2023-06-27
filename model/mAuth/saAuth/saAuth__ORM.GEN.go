package saAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	`database/sql`
	`street/model/mAuth`
	`time`

	_ `github.com/ClickHouse/clickhouse-go/v2`
	chBuffer `github.com/kokizzu/ch-timed-buffer`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Ch`
	`github.com/kokizzu/gotro/L`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saAuth__ORM.GEN.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type saAuth__ORM.GEN.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type saAuth__ORM.GEN.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type saAuth__ORM.GEN.go
// go:generate msgp -tests=false -file saAuth__ORM.GEN.go -o saAuth__MSG.GEN.go

var actionLogsDummy = ActionLogs{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mAuth.TableActionLogs: func(tx *sql.Tx) *sql.Stmt {
		query := actionLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}
type ActionLogs struct {
	Adapter *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	CreatedAt  time.Time
	RequestId  string
	ActorId    uint64
	Action     string
	StatusCode int16
	Traces     string
	Error      string
	IpAddr4    string
	IpAddr6    string
	UserAgent  string
	Lat        float64
	Long       float64
	Latency    float64
}

func NewActionLogs(adapter *Ch.Adapter) *ActionLogs {
	return &ActionLogs{Adapter: adapter}
}

func (a ActionLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mAuth.TableActionLogs
}

func (a *ActionLogs) SqlTableName() string { //nolint:dupl false positive
	return `"actionLogs"`
}

// insert, error if exists
func (a *ActionLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + a.SqlTableName() + `(` + a.SqlAllFields() + `) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`
}

func (a *ActionLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + a.SqlTableName()
}

func (a *ActionLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, requestId
	, actorId
	, action
	, statusCode
	, traces
	, error
	, ipAddr4
	, ipAddr6
	, userAgent
	, lat
	, long
	, latency
	`
}

func (a *ActionLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, requestId, actorId, action, statusCode, traces, error, ipAddr4, ipAddr6, userAgent, lat, long, latency`
}

func (a ActionLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		a.CreatedAt, // 0 
		a.RequestId, // 1 
		a.ActorId, // 2 
		a.Action, // 3 
		a.StatusCode, // 4 
		a.Traces, // 5 
		a.Error, // 6 
		a.IpAddr4, // 7 
		a.IpAddr6, // 8 
		a.UserAgent, // 9 
		a.Lat, // 10 
		a.Long, // 11 
		a.Latency, // 12 
	}
}

func (a *ActionLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (a *ActionLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (a *ActionLogs) IdxRequestId() int { //nolint:dupl false positive
	return 1
}

func (a *ActionLogs) SqlRequestId() string { //nolint:dupl false positive
	return `requestId`
}

func (a *ActionLogs) IdxActorId() int { //nolint:dupl false positive
	return 2
}

func (a *ActionLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (a *ActionLogs) IdxAction() int { //nolint:dupl false positive
	return 3
}

func (a *ActionLogs) SqlAction() string { //nolint:dupl false positive
	return `action`
}

func (a *ActionLogs) IdxStatusCode() int { //nolint:dupl false positive
	return 4
}

func (a *ActionLogs) SqlStatusCode() string { //nolint:dupl false positive
	return `statusCode`
}

func (a *ActionLogs) IdxTraces() int { //nolint:dupl false positive
	return 5
}

func (a *ActionLogs) SqlTraces() string { //nolint:dupl false positive
	return `traces`
}

func (a *ActionLogs) IdxError() int { //nolint:dupl false positive
	return 6
}

func (a *ActionLogs) SqlError() string { //nolint:dupl false positive
	return `error`
}

func (a *ActionLogs) IdxIpAddr4() int { //nolint:dupl false positive
	return 7
}

func (a *ActionLogs) SqlIpAddr4() string { //nolint:dupl false positive
	return `ipAddr4`
}

func (a *ActionLogs) IdxIpAddr6() int { //nolint:dupl false positive
	return 8
}

func (a *ActionLogs) SqlIpAddr6() string { //nolint:dupl false positive
	return `ipAddr6`
}

func (a *ActionLogs) IdxUserAgent() int { //nolint:dupl false positive
	return 9
}

func (a *ActionLogs) SqlUserAgent() string { //nolint:dupl false positive
	return `userAgent`
}

func (a *ActionLogs) IdxLat() int { //nolint:dupl false positive
	return 10
}

func (a *ActionLogs) SqlLat() string { //nolint:dupl false positive
	return `lat`
}

func (a *ActionLogs) IdxLong() int { //nolint:dupl false positive
	return 11
}

func (a *ActionLogs) SqlLong() string { //nolint:dupl false positive
	return `long`
}

func (a *ActionLogs) IdxLatency() int { //nolint:dupl false positive
	return 12
}

func (a *ActionLogs) SqlLatency() string { //nolint:dupl false positive
	return `latency`
}

func (a *ActionLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		a.CreatedAt,  // 0
		a.RequestId,  // 1
		a.ActorId,    // 2
		a.Action,     // 3
		a.StatusCode, // 4
		a.Traces,     // 5
		a.Error,      // 6
		a.IpAddr4,    // 7
		a.IpAddr6,    // 8
		a.UserAgent,  // 9
		a.Lat,        // 10
		a.Long,       // 11
		a.Latency,    // 12
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

