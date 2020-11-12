package http

import (
	"github.com/mchirico/zcovid/pkg/server/http/handles"
	"log"
	"net/http"
	_ "time/tzdata"
)

func SetupHandles() {

	http.HandleFunc("/", handles.BaseRoot)
	http.HandleFunc("/gauge", handles.Gauge)
	http.HandleFunc("/line", handles.Line)
	//http.Handle("/static/", http.StripPrefix("/static", fs))

}

func Server() {
	SetupHandles()
	log.Println("starting server... :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
