package app

import (
	"fmt"
	"log"

	"github.com/airsss993/opendrive/internal/config"
)

// Run initializes and starts the application with all required components
func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("failed to init config: %s", err)
	}
}
