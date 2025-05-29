package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	loggraph "log-graph/log-graph"
	"net/http"
)

//go:embed static/*
var content embed.FS

func main() {
	go loggraph.HandleInput()

	subFS, err := fs.Sub(content, "static")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/ws", loggraph.HandleConnections)
	http.Handle("/", http.FileServer(http.FS(subFS)))

	go loggraph.HandleBroadcast()

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
