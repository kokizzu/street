package TtMigration

import (
	"errors"
	"strings"

	"github.com/alitto/pond"
	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

var DEBUG = false
var MIGRATION_THREAD = 8
var MIGRATION_QUEUE = 1000
var MIGRATION_REC_PER_BATCH = uint64(20_000)

var ErrMiscaonfiguration = errors.New(`misconfiguration`)

func Descr(args ...any) {
	if DEBUG {
		L.ParentDescribe(args...)
	}
}

type TableName string

// types
type DataType string

const (
	// https://www.tarantool.io/en/doc/latest/concepts/data_model/value_store/#data-types

	Unsigned DataType = `unsigned`
	String   DataType = `string`
	Double   DataType = `double`
	Integer  DataType = `integer`
	Boolean  DataType = `boolean`
	Array    DataType = `array`

	//Any      DataType = `any`
	//Number   DataType = `number` // use double instead
	//Decimal  DataType = `decimal` // unsupported
	//DateTime DataType = `datetime` // unsupported
	//Uuid     DataType = `string`  // `uuid` // https://github.com/tarantool/go-tarantool/v2/issues/90
	//Scalar   DataType = `scalar`
	//Map      DataType = `map`

	//ArrayFloat    DataType = `array`
	//ArrayUnsigned DataType = `array`
	//ArrayInteger  DataType = `array`
)

var TypeToConst = map[DataType]string{
	Unsigned: `Tt.Unsigned`,
	String:   `Tt.String`,
	Integer:  `Tt.Integer`,
	Double:   `Tt.Double`,
	Boolean:  `Tt.Boolean`,
	Array:    `Tt.Array`,
}

var TypeToGoType = map[DataType]string{
	//Uuid:     `string`,
	Unsigned: `uint64`,
	String:   `string`,
	Integer:  `int64`,
	Double:   `float64`,
	Boolean:  `bool`,
	Array:    `[]any`,
}
var TypeToGoEmptyValue = map[DataType]string{
	Unsigned: `0`,
	String:   "``",
	Integer:  `0`,
	Double:   `0`,
	Boolean:  `false`,
	Array:    `[]any{}`,
}
var TypeToGoNilValue = map[DataType]string{
	Unsigned: `0`,
	String:   "``",
	Integer:  `0`,
	Double:   `0`,
	Boolean:  `false`,
	Array:    `nil`,
}
var TypeToConvertFunc = map[DataType]string{
	Unsigned: `X.ToU`,
	String:   `X.ToS`,
	Integer:  `X.ToI`,
	Double:   `X.ToF`,
	Boolean:  `X.ToBool`,
	Array:    `X.ToArr`,
}

var Type2TarantoolDefault = map[DataType]string{
	Unsigned: `0`,
	String:   `''`,
	Integer:  `0`,
	Boolean:  `false`,
	Double:   `0`,
	Array:    `[]`,
}

// misc
const Engine = `engine`

type EngineType string

const (
	Vinyl EngineType = `vinyl`
	Memtx EngineType = `memtx`
)
const BoxSpacePrefix = `box.space.`
const IfNotExists = `if_not_exists`

const IdCol = `id`

type TableProp struct {
	Fields []Field

	// indexes
	Unique1 string
	Unique2 string
	Unique3 string
	Uniques []string // multicolumn unique
	Indexes []string
	Spatial string

	Engine          EngineType
	HiddenFields    []string
	AutoIncrementId bool // "id" column will be used to generate sequence, can only be created at beginning
	GenGraphqlType  bool

	// hook
	PreReformatMigrationHook func(*Adapter)
	PreUnique1MigrationHook  func(*Adapter)
	PreUnique2MigrationHook  func(*Adapter)
	PreUnique3MigrationHook  func(*Adapter)
	PreUniquesMigrationHook  func(*Adapter)
	PreSpatialMigrationHook  func(*Adapter)

	AutoCensorFields []string // fields to automatically censor
}

type Field struct { // https://godoc.org/gopkg.in/vmihailenco/msgpack.v2#pkg-examples
	Name string   `msgpack:"name"`
	Type DataType `msgpack:"type"`
}

func (f Field) KeyRenderer() func(string) string {
	switch f.Type {
	case Unsigned:
		return func(structProp string) string {
			return "tarantool.UintKey{I:uint(" + structProp + `)}`
		}
	case Integer:
		return func(structProp string) string {
			return "tarantool.IntKey{I:int(" + structProp + `)}`
		}
	case String:
		return func(structProp string) string {
			return "tarantool.StringKey{S:" + structProp + `}`
		}
	}
	return func(structProp string) string {
		return `DataTypeNotPossibleToBeAKey:` + string(f.Type) + `:` + structProp
	}
}

type NullableField struct {
	Name       string   `msgpack:"name"`
	Type       DataType `msgpack:"type"`
	IsNullable bool     `msgpack:"is_nullable"`
}

type Index struct {
	Parts       []string `msgpack:"parts"`
	IfNotExists bool     `msgpack:"if_not_exists"`
	Unique      bool     `msgpack:"unique"`
	Sequence    string   `msgpack:"sequence,omitempty"`
	Type        string   `msgpack:"type,omitempty"`
}

type MSX map[string]any

func (a *Adapter) UpsertTable(tableName TableName, prop *TableProp) bool {
	Descr(`---UpsertTable-Start-` + tableName + `----------------------------------------------------`)
	defer Descr(`---UpsertTable-End-` + tableName + `-------------------------------------------------------`)
	if prop.Engine == `` {
		prop.Engine = Vinyl
	}
	Descr(`CreateSpace ` + tableName)
	if !a.CreateSpace(string(tableName), prop.Engine) {
		return false
	}
	if prop.PreReformatMigrationHook != nil {
		Descr(`PreReformatMigrationHook`)
		prop.PreReformatMigrationHook(a)
	}
	Descr(`ReformatTable ` + tableName)
	if !a.ReformatTable(string(tableName), prop.Fields) {
		return false // failed to create table
	}
	// create one field unique index
	a.ExecBoxSpace(string(tableName)+`:format`, A.X{})
	if prop.AutoIncrementId {
		if len(prop.Fields) < 1 || prop.Fields[0].Name != IdCol || prop.Fields[0].Type != Unsigned {
			panic(`must create Unsigned id field on first field to use AutoIncrementId`)
		}

		seqName := string(tableName) + `_` + IdCol
		Descr(`box.schema.sequence.create ` + seqName)
		a.ExecTarantoolVerbose(`box.schema.sequence.create`, A.X{
			seqName,
		})
		Descr(string(tableName) + `:create_index ` + seqName)
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			IdCol, Index{
				Sequence:    seqName,
				Parts:       []string{IdCol},
				IfNotExists: true,
				Unique:      true,
			},
		})
	}
	if prop.PreUnique3MigrationHook != nil {
		Descr(`PreUnique3MigrationHook`)
		prop.PreUnique3MigrationHook(a)
	}
	// only create unique if not "id"
	if prop.Unique1 != `` && !(prop.AutoIncrementId && prop.Unique1 == IdCol) {
		Descr(string(tableName) + `:create_index ` + prop.Unique1)
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			prop.Unique1, Index{Parts: []string{prop.Unique1}, IfNotExists: true, Unique: true},
		})
		if prop.Unique2 != `` && prop.Unique1 == prop.Unique2 {
			panic(`Unique1 and Unique2 must be unique`)
		}
		if prop.Unique3 != `` && prop.Unique1 == prop.Unique3 {
			panic(`Unique1 and Unique3 must be unique`)
		}
	}
	if prop.PreUnique2MigrationHook != nil {
		prop.PreUnique2MigrationHook(a)
	}
	if prop.Unique2 != `` && !(prop.AutoIncrementId && prop.Unique2 == IdCol) {
		Descr(string(tableName) + `:create_index ` + prop.Unique2)
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			prop.Unique2, Index{Parts: []string{prop.Unique2}, IfNotExists: true, Unique: true},
		})
		if prop.Unique3 != `` && prop.Unique2 == prop.Unique3 {
			panic(`Unique2 and Unique3 must be unique`)
		}
	}
	if prop.PreUnique3MigrationHook != nil {
		Descr(`PreUnique3MigrationHook`)
		prop.PreUnique3MigrationHook(a)
	}
	if prop.Unique3 != `` && !(prop.AutoIncrementId && prop.Unique3 == IdCol) {
		Descr(string(tableName) + `:create_index ` + prop.Unique3)
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			prop.Unique3, Index{Parts: []string{prop.Unique3}, IfNotExists: true, Unique: true},
		})
	}
	if prop.PreUniquesMigrationHook != nil {
		Descr(`PreUniquesMigrationHook`)
		prop.PreUniquesMigrationHook(a)
	}
	// create multi-field unique index: [col1, col2] will named col1__col2
	if len(prop.Uniques) > 1 {
		Descr(string(tableName) + `:create_index ` + strings.Join(prop.Uniques, `__`))
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			strings.Join(prop.Uniques, `__`), Index{Parts: prop.Uniques, IfNotExists: true, Unique: true},
		})
	}
	if prop.PreSpatialMigrationHook != nil {
		Descr(`PreSpatialMigrationHook`)
		prop.PreSpatialMigrationHook(a)
	}
	// create spatial index (only works for memtx)
	if prop.Spatial != `` {
		Descr(string(tableName) + `:create_index ` + prop.Spatial)
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			prop.Spatial, Index{Parts: []string{prop.Spatial}, IfNotExists: true, Type: `RTREE`},
		})
	}
	// create other indexes
	for _, index := range prop.Indexes {
		Descr(string(tableName) + `:create_index ` + index)
		//a.ExecBoxSpace(tableName+`.index.`+index+`:drop`, AX{index}) // TODO: remove this when index fixed
		a.ExecBoxSpace(string(tableName)+`:create_index`, A.X{
			index, Index{Parts: []string{index}, IfNotExists: true},
		})
	}
	// need refresh index after migrate
	// https://github.com/tarantool/go-tarantool/v2/pull/259#pullrequestreview-1242058107
	a.Connection = a.Reconnect()
	return true
}

