package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/v1k0d3n/fakefish/internal/server"
	"github.com/v1k0d3n/fakefish/pkg/config"
)

func main() {
	configPath := flag.String("c", "config.yaml", "Path to the configuration file")
	flag.Parse()

	absConfigPath, err := filepath.Abs(*configPath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}

	cfg, err := config.LoadConfig(absConfigPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	srv := server.NewServer(cfg)
	srv.Start()
}
