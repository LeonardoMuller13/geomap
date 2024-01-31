package main

import (
	"log"

	"github.com/LeonardoMuller13/geomap/config"
	"github.com/LeonardoMuller13/geomap/simulator"
)

func main() {
	// Config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}

	//RUN SIMULATOR
	if cfg.Simulator.Status {
		simulator.New(cfg.Kafka)
	}

}
