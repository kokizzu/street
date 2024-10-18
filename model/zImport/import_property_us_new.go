package zImport

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"street/model/mProperty/wcProperty"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kpango/fastime"
	"github.com/valyala/tsvreader"
)

func ReadPropertyUSSheet(adapter *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyUSSheet: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}

	defer file.Close()

	switch filepath.Ext(resourcePath) {
	case `.tsv`, `tsv`: 
		type propertyUS struct {
			SourceURL string `json:"source_url"`
			PostName string `json:"post_name"`
			MainListedPrice string `json:"main_listed_price"`
			Type string `json:"type"`
			Address1 string `json:"address1"`
			City string `json:"city"`
			State string `json:"state"`
			ZipCode string `json:"zip_code"`
			Bed string `json:"bed"`
			Baths string `json:"baths"`
			Sqft string `json:"sqft"`
			Description string `json:"description"`
			Parking string `json:"parking"`
			TotalSize string `json:"total_size"`
			PropertySubtype string `json:"property_subtype"`
			Listed string `json:"listed"`
			AgentName string `json:"agent_name"`
			Phone string `json:"phone"`
			BrokerName string `json:"broker_name"`
			SourceMLSNumbert string `json:"source_mls_numbert"`
			PriceHistory string `json:"price_history"`
			Image string `json:"image"`
		}

		var properties []propertyUS

		tsv := tsvreader.New(file)
		for tsv.Next() {
			_ = tsv.String()

			sourceURL := tsv.String()
			postName := tsv.String()
			mainListedPrice := tsv.String()
			typeP := tsv.String()
			address1 := tsv.String()
			city := tsv.String()
			state := tsv.String()
			zipCode := tsv.String()
			bed := tsv.String()
			baths := tsv.String()
			sqft := tsv.String()
			description := tsv.String()
			parking := tsv.String()
			totalSize := tsv.String()
			propertySubtype := tsv.String()
			listed := tsv.String()
			agentName := tsv.String()
			phone := tsv.String()
			brokerName := tsv.String()
			sourceMLSNumbert := tsv.String()
			priceHistory := tsv.String()
			image := tsv.String()

			properties = append(properties, propertyUS{
				SourceURL: S.Trim(sourceURL),
				PostName: S.Trim(postName),
				MainListedPrice: S.Trim(mainListedPrice),
				Type: S.Trim(typeP),
				Address1: S.Trim(address1),
				City: S.Trim(city),
				State: S.Trim(state),
				ZipCode: S.Trim(zipCode),
				Bed: S.Trim(bed),
				Baths: S.Trim(baths),
				Sqft: S.Trim(sqft),
				Description: S.Trim(description),
				Parking: S.Trim(parking),
				TotalSize: S.Trim(totalSize),
				PropertySubtype: S.Trim(propertySubtype),
				Listed: S.Trim(listed),
				AgentName: S.Trim(agentName),
				Phone: S.Trim(phone),
				BrokerName: S.Trim(brokerName),
				SourceMLSNumbert: S.Trim(sourceMLSNumbert),
				PriceHistory: S.Trim(priceHistory),
				Image: S.Trim(image),
			})
		}

		if len(properties) > 0 {
			properties = properties[1:] // remove 1st element
		} else {
			return
		}

		USStates := M.SS{
			"AL": "Alabama",
			"AK": "Alaska",
			"AS": "American Samoa",
			"AZ": "Arizona",
			"AR": "Arkansas",
			"CA": "California",
			"CO": "Colorado",
			"CT": "Connecticut",
			"DE": "Delaware",
			"DC": "District Of Columbia",
			"FM": "Federated States Of Micronesia",
			"FL": "Florida",
			"GA": "Georgia",
			"GU": "Guam",
			"HI": "Hawaii",
			"ID": "Idaho",
			"IL": "Illinois",
			"IN": "Indiana",
			"IA": "Iowa",
			"KS": "Kansas",
			"KY": "Kentucky",
			"LA": "Louisiana",
			"ME": "Maine",
			"MH": "Marshall Islands",
			"MD": "Maryland",
			"MA": "Massachusetts",
			"MI": "Michigan",
			"MN": "Minnesota",
			"MS": "Mississippi",
			"MO": "Missouri",
			"MT": "Montana",
			"NE": "Nebraska",
			"NV": "Nevada",
			"NH": "New Hampshire",
			"NJ": "New Jersey",
			"NM": "New Mexico",
			"NY": "New York",
			"NC": "North Carolina",
			"ND": "North Dakota",
			"MP": "Northern Mariana Islands",
			"OH": "Ohio",
			"OK": "Oklahoma",
			"OR": "Oregon",
			"PW": "Palau",
			"PA": "Pennsylvania",
			"PR": "Puerto Rico",
			"RI": "Rhode Island",
			"SC": "South Carolina",
			"SD": "South Dakota",
			"TN": "Tennessee",
			"TX": "Texas",
			"UT": "Utah",
			"VT": "Vermont",
			"VI": "Virgin Islands",
			"VA": "Virginia",
			"WA": "Washington",
			"WV": "West Virginia",
			"WI": "Wisconsin",
			"WY": "Wyoming",
		}

		for idx, v := range properties {
			property := wcProperty.NewPropertyUSMutator(adapter)
			property.SetLastPrice(v.MainListedPrice)
			property.SetPurpose(v.Type)
			property.SetFormattedAddress(v.Address1)
			property.SetCity(v.City)
			if st, ok := USStates[v.State]; ok {
				property.SetState(st)
			} else {
				property.SetState(v.State)
			}
			property.SetZip(v.ZipCode)

			if v.Bed != `` {
				bedArr := S.Split(v.Bed, ` `)
				property.SetBedroom(S.ToI(bedArr[0]))
			}

			if v.Baths != `` {
				bathArr := S.Split(v.Baths, ` `)
				property.SetBathroom(S.ToI(bathArr[0]))
			}

			if v.Sqft != `` {
				sqftArr := S.Split(v.Sqft, ` `)
				property.SetTotalSqft(S.ToF(sqftArr[0]))
			}

			property.SetNote(v.Description)
			if v.Parking != `` {
				parkingArr := S.Split(v.Parking, ` `)
				property.SetParking(S.ToF(parkingArr[0]))
			}

			property.SetCountryCode(`US`)
			property.SetCountyName(`United State of America`)
			property.SetHouseType(v.PropertySubtype)

			if v.PriceHistory != `` {
				priceHistoryArr := S.Split(v.PriceHistory, ` | `)
				var xprices []any = func() []any {
					var xx []any
					for _, v := range priceHistoryArr {
						xx = append(xx, v)
					}
					return xx
				}()

				if S.Contains(v.Type, `rent`) {
					property.SetPriceHistoriesRent(xprices)
				} else {
					property.SetPriceHistoriesSell(xprices)
				}
			}

			if isValidURL(v.Image) {
				property.SetImages([]any{v.Image})
			}

			property.SetCreatedAt(fastime.UnixNow())
			property.SetUpdatedAt(fastime.UnixNow())

			if !property.DoInsert() {
				msg := fmt.Sprintf("[%d] failed to insert property us data...", idx)
				L.Print(msg)
			}
		}
	case `.xlsx`:
		// TODO
	default:
		L.Print(`invalid file extension, must be xlsx / tsv / csv`)
		os.Exit(1)
	}
}

func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil || parsedURL.Scheme == `` || parsedURL.Host == `` {
		return false
	}
	
	return true
}