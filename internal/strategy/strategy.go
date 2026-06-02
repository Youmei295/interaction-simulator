package strategy

import (
	"interaction-simulator/internal/core"
)

// Strategy defines the interface for all game behaviors.
type Strategy interface {
	// GetAction determines the next move against an opponent.
	GetAction(node *core.Node, opponentID string) core.Action
	
	// ApplyOutcome allows the strategy to update its state based on the result.
	ApplyOutcome(node *core.Node, opponentID string, myAction, oppAction core.Action, scoreDelta float64)
}

var registry = make(map[string]Strategy)

// Register adds a strategy to the global registry.
func Register(name string, s Strategy) {
	registry[name] = s
}

// Get retrieves a strategy by name.
func Get(name string) Strategy {
	if s, exists := registry[name]; exists {
		return s
	}
	// Fallback strategy if not found
	return registry["AlwaysCooperator"]
}

// Available returns a list of all registered strategy names.
func Available() []string {
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	return names
}

// BaseStrategy provides a default implementation for ApplyOutcome
// to reduce boilerplate for simple strategies.
type BaseStrategy struct{}

func (b *BaseStrategy) ApplyOutcome(node *core.Node, opponentID string, myAction, oppAction core.Action, scoreDelta float64) {
	node.Mu.Lock()
	defer node.Mu.Unlock()
	node.Score += scoreDelta
	node.Memory[opponentID] = oppAction
}
