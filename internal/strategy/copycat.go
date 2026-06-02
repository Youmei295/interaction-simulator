package strategy

import "interaction-simulator/internal/core"

func init() {
	Register("Copycat", &Copycat{})
}

type Copycat struct {
	BaseStrategy
}

func (s *Copycat) GetAction(node *core.Node, opponentID string) core.Action {
	node.Mu.RLock()
	defer node.Mu.RUnlock()

	if lastAction, exists := node.Memory[opponentID]; exists {
		return lastAction
	}
	return core.Cooperate // First interaction
}
