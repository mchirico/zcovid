package handles

import (
	"fmt"
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
