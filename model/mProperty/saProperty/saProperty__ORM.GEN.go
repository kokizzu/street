package saProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	"database/sql"
	"time"

	"street/model/mProperty"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	chBuffer "github.com/kokizzu/ch-timed-buffer"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type saProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type saProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type saProperty__ORM.GEN.go
// go:generate msgp -tests=false -file saProperty__ORM.GEN.go -o saProperty__MSG.GEN.go

var Preparators = map[Ch.TableName]chBuffer.Preparator{}
