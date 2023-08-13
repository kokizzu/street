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
  "floorList": [
    {
      "type": "floor",
      "floor": 1,
      "beds": 1,
      "baths": 1,
      "rooms": [
        {
          "name": "bedroom",
          "sizeM2": 12,
          "unit": "m2"
        },
        {
          "name": "bathroom",
          "sizeM2": 8,
          "unit": "m2"
        }
      ],
      "planImageUrl": "/guest/files/D-___.jpg"
    }
  ]
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
		if assert.NotNil(t, out.Property) {
			out.Property.Adapter = nil
			in.Property.PriceHistoriesSell = []any{}
			in.Property.PriceHistoriesRent = []any{}
			assert.Equal(t, in.Property, *out.Property)
		}
	})
}
