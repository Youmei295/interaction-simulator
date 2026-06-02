package strategy

import "interaction-simulator/internal/core"

func init() {
	Register("Detective", &Detective{})
}

type Detective struct{}

func (s *Detective) GetAction(node *core.Node, opponentID string) core.Action {
	node.Mu.RLock()
	defer node.Mu.RUnlock()

	meta := node.Meta[opponentID]
	if meta == nil {
		return core.Cooperate // First action of probe (C)
	}

	probeRound, _ := meta["probe_round"].(int)
	mode, _ := meta["mode"].(string)

	if mode == "cheat" {
		return core.Defect
	}
	if mode == "copycat" {
		if lastAction, exists := node.Memory[opponentID]; exists {
			return lastAction
		}
		return core.Cooperate
	}

	// Still in probe mode
	probeSequence := []core.Action{core.Cooperate, core.Cooperate, core.Defect, core.Cooperate}
	if probeRound < len(probeSequence) {
		return probeSequence[probeRound]
	}
	return core.Cooperate
}

func (s *Detective) ApplyOutcome(node *core.Node, opponentID string, myAction, oppAction core.Action, scoreDelta float64) {
	node.Mu.Lock()
	defer node.Mu.Unlock()

	node.Score += scoreDelta
	node.Memory[opponentID] = oppAction

	meta := node.Meta[opponentID]
	if meta == nil {
		meta = make(map[string]interface{})
		meta["probe_round"] = 0
		meta["opponent_defected"] = false
		node.Meta[opponentID] = meta
	}

	if oppAction == core.Defect {
		meta["opponent_defected"] = true
	}

	probeRound, _ := meta["probe_round"].(int)
	probeRound++
	meta["probe_round"] = probeRound

	if probeRound >= 4 {
		opponentDefected, _ := meta["opponent_defected"].(bool)
		if !opponentDefected {
			meta["mode"] = "cheat"
		} else {
			meta["mode"] = "copycat"
		}
	}
}
