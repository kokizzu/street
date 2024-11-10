package rqAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"street/model/mAuth"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Feedbacks DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqAuth__ORM.GEN.go
type Feedbacks struct {
	Adapter     *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id          uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	CreatedBy   uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	CreatedAt   int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	UpdatedBy   uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	UpdatedAt   int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	DeletedAt   int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	UserMessage string      `json:"userMessage" form:"userMessage" query:"userMessage" long:"userMessage" msg:"userMessage"`
	AdminReply  string      `json:"adminReply" form:"adminReply" query:"adminReply" long:"adminReply" msg:"adminReply"`
}

// NewFeedbacks create new ORM reader/query object
func NewFeedbacks(adapter *Tt.Adapter) *Feedbacks {
	return &Feedbacks{Adapter: adapter}
}

// SpaceName returns full package and table name
func (f *Feedbacks) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableFeedbacks) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (f *Feedbacks) SqlTableName() string { //nolint:dupl false positive
	return `"feedbacks"`
}

func (f *Feedbacks) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (f *Feedbacks) FindById() bool { //nolint:dupl false positive
	res, err := f.Adapter.Select(f.SpaceName(), f.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{f.Id})
	if L.IsError(err, `Feedbacks.FindById failed: `+f.SpaceName()) {
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
func (f *Feedbacks) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "createdBy"
	, "createdAt"
	, "updatedBy"
	, "updatedAt"
	, "deletedAt"
	, "userMessage"
	, "adminReply"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (f *Feedbacks) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "createdBy"
	, "createdAt"
	, "updatedBy"
	, "updatedAt"
	, "deletedAt"
	, "userMessage"
	, "adminReply"
	`
}

// ToUpdateArray generate slice of update command
func (f *Feedbacks) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, f.Id},
		A.X{`=`, 1, f.CreatedBy},
		A.X{`=`, 2, f.CreatedAt},
		A.X{`=`, 3, f.UpdatedBy},
		A.X{`=`, 4, f.UpdatedAt},
		A.X{`=`, 5, f.DeletedAt},
		A.X{`=`, 6, f.UserMessage},
		A.X{`=`, 7, f.AdminReply},
	}
}

// IdxId return name of the index
func (f *Feedbacks) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (f *Feedbacks) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxCreatedBy return name of the index
func (f *Feedbacks) IdxCreatedBy() int { //nolint:dupl false positive
	return 1
}

// SqlCreatedBy return name of the column being indexed
func (f *Feedbacks) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxCreatedAt return name of the index
func (f *Feedbacks) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (f *Feedbacks) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxUpdatedBy return name of the index
func (f *Feedbacks) IdxUpdatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlUpdatedBy return name of the column being indexed
func (f *Feedbacks) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxUpdatedAt return name of the index
func (f *Feedbacks) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (f *Feedbacks) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxDeletedAt return name of the index
func (f *Feedbacks) IdxDeletedAt() int { //nolint:dupl false positive
	return 5
}

// SqlDeletedAt return name of the column being indexed
func (f *Feedbacks) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxUserMessage return name of the index
func (f *Feedbacks) IdxUserMessage() int { //nolint:dupl false positive
	return 6
}

// SqlUserMessage return name of the column being indexed
func (f *Feedbacks) SqlUserMessage() string { //nolint:dupl false positive
	return `"userMessage"`
}

// IdxAdminReply return name of the index
func (f *Feedbacks) IdxAdminReply() int { //nolint:dupl false positive
	return 7
}

// SqlAdminReply return name of the column being indexed
func (f *Feedbacks) SqlAdminReply() string { //nolint:dupl false positive
	return `"adminReply"`
}

// ToArray receiver fields to slice
func (f *Feedbacks) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if f.Id != 0 {
		id = f.Id
	}
	return A.X{
		id,
		f.CreatedBy,   // 1
		f.CreatedAt,   // 2
		f.UpdatedBy,   // 3
		f.UpdatedAt,   // 4
		f.DeletedAt,   // 5
		f.UserMessage, // 6
		f.AdminReply,  // 7
	}
}

// FromArray convert slice to receiver fields
func (f *Feedbacks) FromArray(a A.X) *Feedbacks { //nolint:dupl false positive
	f.Id = X.ToU(a[0])
	f.CreatedBy = X.ToU(a[1])
	f.CreatedAt = X.ToI(a[2])
	f.UpdatedBy = X.ToU(a[3])
	f.UpdatedAt = X.ToI(a[4])
	f.DeletedAt = X.ToI(a[5])
	f.UserMessage = X.ToS(a[6])
	f.AdminReply = X.ToS(a[7])
	return f
}

// FromUncensoredArray convert slice to receiver fields
func (f *Feedbacks) FromUncensoredArray(a A.X) *Feedbacks { //nolint:dupl false positive
	f.Id = X.ToU(a[0])
	f.CreatedBy = X.ToU(a[1])
	f.CreatedAt = X.ToI(a[2])
	f.UpdatedBy = X.ToU(a[3])
	f.UpdatedAt = X.ToI(a[4])
	f.DeletedAt = X.ToI(a[5])
	f.UserMessage = X.ToS(a[6])
	f.AdminReply = X.ToS(a[7])
	return f
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (f *Feedbacks) FindOffsetLimit(offset, limit uint32, idx string) []Feedbacks { //nolint:dupl false positive
	var rows []Feedbacks
	res, err := f.Adapter.Select(f.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Feedbacks.FindOffsetLimit failed: `+f.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Feedbacks{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (f *Feedbacks) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := f.Adapter.Select(f.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Feedbacks.FindOffsetLimit failed: `+f.SpaceName()) {
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
func (f *Feedbacks) Total() int64 { //nolint:dupl false positive
	rows := f.Adapter.CallBoxSpace(f.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// FeedbacksFieldTypeMap returns key value of field name and key
var FeedbacksFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:          Tt.Unsigned,
	`createdBy`:   Tt.Unsigned,
	`createdAt`:   Tt.Integer,
	`updatedBy`:   Tt.Unsigned,
	`updatedAt`:   Tt.Integer,
	`deletedAt`:   Tt.Integer,
	`userMessage`: Tt.String,
	`adminReply`:  Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Sessions DAO reader/query struct
type Sessions struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	SessionToken string      `json:"sessionToken" form:"sessionToken" query:"sessionToken" long:"sessionToken" msg:"sessionToken"`
	UserId       uint64      `json:"userId,string" form:"userId" query:"userId" long:"userId" msg:"userId"`
	ExpiredAt    int64       `json:"expiredAt" form:"expiredAt" query:"expiredAt" long:"expiredAt" msg:"expiredAt"`
	Device       string      `json:"device" form:"device" query:"device" long:"device" msg:"device"`
	LoginAt      int64       `json:"loginAt" form:"loginAt" query:"loginAt" long:"loginAt" msg:"loginAt"`
	LoginIPs     string      `json:"loginIPs" form:"loginIPs" query:"loginIPs" long:"loginIPs" msg:"loginIPs"`
}

// NewSessions create new ORM reader/query object
func NewSessions(adapter *Tt.Adapter) *Sessions {
	return &Sessions{Adapter: adapter}
}

// SpaceName returns full package and table name
func (s *Sessions) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableSessions) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (s *Sessions) SqlTableName() string { //nolint:dupl false positive
	return `"sessions"`
}

// UniqueIndexSessionToken return unique index name
func (s *Sessions) UniqueIndexSessionToken() string { //nolint:dupl false positive
	return `sessionToken`
}

// FindBySessionToken Find one by SessionToken
func (s *Sessions) FindBySessionToken() bool { //nolint:dupl false positive
	res, err := s.Adapter.Select(s.SpaceName(), s.UniqueIndexSessionToken(), 0, 1, tarantool.IterEq, A.X{s.SessionToken})
	if L.IsError(err, `Sessions.FindBySessionToken failed: `+s.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		s.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (s *Sessions) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "sessionToken"
	, "userId"
	, "expiredAt"
	, "device"
	, "loginAt"
	, "loginIPs"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (s *Sessions) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "sessionToken"
	, "userId"
	, "expiredAt"
	, "device"
	, "loginAt"
	, "loginIPs"
	`
}

// ToUpdateArray generate slice of update command
func (s *Sessions) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, s.SessionToken},
		A.X{`=`, 1, s.UserId},
		A.X{`=`, 2, s.ExpiredAt},
		A.X{`=`, 3, s.Device},
		A.X{`=`, 4, s.LoginAt},
		A.X{`=`, 5, s.LoginIPs},
	}
}

