package main

import (
	"log"

	"github.com/David-Kalashir/crs-server/config"
	"github.com/David-Kalashir/crs-server/internal/app"
)

func main() {

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)

}
