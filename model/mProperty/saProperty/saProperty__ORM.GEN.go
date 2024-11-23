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

var scannedAreasDummy = ScannedAreas{}
var scannedPropertiesDummy = ScannedProperties{}
var viewedRoomsDummy = ViewedRooms{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mProperty.TableScannedAreas: func(tx *sql.Tx) *sql.Stmt {
		query := scannedAreasDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TableScannedProperties: func(tx *sql.Tx) *sql.Stmt {
		query := scannedPropertiesDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TableViewedRooms: func(tx *sql.Tx) *sql.Stmt {
		query := viewedRoomsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}

type ScannedAreas struct {
	Adapter   *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	ActorId   uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	CreatedAt time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	Latitude  float64     `json:"latitude" form:"latitude" query:"latitude" long:"latitude" msg:"latitude"`
	Longitude float64     `json:"longitude" form:"longitude" query:"longitude" long:"longitude" msg:"longitude"`
	City      string      `json:"city" form:"city" query:"city" long:"city" msg:"city"`
	State     string      `json:"state" form:"state" query:"state" long:"state" msg:"state"`
}

func NewScannedAreas(adapter *Ch.Adapter) *ScannedAreas {
	return &ScannedAreas{Adapter: adapter}
}

// ScannedAreasFieldTypeMap returns key value of field name and key
var ScannedAreasFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`actorId`:   Ch.UInt64,
	`createdAt`: Ch.DateTime,
	`latitude`:  Ch.Float64,
	`longitude`: Ch.Float64,
	`city`:      Ch.String,
	`state`:     Ch.String,
}

func (s *ScannedAreas) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableScannedAreas
}

func (s *ScannedAreas) SqlTableName() string { //nolint:dupl false positive
	return `"scannedAreas"`
}

func (s *ScannedAreas) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&s.ActorId,
		&s.CreatedAt,
		&s.Latitude,
		&s.Longitude,
		&s.City,
		&s.State,
	)
}

func (s *ScannedAreas) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []ScannedAreas, err error) { //nolint:dupl false positive
	res = make([]ScannedAreas, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row ScannedAreas
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (s *ScannedAreas) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + s.SqlTableName() + `(` + s.SqlAllFields() + `) VALUES (?,?,?,?,?,?)`
}

func (s *ScannedAreas) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + s.SqlTableName()
}

func (s *ScannedAreas) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` actorId
	, createdAt
	, latitude
	, longitude
	, city
	, state
	`
}

func (s *ScannedAreas) SqlAllFields() string { //nolint:dupl false positive
	return `actorId, createdAt, latitude, longitude, city, state`
}

func (s ScannedAreas) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		s.ActorId,   // 0
		s.CreatedAt, // 1
		s.Latitude,  // 2
		s.Longitude, // 3
		s.City,      // 4
		s.State,     // 5
	}
}

func (s *ScannedAreas) IdxActorId() int { //nolint:dupl false positive
	return 0
}

func (s *ScannedAreas) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (s *ScannedAreas) IdxCreatedAt() int { //nolint:dupl false positive
	return 1
}

func (s *ScannedAreas) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (s *ScannedAreas) IdxLatitude() int { //nolint:dupl false positive
	return 2
}

func (s *ScannedAreas) SqlLatitude() string { //nolint:dupl false positive
	return `latitude`
}

func (s *ScannedAreas) IdxLongitude() int { //nolint:dupl false positive
	return 3
}

func (s *ScannedAreas) SqlLongitude() string { //nolint:dupl false positive
	return `longitude`
}

func (s *ScannedAreas) IdxCity() int { //nolint:dupl false positive
	return 4
}

func (s *ScannedAreas) SqlCity() string { //nolint:dupl false positive
	return `city`
}

func (s *ScannedAreas) IdxState() int { //nolint:dupl false positive
	return 5
}

func (s *ScannedAreas) SqlState() string { //nolint:dupl false positive
	return `state`
}

func (s *ScannedAreas) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.ActorId,   // 0
		s.CreatedAt, // 1
		s.Latitude,  // 2
		s.Longitude, // 3
		s.City,      // 4
		s.State,     // 5
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type ScannedProperties struct {
	Adapter     *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	ActorId     uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	CreatedAt   time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CountryCode string      `json:"countryCode" form:"countryCode" query:"countryCode" long:"countryCode" msg:"countryCode"`
	PropertyId  uint64      `json:"propertyId,string" form:"propertyId" query:"propertyId" long:"propertyId" msg:"propertyId"`
}

func NewScannedProperties(adapter *Ch.Adapter) *ScannedProperties {
	return &ScannedProperties{Adapter: adapter}
}

// ScannedPropertiesFieldTypeMap returns key value of field name and key
var ScannedPropertiesFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`actorId`:     Ch.UInt64,
	`createdAt`:   Ch.DateTime,
	`countryCode`: Ch.String,
	`propertyId`:  Ch.UInt64,
}

func (s *ScannedProperties) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableScannedProperties
}

