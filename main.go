package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	loggraph "log-graph/log-graph"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed static/*
var content embed.FS

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config.yaml", "Path to config file")
}

func loadConfig(path string) (*loggraph.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg loggraph.Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	flag.Parse()

	cfg, err := loadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	patterns, err := loggraph.BuildPatterns(cfg)
	if err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	subFS, err := fs.Sub(content, "static")
	if err != nil {
		panic(err)
	}

	handler := loggraph.NewWSHandler(patterns)

	go handler.HandleInput()

	http.HandleFunc("/ws", handler.HandleConnections)
	http.Handle("/", http.FileServer(http.FS(subFS)))

	go handler.HandleBroadcast()

	fmt.Printf("Server started at :%d\n", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
