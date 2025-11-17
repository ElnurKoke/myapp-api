package main

import (
	"elestial/config"
	"elestial/internal/app"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}
	app.Run(cfg)
}
