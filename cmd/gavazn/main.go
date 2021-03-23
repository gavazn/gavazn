package main

import (
	"log"

	"github.com/Gavazn/Gavazn/config"
	"github.com/Gavazn/Gavazn/server"
)

func main() {
	if err := server.Start(config.Get("PORT")); err != nil {
		log.Fatal(err)
	}
}