// ignore return value
func (a *Adapter) ExecTarantool(funcName string, params A.X) bool {
	return a.ExecTarantoolVerbose(funcName, params) == ``
}

func (a *Adapter) ExecTarantoolVerbose(funcName string, params A.X) string {
	Descr(funcName)
	Descr(params)
	res, err := a.Connection.Do(tarantool.NewCallRequest(funcName).Args(params)).Get()
	if err != nil {
		if len(params) > 0 {
			errStr := err.Error()
			if errStr == `Space '`+X.ToS(params[0])+`' already exists` ||
				errStr == `Sequence '`+X.ToS(params[0])+`' already exists` {
				L.IsError(err, `ExecTarantool failed: `+funcName)
				return err.Error()
			}
		}
		if len(params) == 0 {
			L.IsError(err, `ExecTarantool failed: `+funcName)
			return err.Error()
		}
	}
	Descr(res)
	return ``
}

// ignore return value
func (a *Adapter) ExecBoxSpace(funcName string, params A.X) bool {
	return a.ExecBoxSpaceVerbose(funcName, params) == ``
}

func (a *Adapter) ExecBoxSpaceVerbose(funcName string, params A.X) string {
	Descr(funcName)
	Descr(params)
	res, err := a.Connection.Do(tarantool.NewCallRequest(BoxSpacePrefix + funcName).Args(params)).Get()
	if L.IsError(err, `ExecBoxSpace failed: `+funcName) {
		L.Describe(params)
		return err.Error()
	}
	Descr(res)
	return ``
}

