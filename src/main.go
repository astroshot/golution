package main

import (
	"fmt"
	"log"
	"net/http"

	web "modules/src/web"
)

func main() {
	fmt.Println("Initializing...")
	web.RegistHandlers()
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
