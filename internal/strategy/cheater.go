package strategy

import "interaction-simulator/internal/core"

func init() {
	Register("AlwaysCheater", &AlwaysCheater{})
}

type AlwaysCheater struct {
	BaseStrategy
}

func (s *AlwaysCheater) GetAction(node *core.Node, opponentID string) core.Action {
	return core.Defect
}
