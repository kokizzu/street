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
	FullName           = `fullName`

	LoginAt  = `loginAt`
	LoginIPs = `LoginIPs`
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
			{FullName, Tt.String},
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
			{LoginAt, Tt.Integer},
			{LoginIPs, Tt.String},
		},
		Unique1: SessionToken,
	},
}

const (
	TableActionLogs Ch.TableName = `actionLogs`

	RequestId  = `requestId`
	Error      = `error`
	ActorId    = `actorId`
	IpAddr4    = `ipAddr4`
	IpAddr6    = `ipAddr6`
	UserAgent  = `userAgent`
	Action     = `action`
	Traces     = `traces`
	StatusCode = `statusCode`
	Lat        = `lat`
	Long       = `long`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TableActionLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{RequestId, Ch.String},
			{ActorId, Ch.UInt64},
			{Action, Ch.String},
			{StatusCode, Ch.Int16},
			{Traces, Ch.String},
			{Error, Ch.String},
			{IpAddr4, Ch.IPv4},
			{IpAddr6, Ch.IPv6},
			{UserAgent, Ch.String},
			{Lat, Ch.Float64},
			{Long, Ch.Float64},
		},
		Orders: []string{CreatedAt, RequestId, ActorId, Action},
	},
}
