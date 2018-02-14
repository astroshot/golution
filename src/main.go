package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"learning_go/src/calc"
)

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func plotSinc(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	calc.DrawSinc(w)
	mu.Unlock()
}

func plotLissajous(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	calc.Lissajous(w)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", plotLissajous)
	http.HandleFunc("/sinc", plotSinc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