// IdxSessionToken return name of the index
func (s *Sessions) IdxSessionToken() int { //nolint:dupl false positive
	return 0
}

// SqlSessionToken return name of the column being indexed
func (s *Sessions) SqlSessionToken() string { //nolint:dupl false positive
	return `"sessionToken"`
}

// IdxUserId return name of the index
func (s *Sessions) IdxUserId() int { //nolint:dupl false positive
	return 1
}

// SqlUserId return name of the column being indexed
func (s *Sessions) SqlUserId() string { //nolint:dupl false positive
	return `"userId"`
}

// IdxExpiredAt return name of the index
func (s *Sessions) IdxExpiredAt() int { //nolint:dupl false positive
	return 2
}

// SqlExpiredAt return name of the column being indexed
func (s *Sessions) SqlExpiredAt() string { //nolint:dupl false positive
	return `"expiredAt"`
}

// IdxDevice return name of the index
func (s *Sessions) IdxDevice() int { //nolint:dupl false positive
	return 3
}

// SqlDevice return name of the column being indexed
func (s *Sessions) SqlDevice() string { //nolint:dupl false positive
	return `"device"`
}

// IdxLoginAt return name of the index
func (s *Sessions) IdxLoginAt() int { //nolint:dupl false positive
	return 4
}

// SqlLoginAt return name of the column being indexed
func (s *Sessions) SqlLoginAt() string { //nolint:dupl false positive
	return `"loginAt"`
}

