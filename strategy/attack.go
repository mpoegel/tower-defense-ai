package strategy

import (
	"fmt"
	"math/rand"

	"github.com/mpoegel/tower-defense-ai/tdef"
)

type AttackStrategy struct {
	lane  int
	count int
}

func NewAttackStrategy(lane int) *AttackStrategy {
	s := AttackStrategy{lane, 0}
	return &s
}

func (strat *AttackStrategy) Name() string {
	return "Attack Strategy"
}

func (strat *AttackStrategy) Execute(frame *tdef.Frame, player int) string {
	var lane int
	if strat.lane == -1 {
		lane = rand.Int() % 3
	} else {
		lane = strat.lane
	}
	action := fmt.Sprintf("b01 0%d", lane)
	return action
}
