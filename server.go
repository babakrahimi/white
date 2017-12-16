package main

import (
	"net/http"
	"log"
	"github.com/megaminx/white/app"
)

func main() {
	port := "9090"
	router := app.NewRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
