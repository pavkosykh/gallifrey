package main

import (
	"log"

	"github.com/pavkosykh/gallifrey/config"
	"github.com/pavkosykh/gallifrey/internal/cli"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	cli.Run(cfg)
}
