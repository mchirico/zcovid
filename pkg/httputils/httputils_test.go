package httputils

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_US.csv"

	//key := "X-API-Key"
	//value := "1vtlJSvzaaB6bTjJKzyakYnjnxrRzM22Ex3j2SDR"

	h := NewHTTP()
	//h.Header(key, value)

	r, err := h.Get(url)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Println(string(r))

}
