package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"interaction-simulator/internal/core"
)

type Server struct {
	Sim core.Simulator
	mu  sync.RWMutex
}

func NewServer(sim core.Simulator) *Server {
	return &Server{
		Sim: sim,
	}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/state", s.handleGetState)
	mux.HandleFunc("/api/tick", s.handleTick)
	mux.HandleFunc("/api/reset", s.handleReset)
}

func (s *Server) handleGetState(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.Sim.GetState())
}

func (s *Server) handleTick(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	s.mu.Lock()
	s.Sim.AdvanceTick()
	s.mu.Unlock()

	s.handleGetState(w, r)
}

func (s *Server) handleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var config core.SimConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		// Fallback to URL query params if JSON is not provided (for backward compatibility)
		config.Topology = r.URL.Query().Get("topology")
		sizeStr := r.URL.Query().Get("size")
		config.Size = 12
		if sz, err := strconv.Atoi(sizeStr); err == nil && sz > 0 {
			config.Size = sz
		}
	}

	// Ensure distribution exists if not provided
	if config.Distribution == nil {
		config.Distribution = map[string]float64{
			"AlwaysCooperator": 1.0,
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.Sim.Reset(config)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.Sim.GetState())
}
