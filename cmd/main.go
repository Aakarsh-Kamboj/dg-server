package main

import (
	minioclient "dg-server/internal/minioClient"
	"log"
)

func main() {
	// This is the main entry point for the application.
	minioclient.InitMinIO()

	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	_ = server.Start()
}