func (s *ScannedProperties) SqlTableName() string { //nolint:dupl false positive
	return `"scannedProperties"`
}

func (s *ScannedProperties) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&s.ActorId,
		&s.CreatedAt,
		&s.CountryCode,
		&s.PropertyId,
	)
}

func (s *ScannedProperties) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []ScannedProperties, err error) { //nolint:dupl false positive
	res = make([]ScannedProperties, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row ScannedProperties
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (s *ScannedProperties) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + s.SqlTableName() + `(` + s.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (s *ScannedProperties) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + s.SqlTableName()
}

func (s *ScannedProperties) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` actorId
	, createdAt
	, countryCode
	, propertyId
	`
}

func (s *ScannedProperties) SqlAllFields() string { //nolint:dupl false positive
	return `actorId, createdAt, countryCode, propertyId`
}

func (s ScannedProperties) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		s.ActorId,     // 0
		s.CreatedAt,   // 1
		s.CountryCode, // 2
		s.PropertyId,  // 3
	}
}

func (s *ScannedProperties) IdxActorId() int { //nolint:dupl false positive
	return 0
}

func (s *ScannedProperties) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (s *ScannedProperties) IdxCreatedAt() int { //nolint:dupl false positive
	return 1
}

func (s *ScannedProperties) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (s *ScannedProperties) IdxCountryCode() int { //nolint:dupl false positive
	return 2
}

func (s *ScannedProperties) SqlCountryCode() string { //nolint:dupl false positive
	return `countryCode`
}

func (s *ScannedProperties) IdxPropertyId() int { //nolint:dupl false positive
	return 3
}

func (s *ScannedProperties) SqlPropertyId() string { //nolint:dupl false positive
	return `propertyId`
}

func (s *ScannedProperties) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.ActorId,     // 0
		s.CreatedAt,   // 1
		s.CountryCode, // 2
		s.PropertyId,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type ViewedRooms struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	PropertyId uint64      `json:"propertyId,string" form:"propertyId" query:"propertyId" long:"propertyId" msg:"propertyId"`
	RoomLabel  string      `json:"roomLabel" form:"roomLabel" query:"roomLabel" long:"roomLabel" msg:"roomLabel"`
	Country    string      `json:"country" form:"country" query:"country" long:"country" msg:"country"`
}

func NewViewedRooms(adapter *Ch.Adapter) *ViewedRooms {
	return &ViewedRooms{Adapter: adapter}
}

// ViewedRoomsFieldTypeMap returns key value of field name and key
var ViewedRoomsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`actorId`:    Ch.UInt64,
	`createdAt`:  Ch.DateTime,
	`propertyId`: Ch.UInt64,
	`roomLabel`:  Ch.String,
	`country`:    Ch.String,
}

func (v *ViewedRooms) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableViewedRooms
}

func (v *ViewedRooms) SqlTableName() string { //nolint:dupl false positive
	return `"viewedRooms"`
}

func (v *ViewedRooms) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&v.ActorId,
		&v.CreatedAt,
		&v.PropertyId,
		&v.RoomLabel,
		&v.Country,
	)
}

func (v *ViewedRooms) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []ViewedRooms, err error) { //nolint:dupl false positive
	res = make([]ViewedRooms, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row ViewedRooms
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (v *ViewedRooms) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + v.SqlTableName() + `(` + v.SqlAllFields() + `) VALUES (?,?,?,?,?)`
}

func (v *ViewedRooms) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + v.SqlTableName()
}

func (v *ViewedRooms) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` actorId
	, createdAt
	, propertyId
	, roomLabel
	, country
	`
}

func (v *ViewedRooms) SqlAllFields() string { //nolint:dupl false positive
	return `actorId, createdAt, propertyId, roomLabel, country`
}

func (v ViewedRooms) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		v.ActorId,    // 0
		v.CreatedAt,  // 1
		v.PropertyId, // 2
		v.RoomLabel,  // 3
		v.Country,    // 4
	}
}

func (v *ViewedRooms) IdxActorId() int { //nolint:dupl false positive
	return 0
}

func (v *ViewedRooms) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (v *ViewedRooms) IdxCreatedAt() int { //nolint:dupl false positive
	return 1
}

func (v *ViewedRooms) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (v *ViewedRooms) IdxPropertyId() int { //nolint:dupl false positive
	return 2
}

func (v *ViewedRooms) SqlPropertyId() string { //nolint:dupl false positive
	return `propertyId`
}

func (v *ViewedRooms) IdxRoomLabel() int { //nolint:dupl false positive
	return 3
}

func (v *ViewedRooms) SqlRoomLabel() string { //nolint:dupl false positive
	return `roomLabel`
}

func (v *ViewedRooms) IdxCountry() int { //nolint:dupl false positive
	return 4
}

func (v *ViewedRooms) SqlCountry() string { //nolint:dupl false positive
	return `country`
}

func (v *ViewedRooms) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		v.ActorId,    // 0
		v.CreatedAt,  // 1
		v.PropertyId, // 2
		v.RoomLabel,  // 3
		v.Country,    // 4
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go
