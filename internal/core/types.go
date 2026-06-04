package core

import "sync"

// Action represents a choice made by an agent.
type Action string

const (
	Cooperate Action = "Cooperate"
	Defect    Action = "Defect"
)

// Node represents an independent agent in the network.
type Node struct {
	ID       string                            `json:"id"`
	Strategy string                            `json:"strategy"`
	Score    float64                           `json:"score"`
	Age      int                               `json:"age"`
	Memory   map[string]Action                 `json:"-"` // Last action by each neighbor
	Meta     map[string]map[string]interface{} `json:"-"` // Per-opponent metadata for strategies
	Mu       sync.RWMutex                      `json:"-"`
}

// Edge represents a bidirectional connection.
type Edge struct {
	NodeA string `json:"source"`
	NodeB string `json:"target"`
}

// Graph holds the simulation state.
type Graph struct {
	Nodes map[string]*Node `json:"nodes"`
	Edges []Edge           `json:"edges"`
	Tick  int              `json:"tick"`
	Mu    sync.RWMutex     `json:"-"`
}

// SimConfig holds the initialization parameters for a simulation.
type SimConfig struct {
	Topology     string             `json:"topology"`
	Size         int                `json:"size"`
	Distribution map[string]float64 `json:"distribution"` // e.g., {"AlwaysCooperator": 0.5, "AlwaysCheater": 0.5}
}

// Simulator interface allows different engine variations (e.g. static, evolutionary)
// to be driven by the API.
type Simulator interface {
	AdvanceTick()
	GetState() *Graph
	Reset(config SimConfig)
}
