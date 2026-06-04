package topology

import (
	"fmt"
	"math/rand"
	"interaction-simulator/internal/core"
	"interaction-simulator/internal/strategy"
)

// RandomStrategy selects a strategy based on the provided weighted distribution map.
// If the map is empty or all weights are 0, it falls back to a uniform random selection from available.
func RandomStrategy(dist map[string]float64) string {
	available := strategy.Available()
	if len(available) == 0 {
		return "AlwaysCooperator" // Failsafe
	}

	// Validate and sum distribution
	totalWeight := 0.0
	for _, weight := range dist {
		if weight > 0 {
			totalWeight += weight
		}
	}

	// If no valid distribution is provided, default to random uniform
	if totalWeight <= 0 {
		return available[rand.Intn(len(available))]
	}

	// Weighted random selection
	r := rand.Float64() * totalWeight
	for strat, weight := range dist {
		if weight > 0 {
			r -= weight
			if r <= 0 {
				return strat
			}
		}
	}

	return available[0]
}

// GenerateRingGraph generates a ring topology graph
func GenerateRingGraph(config core.SimConfig) *core.Graph {
	g := &core.Graph{
		Nodes: make(map[string]*core.Node),
		Edges: []core.Edge{},
	}
	
	// Add nodes
	for i := 0; i < config.Size; i++ {
		id := fmt.Sprintf("node_%02d", i)
		strat := RandomStrategy(config.Distribution)
		g.Nodes[id] = &core.Node{
			ID:       id,
			Strategy: strat,
			Score:    0,
			Age:      0,
			Memory:   make(map[string]core.Action),
			Meta:     make(map[string]map[string]interface{}),
		}
	}

	// Add edges
	for i := 0; i < config.Size; i++ {
		id1 := fmt.Sprintf("node_%02d", i)
		id2 := fmt.Sprintf("node_%02d", (i+1)%config.Size)
		g.Edges = append(g.Edges, core.Edge{NodeA: id1, NodeB: id2})
	}

	return g
}

// GenerateFullyConnectedGraph generates a complete graph
func GenerateFullyConnectedGraph(config core.SimConfig) *core.Graph {
	g := &core.Graph{
		Nodes: make(map[string]*core.Node),
		Edges: []core.Edge{},
	}
	
	// Add nodes
	for i := 0; i < config.Size; i++ {
		id := fmt.Sprintf("node_%02d", i)
		strat := RandomStrategy(config.Distribution)
		g.Nodes[id] = &core.Node{
			ID:       id,
			Strategy: strat,
			Score:    0,
			Age:      0,
			Memory:   make(map[string]core.Action),
			Meta:     make(map[string]map[string]interface{}),
		}
	}

	// Add edges
	for i := 0; i < config.Size; i++ {
		for j := i + 1; j < config.Size; j++ {
			id1 := fmt.Sprintf("node_%02d", i)
			id2 := fmt.Sprintf("node_%02d", j)
			g.Edges = append(g.Edges, core.Edge{NodeA: id1, NodeB: id2})
		}
	}

	return g
}
