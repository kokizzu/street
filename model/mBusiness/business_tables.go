package mBusiness

import (
	"time"

	"github.com/kokizzu/gotro/D/Tt"
)

type Cache struct {
	CacheName string
	CacheUnixTime int64
	CacheData any
}

// Return expired (true) if cache time more than 1 hour
func (c *Cache) IsExpired() bool {
	cacheTime := time.Unix(c.CacheUnixTime, 0)

	return time.Since(cacheTime) > time.Hour
}

type Revenue struct {
	Revenue int64 `json:"revenue"`
	PropertyBought int64 `json:"propertyBought"`
	SalesDate string `json:"salesDate"`
}

type Order struct {
	Revenue int64 `json:"revenue"`
	TotalTransaction int64 `json:"totalTransaction"`
	SalesDate string `json:"salesDate"`
}

type Buyer struct {
	PropertyId uint64 `json:"propertyId"`
	BuyerEmail string `json:"buyerEmail"`
	PropertyPrice int64 `json:"propertyPrice"`
	SalesDate string `json:"salesDate"`
}

const (
	TableSales Tt.TableName = `sales`

	Id = `id`
	PropertyId = `propertyId`
	PropertyBought = `propertyBought`
	Price = `price`
	RealtorId = `realtorId`
	BuyerId = `buyerId` // fill with 0 if user's email is empty
	BuyerEmail = `buyerEmail`
	EmailNotFound = `emailNotFound`
	PropertyCountry = `propertyCountry`
	SalesDate = `salesDate`
	CreatedAt = `createdAt`
	CreatedBy = `createdBy`
	UpdatedAt = `updatedAt`
	UpdatedBy = `updatedBy`
	DeletedAt = `deletedAt`
)
func IsValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	
	return err == nil
}

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableSales: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{PropertyId, Tt.Unsigned},
			{RealtorId, Tt.Unsigned},
			{PropertyCountry, Tt.String},
			{BuyerId, Tt.Unsigned},
			{Price, Tt.String},
			{BuyerEmail, Tt.String},
			{EmailNotFound, Tt.String},
			{SalesDate, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		Engine:  Tt.Memtx,
		Uniques: []string{PropertyId, SalesDate},
		AutoIncrementId: true,
	},
}