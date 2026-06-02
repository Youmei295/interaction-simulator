package topology

import (
	"fmt"
	"interaction-simulator/internal/core"
	"interaction-simulator/internal/strategy"
)

// GenerateRingGraph generates a ring topology graph
func GenerateRingGraph(size int) *core.Graph {
	g := &core.Graph{
		Nodes: make(map[string]*core.Node),
		Edges: []core.Edge{},
	}
	
	strategies := strategy.Available()
	if len(strategies) == 0 {
		strategies = []string{"AlwaysCooperator"} // Fallback
	}
	
	// Add nodes
	for i := 0; i < size; i++ {
		id := fmt.Sprintf("node_%02d", i)
		strat := strategies[i%len(strategies)]
		g.Nodes[id] = &core.Node{
			ID:       id,
			Strategy: strat,
			Score:    0,
			Memory:   make(map[string]core.Action),
			Meta:     make(map[string]map[string]interface{}),
		}
	}

	// Add edges
	for i := 0; i < size; i++ {
		id1 := fmt.Sprintf("node_%02d", i)
		id2 := fmt.Sprintf("node_%02d", (i+1)%size)
		g.Edges = append(g.Edges, core.Edge{NodeA: id1, NodeB: id2})
	}

	return g
}

// GenerateFullyConnectedGraph generates a complete graph
func GenerateFullyConnectedGraph(size int) *core.Graph {
	g := &core.Graph{
		Nodes: make(map[string]*core.Node),
		Edges: []core.Edge{},
	}
	
	strategies := strategy.Available()
	if len(strategies) == 0 {
		strategies = []string{"AlwaysCooperator"} // Fallback
	}
	
	// Add nodes
	for i := 0; i < size; i++ {
		id := fmt.Sprintf("node_%02d", i)
		strat := strategies[i%len(strategies)]
		g.Nodes[id] = &core.Node{
			ID:       id,
			Strategy: strat,
			Score:    0,
			Memory:   make(map[string]core.Action),
			Meta:     make(map[string]map[string]interface{}),
		}
	}

	// Add edges
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			id1 := fmt.Sprintf("node_%02d", i)
			id2 := fmt.Sprintf("node_%02d", j)
			g.Edges = append(g.Edges, core.Edge{NodeA: id1, NodeB: id2})
		}
	}

	return g
}
