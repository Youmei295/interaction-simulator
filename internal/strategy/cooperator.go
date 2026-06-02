package strategy

import "interaction-simulator/internal/core"

func init() {
	Register("AlwaysCooperator", &AlwaysCooperator{})
}

type AlwaysCooperator struct {
	BaseStrategy
}

func (s *AlwaysCooperator) GetAction(node *core.Node, opponentID string) core.Action {
	return core.Cooperate
}
