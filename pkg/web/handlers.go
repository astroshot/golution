package web

import (
	"fmt"
	"net/http"
	"sync"

	"golution/pkg/calc"
)

var (
	mu    sync.Mutex
	count int
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func PlotSincHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	r.Header.Set("Content-Type", "image/svg+xml")
	calc.DrawSinc(w)
	mu.Unlock()
}

func PlotLissajousHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	calc.Lissajous(w)
	mu.Unlock()
}
