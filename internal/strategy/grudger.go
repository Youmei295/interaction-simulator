package strategy

import "interaction-simulator/internal/core"

func init() {
	Register("Grudger", &Grudger{})
}

type Grudger struct{}

func (s *Grudger) GetAction(node *core.Node, opponentID string) core.Action {
	node.Mu.RLock()
	defer node.Mu.RUnlock()

	if meta, exists := node.Meta[opponentID]; exists {
		if grudge, ok := meta["grudge"].(bool); ok && grudge {
			return core.Defect
		}
	}
	return core.Cooperate
}

func (s *Grudger) ApplyOutcome(node *core.Node, opponentID string, myAction, oppAction core.Action, scoreDelta float64) {
	node.Mu.Lock()
	defer node.Mu.Unlock()
	
	node.Score += scoreDelta
	node.Memory[opponentID] = oppAction

	if oppAction == core.Defect {
		if node.Meta[opponentID] == nil {
			node.Meta[opponentID] = make(map[string]interface{})
		}
		node.Meta[opponentID]["grudge"] = true
	}
}
