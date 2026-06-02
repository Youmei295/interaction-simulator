package v1_static

import (
	"math/rand"
	"time"

	"interaction-simulator/internal/core"
	"interaction-simulator/internal/strategy"
	"interaction-simulator/internal/topology"
)

// Engine is the v1 implementation of the Simulator interface.
// It runs a static graph where nodes interact but do not evolve.
type Engine struct {
	graph *core.Graph
}

func NewEngine() *Engine {
	// Initialize with a default graph
	return &Engine{
		graph: topology.GenerateRingGraph(12),
	}
}

func (e *Engine) GetState() *core.Graph {
	return e.graph
}

func (e *Engine) Reset(topo string, size int) {
	if topo == "complete" {
		e.graph = topology.GenerateFullyConnectedGraph(size)
	} else {
		e.graph = topology.GenerateRingGraph(size)
	}
}

func (e *Engine) AdvanceTick() {
	e.graph.Mu.Lock()
	defer e.graph.Mu.Unlock()

	// Phase 1: Queue Generation
	type pair struct{ A, B string }
	queue := make([]pair, len(e.graph.Edges))
	for i, edge := range e.graph.Edges {
		queue[i] = pair{edge.NodeA, edge.NodeB}
	}

	// Phase 2: Game Execution - Shuffle
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(queue), func(i, j int) { queue[i], queue[j] = queue[j], queue[i] })

	// We calculate all actions first to ensure simultaneous evaluation
	type interaction struct {
		nodeA, nodeB     string
		actionA, actionB core.Action
	}
	interactions := make([]interaction, len(queue))

	for i, p := range queue {
		nodeA := e.graph.Nodes[p.A]
		nodeB := e.graph.Nodes[p.B]

		stratA := strategy.Get(nodeA.Strategy)
		stratB := strategy.Get(nodeB.Strategy)

		actA := stratA.GetAction(nodeA, p.B)
		actB := stratB.GetAction(nodeB, p.A)

		interactions[i] = interaction{
			nodeA:   p.A,
			nodeB:   p.B,
			actionA: actA,
			actionB: actB,
		}
	}

	// Phase 3: State Resolution & Payoff Matrix
	for _, intx := range interactions {
		nodeA := e.graph.Nodes[intx.nodeA]
		nodeB := e.graph.Nodes[intx.nodeB]

		var scoreA, scoreB float64

		if intx.actionA == core.Cooperate && intx.actionB == core.Cooperate {
			scoreA, scoreB = 2, 2
		} else if intx.actionA == core.Defect && intx.actionB == core.Cooperate {
			scoreA, scoreB = 3, -1
		} else if intx.actionA == core.Cooperate && intx.actionB == core.Defect {
			scoreA, scoreB = -1, 3
		} else if intx.actionA == core.Defect && intx.actionB == core.Defect {
			scoreA, scoreB = 0, 0
		}

		stratA := strategy.Get(nodeA.Strategy)
		stratB := strategy.Get(nodeB.Strategy)

		stratA.ApplyOutcome(nodeA, intx.nodeB, intx.actionA, intx.actionB, scoreA)
		stratB.ApplyOutcome(nodeB, intx.nodeA, intx.actionB, intx.actionA, scoreB)
	}

	e.graph.Tick++
}