// CallBoxSpace ignore return value
func (a *Adapter) CallBoxSpace(funcName string, params A.X) (rows [][]any) {
	Descr(funcName)
	Descr(params)
	res, err := a.Connection.Do(tarantool.NewCallRequest(BoxSpacePrefix + funcName).Args(params)).Get()
	Descr(res)
	if L.IsError(err, `ExecBoxSpace failed: `+funcName) {
		L.Describe(params)
		return
	}
	rows = make([][]any, 0)
	for _, row := range res {
		if row, ok := row.([]any); ok {
			rows = append(rows, row)
		}
	}
	return
}

func (a *Adapter) DropTable(tableName string) bool {
	return a.ExecBoxSpace(tableName+`:drop`, A.X{})
}

func (a *Adapter) TruncateTable(tableName string) bool {
	return a.ExecBoxSpace(tableName+`:truncate`, A.X{})
}

func (a *Adapter) ReformatTable(tableName string, fields []Field) bool {
	// check old schema
	a.Connection = a.Reconnect() // need reconnect after creating space or a.Schema.Spaces will be empty
	schema, err := tarantool.GetSchema(a.Connection)
	if L.IsError(err, `GetSchema failed`) {
		return false
	}
	table := schema.Spaces[tableName]
	if len(table.Fields) == 0 { // new table, create anyway
		return a.ExecBoxSpace(tableName+`:format`, A.X{
			&fields,
		}) // failed to create table
	}
	// table already exists
	var newFields []NullableField
	nullFields := map[string]DataType{}
	haveIdField := false
	for idx, newField := range fields { // diff and create nullable field
		origField := table.FieldsById[uint32(idx)]
		if origField.Type == string(newField.Type) {
			newFields = append(newFields, NullableField{Name: newField.Name, Type: newField.Type})
		} else {
			newFields = append(newFields, NullableField{Name: newField.Name, Type: newField.Type, IsNullable: true})
			nullFields[newField.Name] = newField.Type
		}
		if newField.Name == IdCol {
			haveIdField = true
		}
	}
	res := a.ExecBoxSpaceVerbose(tableName+`:format`, A.X{
		&newFields,
	})
	if res != `` {
		L.Describe(res)
	}
	a.Connection = a.Reconnect()
	// update all column
	if len(nullFields) > 0 {
		updateCols := []string{}
		for col, typ := range nullFields {
			defaultvalue := Type2TarantoolDefault[typ]
			if defaultvalue == `` {
				panic(`please set Type2TarantoolDefault for: ` + typ)
			}
			updateCols = append(updateCols, dq(col)+` = `+defaultvalue)
		}
		updateQuery := `UPDATE ` + dq(tableName) + ` SET ` + A.StrJoin(updateCols, `, `)
		if !haveIdField {
			a.ExecSql(updateQuery) // update whole thing since there's no id field
		} else {
			var count uint64
			a.QuerySql(`SELECT COUNT(*) FROM `+dq(tableName), func(row []any) {
				count = X.ToU(row[0])
			})
			pool := pond.New(MIGRATION_THREAD, MIGRATION_QUEUE)
			for i := uint64(0); i <= count; i += MIGRATION_REC_PER_BATCH {
				i := i // for go <1.21
				pool.Submit(func() {
					if i+MIGRATION_REC_PER_BATCH >= count { // remaining
						a.ExecSql(updateQuery + ` WHERE "id" >= ` + X.ToS(i))
					} else {
						a.ExecSql(updateQuery + ` WHERE "id" >= ` + X.ToS(i) + ` AND "id" < ` + X.ToS(i+MIGRATION_REC_PER_BATCH))
					}
					Descr(`Overwrite nulls ` + tableName + ` ` + X.ToS(i) + ` to ` + X.ToS(i+MIGRATION_REC_PER_BATCH) + ` of ` + X.ToS(count))
				})
			}
			pool.StopAndWait()
			a.ExecSql(updateQuery + ` WHERE "id" >= ` + X.ToS(count)) // update once more just in case
			L.Print(`if next step is is with error 'Tuple field N (xxx) required by space format is missing', run this query: ` + updateQuery + ` WHERE "id" >= ` + X.ToS(count))
			// can be caused by insert happened during migration
		}
	}

	return a.ExecBoxSpace(tableName+`:format`, A.X{
		&fields,
	})
}

