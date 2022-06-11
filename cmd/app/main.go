package main

import (
	"log"
)

func main() {
	logger := initLogger()

	conf, cErr := initConfig()
	if cErr != nil {
		logger.Error(cErr)
		return
	}

	app, err := boot(logger, conf.Application)
	if err != nil {
		logger.Fatalf("Something went wrong while utilizing the server. %v", err)
	}
	sv := initServer(app)

	for {
		go log.Fatal(sv.Listen(":8080"))
	}
}
