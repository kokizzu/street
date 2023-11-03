package domain

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"

	"street/model/mProperty/rqProperty"
)

func TestRealtorUpsertProperty(t *testing.T) {
	const propJson = `{
  "formattedAddress": "Jl. Gayungsari Tim. III No.40, Menanggal, Kec. Gayungan, Surabaya, Jawa Timur 60234, Indonesia",
  "coord": [
    -7.338162399999999,
    12.4456
  ],
  "houseType": "house",
  "purpose": "rent",
  "images": [
    "/guest/files/C-___.png"
  ],
  "beds": 3,
  "baths": 4,
  "sizeM2": "200",
  "mainUse": "swimming pool",
  "note": "near railroad",
  "numberOfFloors": "1",
  "approvalState": "pending"
}`
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {

		in := RealtorUpsertPropertyIn{
			RequestCommon: testAdminRequestCommon(RealtorUpsertPropertyAction),
			Property:      rqProperty.Property{},
		}

		err := json.Unmarshal([]byte(propJson), &in.Property)
		assert.NoError(t, err)

		out := d.RealtorUpsertProperty(&in)

		assert.Empty(t, out.Error)

		assert.NotZero(t, out.Property.Id)
		in.Property.Id = out.Property.Id
		assert.NotEmpty(t, out.Property.UniqPropKey)
		in.Property.UniqPropKey = out.Property.UniqPropKey
		assert.NotZero(t, out.Property.CreatedAt)
		in.Property.CreatedAt = out.Property.CreatedAt
		assert.Equal(t, out.Property.Id, out.Property.CreatedBy)
		in.Property.CreatedBy = out.Property.CreatedBy
		assert.NotZero(t, out.Property.UpdatedAt)
		in.Property.UpdatedAt = out.Property.UpdatedAt
		assert.Equal(t, out.Property.Id, out.Property.UpdatedBy)
		in.Property.UpdatedBy = out.Property.UpdatedBy

		if assert.NotNil(t, out.Property) {
			out.Property.Adapter = nil
			in.Property.PriceHistoriesSell = []any{}
			in.Property.PriceHistoriesRent = []any{}
			in.Property.FloorList = []any{}
			in.Property.OtherFees = []any{}
			in.Property.ImageLabels = []any{}
			assert.Equal(t, in.Property, *out.Property)
		}
	})
}
