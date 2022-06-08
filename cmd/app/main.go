package main

import (
	"log"
)

func main() {
	logger := initLogger()
	app, err := boot(logger)
	if err != nil {
		logger.Fatalf("Something went wrong while utilizing the server. %v", err)
	}
	sv := initServer(app)

	go log.Fatal(sv.Listen("8080"))
}
