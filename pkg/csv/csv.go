package csv

import (
	"encoding/csv"
	"github.com/mchirico/zcovid/pkg/httputils"
	"log"
	"strings"
)

func GetCSVRecords(url string) (records [][]string, err error) {

	h := httputils.NewHTTP()
	r, err := h.Get(url)
	if err != nil {
		return nil, nil
	}

	rr := csv.NewReader(strings.NewReader(string(r)))
	records, err = rr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records, err

}
