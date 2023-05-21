package mAuth

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TableUsers Tt.TableName = `users`

	Id                 = `id`
	Email              = `email`
	Password           = `password`
	CreatedAt          = `createdAt`
	CreatedBy          = `createdBy`
	UpdatedAt          = `updatedAt`
	UpdatedBy          = `updatedBy`
	DeletedAt          = `deletedAt`
	PasswordSetAt      = `passwordSetAt`
	SecretCode         = `secretCode`
	SecretCodeAt       = `secretCodeAt`
	VerificationSentAt = `verificationSentAt`
	VerifiedAt         = `verifiedAt`
	LastLoginAt        = `lastLoginAt`
)

const (
	TableSessions Tt.TableName = `sessions`

	SessionToken = `sessionToken`
	UserId       = `userId`
	ExpiredAt    = `expiredAt`
	Device       = `device`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableUsers: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{Email, Tt.String},
			{Password, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{PasswordSetAt, Tt.Integer},
			{SecretCode, Tt.String},
			{SecretCodeAt, Tt.Integer},
			{VerificationSentAt, Tt.Integer},
			{VerifiedAt, Tt.Integer},
			{LastLoginAt, Tt.Integer},
		},
		AutoIncrementId: true,
		Unique1:         Email,
	},
	TableSessions: {
		Fields: []Tt.Field{
			{SessionToken, Tt.String},
			{UserId, Tt.Unsigned},
			{ExpiredAt, Tt.Integer},
			{Device, Tt.String},
		},
		Unique1: SessionToken,
	},
}

const (
	TableUserLogs Ch.TableName = `userLogs`

	RequestId  = `requestId`
	Error      = `error`
	ActorId    = `actorId`
	IpAddr4    = `ipAddr4`
	IpAddr6    = `ipAddr6`
	UserAgent  = `userAgent`
	Action     = `action`
	Traces     = `traces`
	StatusCode = `statusCode`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TableUserLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{RequestId, Ch.String},
			{Error, Ch.String},
			{ActorId, Ch.UInt64},
			{IpAddr4, Ch.IPv4},
			{IpAddr6, Ch.IPv6},
			{UserAgent, Ch.String},
			{Action, Ch.String},
			{Traces, Ch.String},
			{StatusCode, Ch.Int16},
		},
		Orders: []string{ActorId, RequestId},
	},
}
