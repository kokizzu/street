package mBusiness

import "github.com/kokizzu/gotro/D/Tt"

const (
	TableRevenue Tt.TableName = `revenue`
	TableSales Tt.TableName = `sales`

	Id = `id`
	PropertyId = `propertyId`
	PropertyBought = `propertyBought`
	Price = `price`
	RealtorId = `realtorId`
	BuyerId = `buyerId` // fill with 0 if user's email is empty
	BuyerEmail = `buyerEmail`
	EmailNotFound = `emailNotFound`
	SalesDate = `salesDate`
	CreatedAt = `createdAt`
	CreatedBy = `createdBy`
	UpdatedAt = `updatedAt`
	UpdatedBy = `updatedBy`
	DeletedAt = `deletedAt`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableSales: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{PropertyId, Tt.Unsigned},
			{RealtorId, Tt.Unsigned},
			{BuyerId, Tt.Unsigned},
			{Price, Tt.String},
			{BuyerEmail, Tt.String},
			{SalesDate, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		Engine:  Tt.Memtx,
		Uniques: []string{PropertyId, SalesDate},
	},
	TableRevenue: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{RealtorId, Tt.Unsigned},
			{PropertyId, Tt.Unsigned},
			{PropertyBought, Tt.Integer},
			{BuyerEmail, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		Engine: Tt.Memtx,
		Uniques: []string{RealtorId, PropertyId},
	},
}