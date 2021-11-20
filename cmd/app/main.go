package main

import (
	"fiber_blog/config"
	"fiber_blog/internal/app"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

func main() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(&cfg)
}
