package zImport

import (
	"errors"
	"net/url"
	"os"
	"path/filepath"

	"street/model"
	"street/model/mProperty/wcProperty"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kpango/fastime"
	"github.com/valyala/tsvreader"
)

type OtherFacilityInfo struct {
	ParkingFeatures []string `json:"parkingFeatures,omitempty"`
	Gartereg        int64    `json:"gartereg,omitempty"`
	LotSize         string   `json:"lotSize,omitempty"`
	PostName        string   `json:"postName,omitempty"`
}

func ReadPropertyUS_TruliaCom(adapter *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyUS_TruliaCom: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}

	defer file.Close()
	type propertyUS struct {
		SourceURL        string
		PostName         string
		MainListedPrice  string
		TypePurpose      string
		Address1         string
		City             string
		State            string
		ZipCode          string
		Bed              string
		Baths            string
		Sqft             string
		Description      string
		Parking          string
		TotalSize        string
		PropertySubtype  string
		Listed           string
		AgentName        string
		Phone            string
		BrokerName       string
		SourceMLSNumbert string
		PriceHistory     string
		Image            string
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
			SourceURL:        S.Trim(sourceURL),
			PostName:         S.Trim(postName),
			MainListedPrice:  S.Trim(mainListedPrice),
			TypePurpose:      S.Trim(typeP),
			Address1:         S.Trim(address1),
			City:             S.Trim(city),
			State:            S.Trim(state),
			ZipCode:          S.Trim(zipCode),
			Bed:              S.Trim(bed),
			Baths:            S.Trim(baths),
			Sqft:             S.Trim(sqft),
			Description:      S.Trim(description),
			Parking:          S.Trim(parking),
			TotalSize:        S.Trim(totalSize),
			PropertySubtype:  S.Trim(propertySubtype),
			Listed:           S.Trim(listed),
			AgentName:        S.Trim(agentName),
			Phone:            S.Trim(phone),
			BrokerName:       S.Trim(brokerName),
			SourceMLSNumbert: S.Trim(sourceMLSNumbert),
			PriceHistory:     S.Trim(priceHistory),
			Image:            S.Trim(image),
		})
	}

	if len(properties) > 0 {
		properties = properties[1:] // remove 1st element
	} else {
		return
	}

	stat := &model.ImporterStat{Total: len(properties)}
	defer stat.Print(`last`)

	for _, v := range properties {
		stat.Print()

		property := wcProperty.NewPropertyUSMutator(adapter)
		if isValidURL(v.SourceURL) {
			propKey, err := getSheet001UniqPropKey(v.SourceURL)
			if err != nil {
				stat.Skip()
				continue
			}
			property.SetUniqPropKey(propKey)
		} else {
			stat.Skip()
			continue
		}
		property.SetLastPrice(v.MainListedPrice)
		property.SetPurpose(v.TypePurpose)
		property.SetFormattedAddress(v.Address1)
		property.SetCity(v.City)
		property.SetState(getUSStateByCode(v.State))
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
		property.SetCountryCode(`US`)
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

			if S.Contains(v.TypePurpose, `rent`) {
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
		property.SetCoord([]any{0, 0})

		stat.Ok(property.DoUpsert())

		stat.Print()
		propertyExtra := wcProperty.NewPropertyExtraUSMutator(adapter)
		propertyExtra.SetPropertyKey(property.UniqPropKey)

		facility := OtherFacilityInfo{
			ParkingFeatures: []string{v.Parking},
			PostName:        v.PostName,
		}

		facilityJson, err := json.Marshal(facility)
		if err != nil {
			stat.Warn(`facility marshal error`)
		} else {
			propertyExtra.SetFacilityInfo(string(facilityJson))
		}

		mlsDisclaimerInfo := MlsDisclaimerInfo{
			ListingBrokerName:   v.BrokerName,
			ListingBrokerNumber: v.SourceMLSNumbert,
			ListingAgentName:    v.AgentName,
			ListingAgentNumber:  v.Phone,
		}

		mlsDisclaimerInfoJson, err := json.Marshal(mlsDisclaimerInfo)
		if err != nil {
			stat.Warn(`mlsDisclaimerInfo marhsal error`)
		} else {
			propertyExtra.SetMlsDisclaimerInfo(string(mlsDisclaimerInfoJson))
		}

		stat.Ok(propertyExtra.DoUpsert())
	}
}