func (a *Adapter) CreateSpace(tableName string, engine EngineType) bool {
	err := a.ExecTarantoolVerbose(`box.schema.space.create`, A.X{
		tableName,
		map[string]any{
			Engine: string(engine),
			//IfNotExists: true,
		},
	})
	if err == `` || S.StartsWith(err, `unsupported Lua type 'function'`) {
		return true // ignore
	}
	return true
}

func (a *Adapter) MigrateTables(tables map[TableName]*TableProp) {
	for name, props := range tables {
		Descr(name, props)
		// validity check
		if props.Engine == Vinyl && props.Spatial != `` {
			L.PanicIf(ErrMiscaonfiguration, `spatial index is not supported in vinyl engine`)
		}
		a.UpsertTable(name, props)
	}
}

func CheckTarantoolTables(tConn *Adapter, tables map[TableName]*TableProp) {
	for tableName := range tables {
		tableName := string(tableName)
		res, err := tConn.Connection.Do(tarantool.NewCallRequest(`box.space.` + tableName + `:format`).Args([]any{})).Get()
		L.PanicIf(err, `please run TABLE migration on %v for tarantool`, tableName)
		if len(res) != 1 {
			L.Panic(`please run FIELDS migration on %v for tarantool`, tableName)
		}
		columnNameTypes := map[string]string{}
		if len(res) > 0 {
			if rows, ok := res[0].([]any); ok {
				for _, row := range rows {
					m, ok := row.(map[any]any)
					if !ok {
						L.Panic(`please run FIELD migration on %v for tarantool: %v`, tableName, row)
					}
					columnNameTypes[X.ToS(m[`name`])] = X.ToS(m[`type`])
				}
			}
		}
		for _, field := range tables[TableName(tableName)].Fields {
			existingType := columnNameTypes[field.Name]
			if existingType != string(field.Type) {
				L.Panic(`please run FIELD_TYPE %v migration on %v for tarantool: %v <> %v`, field.Name, tableName, existingType, field.Type)
			}
		}
	}
}
