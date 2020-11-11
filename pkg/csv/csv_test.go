package csv

import (
	"fmt"
	"testing"
)

func TestGetCSVRecords(t *testing.T) {
	url := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_US.csv"

	records, err := GetCSVRecords(url)
	if err != nil {
		t.Fatalf("TestCSV failed: %s\n", err)
	}
	fmt.Println(records[0][20])
}