func ReadPropertyUS_ZillowCom(adapter *Tt.Adapter, resourcePath string) {
	defer subTaskPrint(`ReadPropertyUS_ZillowCom: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if L.IsError(err, `failed to open file "`+path+`"`) {
		os.Exit(1)
	}

	defer file.Close()

	type propertyUS struct {
		SourceURL       string
		PostName        string
		TypePurpose     string
		Address1        string
		City            string
		State           string
		ZipCode         string
		Latitude        string
		Longitude       string
		MainPrice       string
		Bed             string
		Bath            string
		Sqft            string
		AgentName       string
		AgentPhone      string
		BrokerName      string
		BrokerNumber    string
		NMLSId          string
		LotSize         string
		Parking         string
		ParkingFeatures string
		Gartereg        string
		PropertyType    string
		LastUpdate      string
		LastChecked     string
		HistoryListed   string
		Image           string
	}

	var properties []propertyUS

	tsv := tsvreader.New(file)
	for tsv.Next() {
		_ = tsv.String()

		sourceUrl := tsv.String()
		postName := tsv.String()
		typePurpose := tsv.String()
		address1 := tsv.String()
		city := tsv.String()
		state := tsv.String()
		zipCode := tsv.String()
		latitude := tsv.String()
		longitude := tsv.String()
		mainPrice := tsv.String()
		bed := tsv.String()
		bath := tsv.String()
		sqft := tsv.String()
		agentName := tsv.String()
		agentPhone := tsv.String()
		brokerName := tsv.String()
		brokerNumber := tsv.String()
		nmlsId := tsv.String()
		lotSize := tsv.String()
		parking := tsv.String()
		parkingFeatures := tsv.String()
		gartereg := tsv.String()
		propertyType := tsv.String()
		lastUpdate := tsv.String()
		lastChecked := tsv.String()
		historyListed := tsv.String()
		image := tsv.String()

		properties = append(properties, propertyUS{
			SourceURL:       S.Trim(sourceUrl),
			PostName:        S.Trim(postName),
			TypePurpose:     S.Trim(typePurpose),
			Address1:        S.Trim(address1),
			City:            S.Trim(city),
			State:           S.Trim(state),
			ZipCode:         S.Trim(zipCode),
			Latitude:        S.Trim(latitude),
			Longitude:       S.Trim(longitude),
			MainPrice:       S.Trim(mainPrice),
			Bed:             S.Trim(bed),
			Bath:            S.Trim(bath),
			Sqft:            S.Trim(sqft),
			AgentName:       S.Trim(agentName),
			AgentPhone:      S.Trim(agentPhone),
			BrokerName:      S.Trim(brokerName),
			BrokerNumber:    S.Trim(brokerNumber),
			NMLSId:          S.Trim(nmlsId),
			LotSize:         S.Trim(lotSize),
			Parking:         S.Trim(parking),
			ParkingFeatures: S.Trim(parkingFeatures),
			Gartereg:        S.Trim(gartereg),
			PropertyType:    S.Trim(propertyType),
			LastUpdate:      S.Trim(lastUpdate),
			LastChecked:     S.Trim(lastChecked),
			HistoryListed:   S.Trim(historyListed),
			Image:           S.Trim(image),
		})
	}

	if len(properties) > 0 {
		properties = properties[1:] // remove 1st element
	} else {
		return
	}

	stat := &model.ImporterStat{Total: len(properties) * 2}
	defer stat.Print(`last`)

	for _, v := range properties {
		stat.Print()

		property := wcProperty.NewPropertyUSMutator(adapter)
		if isValidURL(v.SourceURL) {
			propKey, err := getSheet002UniqPropKey(v.SourceURL)
			if err != nil {
				stat.Skip()
				continue
			}
			property.SetUniqPropKey(propKey)
		} else {
			stat.Skip()
			continue
		}
		property.SetPurpose(v.TypePurpose)
		property.SetFormattedAddress(v.Address1)
		property.SetCity(v.City)
		property.SetState(getUSStateByCode(v.State))
		property.SetZip(v.ZipCode)

		coord := [2]any{0, 0}

		if v.Latitude != `` {
			latitude := S.Replace(v.Latitude, `,`, `.`)
			coord[0] = S.ToF(latitude)
		}

		if v.Longitude != `` {
			longitude := S.Replace(v.Latitude, `,`, `.`)
			coord[1] = S.ToF(longitude)
		}

		property.SetCoord(coord[:])

		if v.MainPrice != `` {
			price := S.Replace(S.Replace(v.MainPrice, `$`, ``), `.`, ``)
			property.SetLastPrice(price)
		}

		if !(S.Contains(v.Bed, `--`) || S.Contains(v.Bed, `"`) || S.Contains(v.Bed, `#`)) {
			if S.Contains(v.Bed, `.`) {
				bedStr := S.Split(v.Bed, `.`)[0]
				property.SetBedroom(S.ToI(S.Trim(bedStr)))
			}
		}

		bathStr := S.Trim(S.Replace(v.Bath, `,`, ``))
		property.SetBathroom(S.ToI(bathStr))

		property.SetTotalSqft(S.ToF(v.Sqft))

		if v.Parking != `` {
			parkingArr := S.Split(v.Parking, ` `)
			if len(parkingArr) == 3 {
				property.SetParking(S.ToF(parkingArr[2]))
			}
		}

		property.SetHouseType(v.PropertyType)
		if isValidURL(v.Image) {
			property.SetImages([]any{v.Image})
		}

		property.SetCreatedAt(fastime.UnixNow())
		property.SetUpdatedAt(fastime.UnixNow())
		property.SetCoord([]any{0, 0})

		stat.Ok(property.DoUpsert())

		stat.Print()
		propertyExtra := wcProperty.NewPropertyExtraUSMutator(adapter)
		propertyExtra.SetPropertyKey(property.UniqPropKey)

		facility := OtherFacilityInfo{
			ParkingFeatures: S.Split(v.ParkingFeatures, `, `),
			Gartereg:        S.ToI(v.Gartereg),
			LotSize:         v.LotSize,
			PostName:        v.PostName,
		}

		facilityJson, err := json.Marshal(facility)
		if err != nil {
			stat.Warn(`facility marshal error`)
		} else {
			propertyExtra.SetFacilityInfo(string(facilityJson))
		}

		mlsDisclaimerInfo := MlsDisclaimerInfo{
			ListingBrokerName:   v.BrokerName,
			ListingBrokerNumber: v.BrokerNumber,
			ListingAgentName:    v.AgentName,
			ListingAgentNumber:  v.AgentPhone,
		}

		mlsDisclaimerInfoJson, err := json.Marshal(mlsDisclaimerInfo)
		if err != nil {
			stat.Warn(`mlsDisclaimerInfo marhsal error`)
		} else {
			propertyExtra.SetMlsDisclaimerInfo(string(mlsDisclaimerInfoJson))
		}

		stat.Ok(propertyExtra.DoUpsert())
	}
}

func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	if err != nil || parsedURL.Scheme == `` || parsedURL.Host == `` {
		return false
	}

	return true
}

func getSheet001UniqPropKey(sourceURL string) (string, error) {
	parsedUrl, err := url.Parse(sourceURL)
	if err != nil {
		return ``, errors.New(`invalid URL`)
	}

	pathSegments := S.Split(parsedUrl.Path, `/`)
	if len(pathSegments) != 3 {
		return ``, errors.New(`cannot find unique property key`)
	}

	propKey := pathSegments[len(pathSegments)-1]
	propKeyArr := S.Split(propKey, `-`)

	return propKeyArr[len(propKeyArr)-1], nil
}

func getSheet002UniqPropKey(sourceURL string) (string, error) {
	parsedUrl, err := url.Parse(sourceURL)
	if err != nil {
		return ``, errors.New(`invalid URL`)
	}

	pathSegments := S.Split(parsedUrl.Path, `/`)
	if len(pathSegments) != 5 {
		return ``, errors.New(`cannot find unique property key`)
	}

	propKey := pathSegments[len(pathSegments)-2]

	return propKey, nil
}

func getUSStateByCode(stateCode string) string {
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

	if st, ok := USStates[stateCode]; ok {
		return st
	}

	return stateCode
}
