package mAuth

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

type (
	UserRegisterStat struct {
		Date  string `json:"date"`
		Count int64  `json:"count"`
	}
	RealtorStat struct {
		Date          string `json:"date"`
		TotalActivity int64  `json:"totalActivity"`
	}
	BuyerStat struct {
		Date          string `json:"date"`
		TotalActivity int64  `json:"totalActivity"`
	}
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
	UserName           = `userName`
	Country            = `country`  // 2-letters ISO country code
	Language           = `language` // 2-letters ISO country code (only EN and TW available, if empty assume EN)
	PropertyCount      = `propertyCount`
	PropertyBought     = `propertyBought`
)

const (
	TableSessions Tt.TableName = `sessions`

	SessionToken = `sessionToken`
	UserId       = `userId`
	ExpiredAt    = `expiredAt`
	Device       = `device`

	LoginAt  = `loginAt`
	LoginIPs = `loginIPs`
)

const (
	TableFeedbacks Tt.TableName = `feedbacks`

	UserMessage = `userMessage`
	AdminReply  = `adminReply`
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
			{UserName, Tt.String},
			{Country, Tt.String},
			{Language, Tt.String},
			{PropertyCount, Tt.Integer},
			{PropertyBought, Tt.Integer},
		},
		AutoIncrementId:  true,
		Unique1:          Email,
		Unique2:          UserName, // after migration setting default usernames code done
		AutoCensorFields: []string{Password, SecretCode, SecretCodeAt},
		Engine:           Tt.Memtx,
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
		Engine:  Tt.Memtx,
	},
	TableFeedbacks: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{CreatedBy, Tt.Unsigned}, // user
			{CreatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned}, // admin
			{UpdatedAt, Tt.Integer},
			{DeletedAt, Tt.Integer}, // solved or invalid
			{UserMessage, Tt.String},
			{AdminReply, Tt.String},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
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

	Latency = `latency` // in seconds

	RefId = `refId`
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
			{Latency, Ch.Float64},
			{RefId, Ch.UInt64},
		},
		Orders: []string{CreatedAt, RequestId, ActorId, Action},
	},
}
