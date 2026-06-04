package v2_evo

import (
	"math/rand"
	"time"

	"interaction-simulator/internal/core"
	"interaction-simulator/internal/strategy"
)

// processInteractions handles Phase 1 and 2: Queue generation, strategy evaluation,
// and payoff matrix calculation.
func processInteractions(graph *core.Graph) {
	// Phase 1: Queue Generation
	type pair struct{ A, B string }
	queue := make([]pair, len(graph.Edges))
	for i, edge := range graph.Edges {
		queue[i] = pair{edge.NodeA, edge.NodeB}
	}

	// Phase 2: Game Execution - Shuffle
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(queue), func(i, j int) { queue[i], queue[j] = queue[j], queue[i] })

	// Calculate all actions first for simultaneous evaluation
	type interaction struct {
		nodeA, nodeB     string
		actionA, actionB core.Action
	}
	interactions := make([]interaction, len(queue))

	for i, p := range queue {
		nodeA := graph.Nodes[p.A]
		nodeB := graph.Nodes[p.B]

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
		nodeA := graph.Nodes[intx.nodeA]
		nodeB := graph.Nodes[intx.nodeB]

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
}
