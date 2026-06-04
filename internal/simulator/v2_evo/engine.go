package v2_evo

import (
	"interaction-simulator/internal/core"
	"interaction-simulator/internal/topology"
)

// Engine is the v2 implementation of the Simulator interface.
// It includes aging, death (via a Sigmoid probability curve), and
// strategy reproduction from the original distribution.
type Engine struct {
	graph         *core.Graph
	initialConfig core.SimConfig
}

// NewEngine creates a new V2 evolutionary engine.
func NewEngine() *Engine {
	defaultConfig := core.SimConfig{
		Topology: "ring",
		Size:     12,
		Distribution: map[string]float64{
			"AlwaysCooperator": 0.33,
			"AlwaysCheater":    0.33,
			"Copycat":          0.34,
		},
	}
	return &Engine{
		graph:         topology.GenerateRingGraph(defaultConfig),
		initialConfig: defaultConfig,
	}
}

func (e *Engine) GetState() *core.Graph {
	return e.graph
}

func (e *Engine) Reset(config core.SimConfig) {
	e.initialConfig = config
	if config.Topology == "complete" {
		e.graph = topology.GenerateFullyConnectedGraph(config)
	} else {
		e.graph = topology.GenerateRingGraph(config)
	}
}

func (e *Engine) AdvanceTick() {
	e.graph.Mu.Lock()
	defer e.graph.Mu.Unlock()

	// 1. Process Prisoner's Dilemma interactions
	processInteractions(e.graph)

	// 2. Process Age, Death, and Reproduction
	processLifecycle(e.graph, e.initialConfig)

	e.graph.Tick++
}
