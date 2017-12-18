package main

import (
	"github.com/megaminx/white/cmd/server"
	"log"
	"net/http"
)

func main() {
	c, err := server.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":"+c.Port, server.GetRouter()))
}
