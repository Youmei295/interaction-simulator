package main

import (
	"fmt"
	"log"
	"net/http"

	"interaction-simulator/internal/api"
	"interaction-simulator/internal/simulator/v1_static"
)

func main() {
	// Initialize the simulation engine (v1_static for now)
	engine := v1_static.NewEngine()

	// Initialize the API server
	apiServer := api.NewServer(engine)

	// Setup routing
	mux := http.NewServeMux()
	apiServer.RegisterRoutes(mux)

	// Serve Static Files
	fs := http.FileServer(http.Dir("../../static")) // Note relative path from cmd/server to static/
	mux.Handle("/", fs)

	port := 8081
	fmt.Printf("Server starting on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
