package zImport

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"github.com/valyala/tsvreader"
)

func GoogleSheetTranslationToJson(docId string) {
	// wget --output-file="logs.csv" "https://docs.google.com/spreadsheets/d/<document_id>/export?format=tsv&gid=<sheet_id>" -O "downloaded_content.tsv"
	// download and convert document id to tsv
	url := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/export?format=tsv&gid=%d", docId, 0)
	resp, err := http.Get(url)
	L.PanicIf(err, `GoogleSheetTranslationToJson.http.Get`)
	if resp.StatusCode != http.StatusOK {
		L.Print(`GoogleSheetTranslationToJson.http.Get:`, resp.StatusCode)
		return
	}
	if resp == nil || resp.Body == nil {
		L.Print(`GoogleSheetTranslationToJson.http.Get: empty response`)
		return
	}

	// convert tsv to json
	// columns: key	EN	TW
	r := tsvreader.New(resp.Body)
	m := map[string]string{}
	for r.Next() {
		key := r.String()
		if key == `` { // skip if have no key (haoji make it duplicate)
			for r.HasCols() {
				_ = r.String()
			}
			continue
		}
		en := r.String()
		tw := r.String()
		m[key] = en
		m[key+`TW`] = tw
		// ignore columnt 4+
		for r.HasCols() {
			_ = r.String()
		}
	}
	err = r.Error()
	if err != nil && !strings.Contains(err.Error(), `cannot find newline at the end of row`) {
		L.PanicIf(err, `GoogleSheetTranslationToJson.tsvreader.Error`)
	}
	fo, err := os.Create(`./svelte/translation.json`)
	defer fo.Close()
	L.PanicIf(err, `GoogleSheetTranslationToJson.os.Create`)
	_, err = fo.WriteString(X.ToJson(m))
	L.PanicIf(err, `GoogleSheetTranslationToJson.WriteString`)
}
