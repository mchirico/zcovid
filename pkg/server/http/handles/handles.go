package handles

import (
	"bytes"
	"fmt"
	"github.com/mchirico/zcovid/pkg/echarts/gauge"
	"net/http"

)

var Count = 0
var CountStatus = 0

func BaseRoot(w http.ResponseWriter, r *http.Request) {

	version := "v0.0.1"

	switch r.Method {
	case "GET":
		Count += 1
		msg := fmt.Sprintf("\nversion: %v\nzcovid: %v\n", version, Count)
		w.Write([]byte(msg))
	case "POST":
		// msg := fmt.Sprintf("Hello world: POST: %v", r.FormValue("user"))
		w.Write([]byte("post"))
	default:
		w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
	}

}

func Gauge(w http.ResponseWriter, r *http.Request) {

	g := gauge.GaugeExamples{}
	buf := bytes.NewBufferString("")
	g.Examples(buf)
    w.Write(buf.Bytes())

}
