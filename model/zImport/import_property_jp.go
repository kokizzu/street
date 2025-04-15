package zImport

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"street/model"
	"street/model/mProperty"
	"street/model/mProperty/wcProperty"
	"street/model/mStorage/wcStorage"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kpango/fastime"
	"github.com/valyala/tsvreader"
)

func ReadPropertyJP(conn *Tt.Adapter, resourcePath string, uploadDir string) {
	defer subTaskPrint(`ReadPropertyJP: import property data`)()

	path, err := filepath.Abs(resourcePath)
	if L.IsError(err, `failed to get path to file "`+resourcePath+`"`) {
		os.Exit(1)
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(`failed to open file "`+path+`" : `, err)
		os.Exit(1)
	}
	defer file.Close()

	type propertyJp struct {
		id             string
		sellingPrice   string
		buildingName   string
		city           string
		district       string
		address        string
		roomType       string
		sizeM2         string
		floor          string
		year           string
		month          string
		closestStation string
		minToStation   string
		status         string
		coord          string
		file3d         string
	}

	var properties []propertyJp

	tsv := tsvreader.New(file)
	for tsv.Next() {
		var propertyJp propertyJp

		propertyJp.id = tsv.String()
		propertyJp.sellingPrice = tsv.String()
		propertyJp.buildingName = tsv.String()
		propertyJp.city = tsv.String()
		propertyJp.district = tsv.String()
		propertyJp.address = tsv.String()
		propertyJp.roomType = tsv.String()
		propertyJp.sizeM2 = tsv.String()
		propertyJp.floor = tsv.String()
		propertyJp.year = tsv.String()
		propertyJp.month = tsv.String()
		propertyJp.closestStation = tsv.String()
		propertyJp.minToStation = tsv.String()
		propertyJp.status = tsv.String()
		propertyJp.coord = tsv.String()
		propertyJp.file3d = tsv.String()

		properties = append(properties, propertyJp)
	}

	properties = properties[1:]

	stat := &model.ImporterStat{
		Total: len(properties),
	}
	defer stat.Print(`last`)

	rgxFileId := regexp.MustCompile(`drive\.google\.com\/file\/d\/([a-zA-Z0-9_-]+)`)
	rgxConfirmToken := regexp.MustCompile(`confirm=([0-9A-Za-z_]+)`)

	for _, v := range properties {
		stat.Print()

		prop := wcProperty.NewPropertyMutator(conn)

		uniqPropKey := v.id + `_jp`

		prop.UniqPropKey = uniqPropKey
		if !prop.FindByUniqPropKey() {
			prop.SetUniqPropKey(uniqPropKey)
		}
		prop.SetCity(v.city)
		prop.SetDistrict(v.district)
		prop.SetAddress(v.address)
		prop.SetSizeM2(v.sizeM2)

		floorSlice := S.Split(v.floor, `/`)

		if len(floorSlice) == 2 {
			prop.SetNumberOfFloors(floorSlice[1])
		} else {
			prop.SetNumberOfFloors(floorSlice[0])
		}

		prop.SetYearBuilt(S.ToI(v.year))

		attrByt, err := json.Marshal(mProperty.PropertyAttribute{
			ClosestStation:   v.closestStation,
			MinutesToStation: S.ToI(v.minToStation),
		})
		if err != nil {
			fmt.Println(`failed to marshal property attribute: `, err)
			stat.Skip()
		}
		prop.SetAttribute(string(attrByt))

		price := convertJPYToUSD(v.sellingPrice)
		prop.SetLastPrice(price)

		coordStrSlice := S.Split(v.coord, ",")
		if len(coordStrSlice) == 2 {
			coordAnySlice := []any{
				S.ToF(S.Trim(coordStrSlice[0])),
				S.ToF(S.Trim(coordStrSlice[1])),
			}
			prop.SetCoord(coordAnySlice)
		}
		prop.SetCountryCode(`JP`)

		if prop.DoUpsert() {
			if v.file3d != `` {
				matchFileId := rgxFileId.FindStringSubmatch(v.file3d)
				if len(matchFileId) > 1 {
					fileId := matchFileId[1]
					countryPropId := fmt.Sprintf("%s:%d", `JP`, prop.Id)
					save3dFile(conn, fileId, countryPropId, uploadDir, prop.Id, rgxConfirmToken)
				}
			}
		} else {
			stat.Ok(false)
			continue
		}

		stat.Ok(true)
	}
}

func save3dFile(conn *Tt.Adapter, fileId, countryPropId, uploadDir string, propId uint64, rgxConfirmToken *regexp.Regexp) error {
	url := fmt.Sprintf("https://drive.google.com/uc?export=download&id=%s", fileId)
	client := &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		L.LOG.Error("failed to download 3d file", err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		L.LOG.Error("failed to read 3d file", err)
		return err
	}
	bodyStr := string(bodyBytes)

	// Try to extract confirmation token
	match := rgxConfirmToken.FindStringSubmatch(bodyStr)

	downloadUrl := url
	if len(match) > 2 {
		confirmToken := match[1]
		downloadUrl = fmt.Sprintf("https://drive.google.com/uc?export=download&confirm=%s&id=%s", confirmToken, fileId)
		fmt.Println("Using confirmation token:", confirmToken)
	} else {
		fmt.Println("No confirmation token found, downloading directly...")
	}

	respDownload, err := client.Get(downloadUrl)
	if err != nil {
		L.LOG.Error("failed to download 3d file", err)
		return err
	}
	defer respDownload.Body.Close()

	fileBytes, err := io.ReadAll(respDownload.Body)
	if err != nil {
		L.LOG.Error("failed to read 3d file", err)
		return err
	}

	fileName := fmt.Sprintf("%s_%d.glb", fileId, propId)

	designFile := wcStorage.NewDesignFilesMutator(conn)
	designFile.CountryPropId = countryPropId
	if !designFile.FindByCountryPropId() {
		designFile.SetCountryPropId(countryPropId)
		designFile.SetCreatedAt(fastime.UnixNow())
		if !designFile.DoInsert() {
			return errors.New("failed to insert design file")
		}
	}

	pathName := uploadDir + fileName
	writer, err := os.Create(pathName)
	if err != nil {
		L.LOG.Error("failed to create file", err)
		return err
	}

	_, err = writer.Write(fileBytes)
	if err != nil {
		L.LOG.Error("failed to write file", err)
		return err
	}

	err = writer.Close()
	if err != nil {
		L.LOG.Error("failed to close file", err)
		return err
	}

	designFile.SetFilePath(fileName)
	if !designFile.DoUpdateById() {
		return errors.New("failed to update design file")
	}

	return nil
}

func convertJPYToUSD(jpyStr string) string {
	jpyStr = S.Trim(jpyStr)
	jpyStr = S.Replace(jpyStr, `.`, ``)

	jpyFloat := S.ToF(jpyStr)

	var exchangeRate float64 = 0.0067

	usdAmount := (jpyFloat * exchangeRate)

	return fmt.Sprintf("%.2f", usdAmount)
}

func randChar(length int) string {
	const letterBytes = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
