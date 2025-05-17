package main

import "log"

func main() {
	// This is the main entry point for the application.
	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	_ = server.Start()
}
