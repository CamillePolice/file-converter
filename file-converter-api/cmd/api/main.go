package main

import (
	"file-converter-api/internal/config"
	"file-converter-api/internal/routes"
	"file-converter-api/internal/server"
	"log"
)

func main() {
	// Initialize config
	cfg := config.NewServerConf()

	// Create server
	srv := server.NewServer(cfg)

	// Setup routes
	r := routes.NewRoutes(srv)
	r.Setup(srv.Echo())

	// Start server
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
