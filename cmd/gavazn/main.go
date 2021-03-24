package main

import (
	"log"

	"github.com/Gavazn/Gavazn/config"
	"github.com/Gavazn/Gavazn/database"
	"github.com/Gavazn/Gavazn/server"
)

func main() {
	// connect to database
	if err := database.Connect(config.Get("MONGO_HOST"), config.Get("MONGO_DATABASE"), config.Get("MONGO_USER"), config.Get("MONGO_PASSWORD")); err != nil {
		log.Fatal(err)
	}

	// start web server
	if err := server.Start(config.Get("PORT")); err != nil {
		log.Fatal(err)
	}
}
