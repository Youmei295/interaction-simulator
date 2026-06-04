package v2_evo

import (
	"math"
	"math/rand"

	"interaction-simulator/internal/core"
	"interaction-simulator/internal/topology"
)

const (
	LifespanMedian = 50.0 // L: Age at which death probability is exactly 50%
	Steepness      = 0.2  // k: Controls how sharply the death rate rises
)

// processLifecycle increments age and handles the death/reproduction mechanics.
func processLifecycle(graph *core.Graph, config core.SimConfig) {
	for _, node := range graph.Nodes {
		node.Mu.Lock()
		node.Age++

		// Calculate Death Probability using Sigmoid Function
		// P(death) = 1 / (1 + e^(-k * (Age - L)))
		exponent := -Steepness * (float64(node.Age) - LifespanMedian)
		deathProb := 1.0 / (1.0 + math.Exp(exponent))

		// Roll for death
		r := rand.Float64()
		if r < deathProb {
			replaceNode(node, config.Distribution)
		}
		
		node.Mu.Unlock()
	}
}

// replaceNode completely resets a node and gives it a new strategy
// based on the original distribution pool.
func replaceNode(node *core.Node, dist map[string]float64) {
	node.Score = 0
	node.Age = 0
	node.Memory = make(map[string]core.Action)
	node.Meta = make(map[string]map[string]interface{})
	
	// Draw a new strategy from the distribution
	node.Strategy = topology.RandomStrategy(dist)
}
