package web

import (
	"net/http"
)

// RegistHandlers that handles request
func RegistHandlers() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/count", CounterHandler)
	http.HandleFunc("/lissajous", PlotLissajousHandler)
	http.HandleFunc("/sinc", PlotSincHandler)
}
