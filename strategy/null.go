package strategy

import "github.com/mpoegel/tower-defense-ai/tdef"

type NullStrategy struct {
	count int
}

func NewNullStrategy() *NullStrategy {
	s := NullStrategy{0}
	return &s
}

func (strat *NullStrategy) Name() string {
	return "Null Strategy"
}

func (strat *NullStrategy) Init(player int) {
	return
}

func (strat *NullStrategy) Execute(frame *tdef.Frame, player int) string {
	action := ""
	return action
}

func (strat *NullStrategy) Save() {
	return
}
