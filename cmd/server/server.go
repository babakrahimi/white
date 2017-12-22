package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must set")
	}
	log.Fatal(http.ListenAndServe(":"+port, NewRouter()))
}