// IdxLoginIPs return name of the index
func (s *Sessions) IdxLoginIPs() int { //nolint:dupl false positive
	return 5
}

// SqlLoginIPs return name of the column being indexed
func (s *Sessions) SqlLoginIPs() string { //nolint:dupl false positive
	return `"loginIPs"`
}

// ToArray receiver fields to slice
func (s *Sessions) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.SessionToken, // 0
		s.UserId,       // 1
		s.ExpiredAt,    // 2
		s.Device,       // 3
		s.LoginAt,      // 4
		s.LoginIPs,     // 5
	}
}

// FromArray convert slice to receiver fields
func (s *Sessions) FromArray(a A.X) *Sessions { //nolint:dupl false positive
	s.SessionToken = X.ToS(a[0])
	s.UserId = X.ToU(a[1])
	s.ExpiredAt = X.ToI(a[2])
	s.Device = X.ToS(a[3])
	s.LoginAt = X.ToI(a[4])
	s.LoginIPs = X.ToS(a[5])
	return s
}

// FromUncensoredArray convert slice to receiver fields
func (s *Sessions) FromUncensoredArray(a A.X) *Sessions { //nolint:dupl false positive
	s.SessionToken = X.ToS(a[0])
	s.UserId = X.ToU(a[1])
	s.ExpiredAt = X.ToI(a[2])
	s.Device = X.ToS(a[3])
	s.LoginAt = X.ToI(a[4])
	s.LoginIPs = X.ToS(a[5])
	return s
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (s *Sessions) FindOffsetLimit(offset, limit uint32, idx string) []Sessions { //nolint:dupl false positive
	var rows []Sessions
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Sessions{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (s *Sessions) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
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
func (s *Sessions) Total() int64 { //nolint:dupl false positive
	rows := s.Adapter.CallBoxSpace(s.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// SessionsFieldTypeMap returns key value of field name and key
var SessionsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`sessionToken`: Tt.String,
	`userId`:       Tt.Unsigned,
	`expiredAt`:    Tt.Integer,
	`device`:       Tt.String,
	`loginAt`:      Tt.Integer,
	`loginIPs`:     Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Users DAO reader/query struct
type Users struct {
	Adapter            *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                 uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	Email              string      `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	Password           string      `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	CreatedAt          int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy          uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt          int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy          uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt          int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	PasswordSetAt      int64       `json:"passwordSetAt" form:"passwordSetAt" query:"passwordSetAt" long:"passwordSetAt" msg:"passwordSetAt"`
	SecretCode         string      `json:"secretCode" form:"secretCode" query:"secretCode" long:"secretCode" msg:"secretCode"`
	SecretCodeAt       int64       `json:"secretCodeAt" form:"secretCodeAt" query:"secretCodeAt" long:"secretCodeAt" msg:"secretCodeAt"`
	VerificationSentAt int64       `json:"verificationSentAt" form:"verificationSentAt" query:"verificationSentAt" long:"verificationSentAt" msg:"verificationSentAt"`
	VerifiedAt         int64       `json:"verifiedAt" form:"verifiedAt" query:"verifiedAt" long:"verifiedAt" msg:"verifiedAt"`
	LastLoginAt        int64       `json:"lastLoginAt" form:"lastLoginAt" query:"lastLoginAt" long:"lastLoginAt" msg:"lastLoginAt"`
	FullName           string      `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
	UserName           string      `json:"userName" form:"userName" query:"userName" long:"userName" msg:"userName"`
	Country            string      `json:"country" form:"country" query:"country" long:"country" msg:"country"`
	Language           string      `json:"language" form:"language" query:"language" long:"language" msg:"language"`
	PropertyCount      int64       `json:"propertyCount" form:"propertyCount" query:"propertyCount" long:"propertyCount" msg:"propertyCount"`
	PropertyBought     int64       `json:"propertyBought" form:"propertyBought" query:"propertyBought" long:"propertyBought" msg:"propertyBought"`
}

// NewUsers create new ORM reader/query object
func NewUsers(adapter *Tt.Adapter) *Users {
	return &Users{Adapter: adapter}
}

// SpaceName returns full package and table name
func (u *Users) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableUsers) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (u *Users) SqlTableName() string { //nolint:dupl false positive
	return `"users"`
}

func (u *Users) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (u *Users) FindById() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{u.Id})
	if L.IsError(err, `Users.FindById failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexEmail return unique index name
func (u *Users) UniqueIndexEmail() string { //nolint:dupl false positive
	return `email`
}

// FindByEmail Find one by Email
func (u *Users) FindByEmail() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexEmail(), 0, 1, tarantool.IterEq, A.X{u.Email})
	if L.IsError(err, `Users.FindByEmail failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexUserName return unique index name
func (u *Users) UniqueIndexUserName() string { //nolint:dupl false positive
	return `userName`
}

// FindByUserName Find one by UserName
func (u *Users) FindByUserName() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexUserName(), 0, 1, tarantool.IterEq, A.X{u.UserName})
	if L.IsError(err, `Users.FindByUserName failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (u *Users) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "email"
	, "password"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "passwordSetAt"
	, "secretCode"
	, "secretCodeAt"
	, "verificationSentAt"
	, "verifiedAt"
	, "lastLoginAt"
	, "fullName"
	, "userName"
	, "country"
	, "language"
	, "propertyCount"
	, "propertyBought"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (u *Users) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "email"
	, "password"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "passwordSetAt"
	, "secretCode"
	, "secretCodeAt"
	, "verificationSentAt"
	, "verifiedAt"
	, "lastLoginAt"
	, "fullName"
	, "userName"
	, "country"
	, "language"
	, "propertyCount"
	, "propertyBought"
	`
}

// ToUpdateArray generate slice of update command
func (u *Users) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, u.Id},
		A.X{`=`, 1, u.Email},
		A.X{`=`, 2, u.Password},
		A.X{`=`, 3, u.CreatedAt},
		A.X{`=`, 4, u.CreatedBy},
		A.X{`=`, 5, u.UpdatedAt},
		A.X{`=`, 6, u.UpdatedBy},
		A.X{`=`, 7, u.DeletedAt},
		A.X{`=`, 8, u.PasswordSetAt},
		A.X{`=`, 9, u.SecretCode},
		A.X{`=`, 10, u.SecretCodeAt},
		A.X{`=`, 11, u.VerificationSentAt},
		A.X{`=`, 12, u.VerifiedAt},
		A.X{`=`, 13, u.LastLoginAt},
		A.X{`=`, 14, u.FullName},
		A.X{`=`, 15, u.UserName},
		A.X{`=`, 16, u.Country},
		A.X{`=`, 17, u.Language},
		A.X{`=`, 18, u.PropertyCount},
		A.X{`=`, 19, u.PropertyBought},
	}
}

// IdxId return name of the index
func (u *Users) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (u *Users) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxEmail return name of the index
func (u *Users) IdxEmail() int { //nolint:dupl false positive
	return 1
}

// SqlEmail return name of the column being indexed
func (u *Users) SqlEmail() string { //nolint:dupl false positive
	return `"email"`
}

// IdxPassword return name of the index
func (u *Users) IdxPassword() int { //nolint:dupl false positive
	return 2
}

// SqlPassword return name of the column being indexed
func (u *Users) SqlPassword() string { //nolint:dupl false positive
	return `"password"`
}

// IdxCreatedAt return name of the index
func (u *Users) IdxCreatedAt() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedAt return name of the column being indexed
func (u *Users) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (u *Users) IdxCreatedBy() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedBy return name of the column being indexed
func (u *Users) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (u *Users) IdxUpdatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedAt return name of the column being indexed
func (u *Users) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (u *Users) IdxUpdatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedBy return name of the column being indexed
func (u *Users) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (u *Users) IdxDeletedAt() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedAt return name of the column being indexed
func (u *Users) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxPasswordSetAt return name of the index
func (u *Users) IdxPasswordSetAt() int { //nolint:dupl false positive
	return 8
}

// SqlPasswordSetAt return name of the column being indexed
func (u *Users) SqlPasswordSetAt() string { //nolint:dupl false positive
	return `"passwordSetAt"`
}

// IdxSecretCode return name of the index
func (u *Users) IdxSecretCode() int { //nolint:dupl false positive
	return 9
}

// SqlSecretCode return name of the column being indexed
func (u *Users) SqlSecretCode() string { //nolint:dupl false positive
	return `"secretCode"`
}

// IdxSecretCodeAt return name of the index
func (u *Users) IdxSecretCodeAt() int { //nolint:dupl false positive
	return 10
}

// SqlSecretCodeAt return name of the column being indexed
func (u *Users) SqlSecretCodeAt() string { //nolint:dupl false positive
	return `"secretCodeAt"`
}

// IdxVerificationSentAt return name of the index
func (u *Users) IdxVerificationSentAt() int { //nolint:dupl false positive
	return 11
}

// SqlVerificationSentAt return name of the column being indexed
func (u *Users) SqlVerificationSentAt() string { //nolint:dupl false positive
	return `"verificationSentAt"`
}

// IdxVerifiedAt return name of the index
func (u *Users) IdxVerifiedAt() int { //nolint:dupl false positive
	return 12
}

// SqlVerifiedAt return name of the column being indexed
func (u *Users) SqlVerifiedAt() string { //nolint:dupl false positive
	return `"verifiedAt"`
}

// IdxLastLoginAt return name of the index
func (u *Users) IdxLastLoginAt() int { //nolint:dupl false positive
	return 13
}

// SqlLastLoginAt return name of the column being indexed
func (u *Users) SqlLastLoginAt() string { //nolint:dupl false positive
	return `"lastLoginAt"`
}

// IdxFullName return name of the index
func (u *Users) IdxFullName() int { //nolint:dupl false positive
	return 14
}

// SqlFullName return name of the column being indexed
func (u *Users) SqlFullName() string { //nolint:dupl false positive
	return `"fullName"`
}

// IdxUserName return name of the index
func (u *Users) IdxUserName() int { //nolint:dupl false positive
	return 15
}

// SqlUserName return name of the column being indexed
func (u *Users) SqlUserName() string { //nolint:dupl false positive
	return `"userName"`
}

// IdxCountry return name of the index
func (u *Users) IdxCountry() int { //nolint:dupl false positive
	return 16
}

// SqlCountry return name of the column being indexed
func (u *Users) SqlCountry() string { //nolint:dupl false positive
	return `"country"`
}

// IdxLanguage return name of the index
func (u *Users) IdxLanguage() int { //nolint:dupl false positive
	return 17
}

// SqlLanguage return name of the column being indexed
func (u *Users) SqlLanguage() string { //nolint:dupl false positive
	return `"language"`
}

// IdxPropertyCount return name of the index
func (u *Users) IdxPropertyCount() int { //nolint:dupl false positive
	return 18
}

// SqlPropertyCount return name of the column being indexed
func (u *Users) SqlPropertyCount() string { //nolint:dupl false positive
	return `"propertyCount"`
}

// IdxPropertyBought return name of the index
func (u *Users) IdxPropertyBought() int { //nolint:dupl false positive
	return 19
}

// SqlPropertyBought return name of the column being indexed
func (u *Users) SqlPropertyBought() string { //nolint:dupl false positive
	return `"propertyBought"`
}

// CensorFields remove sensitive fields for output
func (u *Users) CensorFields() { //nolint:dupl false positive
	u.Password = ``
	u.SecretCode = ``
	u.SecretCodeAt = 0
}

// ToArray receiver fields to slice
func (u *Users) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if u.Id != 0 {
		id = u.Id
	}
	return A.X{
		id,
		u.Email,              // 1
		u.Password,           // 2
		u.CreatedAt,          // 3
		u.CreatedBy,          // 4
		u.UpdatedAt,          // 5
		u.UpdatedBy,          // 6
		u.DeletedAt,          // 7
		u.PasswordSetAt,      // 8
		u.SecretCode,         // 9
		u.SecretCodeAt,       // 10
		u.VerificationSentAt, // 11
		u.VerifiedAt,         // 12
		u.LastLoginAt,        // 13
		u.FullName,           // 14
		u.UserName,           // 15
		u.Country,            // 16
		u.Language,           // 17
		u.PropertyCount,      // 18
		u.PropertyBought,     // 19
	}
}

// FromArray convert slice to receiver fields
func (u *Users) FromArray(a A.X) *Users { //nolint:dupl false positive
	u.Id = X.ToU(a[0])
	u.Email = X.ToS(a[1])
	u.Password = X.ToS(a[2])
	u.CreatedAt = X.ToI(a[3])
	u.CreatedBy = X.ToU(a[4])
	u.UpdatedAt = X.ToI(a[5])
	u.UpdatedBy = X.ToU(a[6])
	u.DeletedAt = X.ToI(a[7])
	u.PasswordSetAt = X.ToI(a[8])
	u.SecretCode = X.ToS(a[9])
	u.SecretCodeAt = X.ToI(a[10])
	u.VerificationSentAt = X.ToI(a[11])
	u.VerifiedAt = X.ToI(a[12])
	u.LastLoginAt = X.ToI(a[13])
	u.FullName = X.ToS(a[14])
	u.UserName = X.ToS(a[15])
	u.Country = X.ToS(a[16])
	u.Language = X.ToS(a[17])
	u.PropertyCount = X.ToI(a[18])
	u.PropertyBought = X.ToI(a[19])
	return u
}

// FromUncensoredArray convert slice to receiver fields
func (u *Users) FromUncensoredArray(a A.X) *Users { //nolint:dupl false positive
	u.Id = X.ToU(a[0])
	u.Email = X.ToS(a[1])
	u.CreatedAt = X.ToI(a[3])
	u.CreatedBy = X.ToU(a[4])
	u.UpdatedAt = X.ToI(a[5])
	u.UpdatedBy = X.ToU(a[6])
	u.DeletedAt = X.ToI(a[7])
	u.PasswordSetAt = X.ToI(a[8])
	u.VerificationSentAt = X.ToI(a[11])
	u.VerifiedAt = X.ToI(a[12])
	u.LastLoginAt = X.ToI(a[13])
	u.FullName = X.ToS(a[14])
	u.UserName = X.ToS(a[15])
	u.Country = X.ToS(a[16])
	u.Language = X.ToS(a[17])
	u.PropertyCount = X.ToI(a[18])
	u.PropertyBought = X.ToI(a[19])
	return u
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (u *Users) FindOffsetLimit(offset, limit uint32, idx string) []Users { //nolint:dupl false positive
	var rows []Users
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Users{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (u *Users) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
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
func (u *Users) Total() int64 { //nolint:dupl false positive
	rows := u.Adapter.CallBoxSpace(u.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// UsersFieldTypeMap returns key value of field name and key
var UsersFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                 Tt.Unsigned,
	`email`:              Tt.String,
	`password`:           Tt.String,
	`createdAt`:          Tt.Integer,
	`createdBy`:          Tt.Unsigned,
	`updatedAt`:          Tt.Integer,
	`updatedBy`:          Tt.Unsigned,
	`deletedAt`:          Tt.Integer,
	`passwordSetAt`:      Tt.Integer,
	`secretCode`:         Tt.String,
	`secretCodeAt`:       Tt.Integer,
	`verificationSentAt`: Tt.Integer,
	`verifiedAt`:         Tt.Integer,
	`lastLoginAt`:        Tt.Integer,
	`fullName`:           Tt.String,
	`userName`:           Tt.String,
	`country`:            Tt.String,
	`language`:           Tt.String,
	`propertyCount`:      Tt.Integer,
	`propertyBought`:     Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
