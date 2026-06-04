package main

import (
	"fmt"
	"log"
	"net/http"

	"interaction-simulator/internal/api"
	"interaction-simulator/internal/simulator/v2_evo"
)

func main() {
	// Initialize the simulation engine (v2_evo with lifespan and reproduction)
	engine := v2_evo.NewEngine()

	// Initialize the API server
	apiServer := api.NewServer(engine)

	// Setup routing
	mux := http.NewServeMux()
	apiServer.RegisterRoutes(mux)

	// Serve Static Files
	fs := http.FileServer(http.Dir("../../static")) // Note relative path from cmd/server to static/
	mux.Handle("/", fs)

	port := 8082
	fmt.Printf("Server starting on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
