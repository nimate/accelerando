package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port         int
	ShutdownSecs int
	ServiceName  string
}

func loadConfig() (Config, error) {
	cfg := Config{
		Port:         8080,
		ShutdownSecs: 10,
		ServiceName:  "demo-service",
	}

	if raw := os.Getenv("PORT"); raw != "" {
		port, err := strconv.Atoi(raw)
		if err != nil {
			return Config{}, fmt.Errorf("parse PORT: %w", err)
		}
		cfg.Port = port
	}

	if raw := os.Getenv("SHUTDOWN_SECS"); raw != "" {
		secs, err := strconv.Atoi(raw)
		if err != nil {
			return Config{}, fmt.Errorf("parse SHUTDOWN_SECS: %w", err)
		}
		cfg.ShutdownSecs = secs
	}

	if raw := os.Getenv("SERVICE_NAME"); raw != "" {
		cfg.ServiceName = raw
	}

	return cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	log.Printf("config: %+v", cfg)
}
