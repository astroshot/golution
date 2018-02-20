package main

import (
	"fmt"
	"log"
	"net/http"

	web "learning_go/src/web"
)

func main() {
	fmt.Println("Initializing...")
	http.HandleFunc("/", web.DefaultHandler)
	http.HandleFunc("/count", web.CounterHandler)
	http.HandleFunc("/lissajous", web.PlotLissajousHandler)
	http.HandleFunc("/sinc", web.PlotSincHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
