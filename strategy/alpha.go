package strategy

import (
	"fmt"
	"math/rand"

	"github.com/mpoegel/tower-defense-ai/tdef"
)

type AlphaStrategy struct {
	count           int
	weights         []float64
	usedGhandi      bool
	player          int
	lastAction      string
	numberOfActions int
	actionIDs       []int
	r               *rand.Rand
}

func NewAlphaStrategy(seed int64) *AlphaStrategy {
	source := rand.NewSource(seed)
	r := rand.New(source)
	numberOfActions := (len(tdef.AllTowers)+len(tdef.AllTroops))*3 + 1
	actionIDs := makeRange(0, numberOfActions)
	weights := make([]float64, numberOfActions)
	for i := range weights {
		weights[i] = float64(1 / numberOfActions)
	}
	s := AlphaStrategy{0, weights, false, -1, "", numberOfActions, actionIDs, r}
	return &s
}

func (strat *AlphaStrategy) Init(player int) {
	strat.player = player
}

func (strat *AlphaStrategy) Name() string {
	return "Alpha Strategy"
}

func (strat *AlphaStrategy) Execute(frame *tdef.Frame, player int) string {
	strat.count++
	i := choiceDistribution(strat.r, strat.actionIDs, strat.weights)
	action := strat.intToAction(i)
	strat.lastAction = action
	return action
}

func (strat *AlphaStrategy) intToAction(i int) string {
	if i == strat.numberOfActions {
		return ""
	}
	unit := i % (strat.numberOfActions - 1)
	lane := i / (strat.numberOfActions - 1)
	if unit < len(tdef.AllTowers) {
		unit = tdef.AllTowers[unit].Enum
	} else {
		unit = tdef.AllTroops[unit-len(tdef.AllTowers)].Enum
	}
	var action string
	if unit < 10 {
		action = fmt.Sprintf("b0%d 0%d", unit, lane)
	} else {
		action = fmt.Sprintf("b%d 0%d", unit, lane)
	}
	return action
}

func (strat *AlphaStrategy) getPlayerData(frame *tdef.Frame) tdef.PlayerData {
	if strat.player == 1 {
		return frame.P1
	}
	return frame.P2
}

func (strat *AlphaStrategy) getOpponentData(frame *tdef.Frame) tdef.PlayerData {
	if strat.player == 1 {
		return frame.P2
	}
	return frame.P1
}

func (strat *AlphaStrategy) utility(frame *tdef.Frame) float64 {
	pdata := strat.getPlayerData(frame)
	val := float64(pdata.MainCore.Hp)
	return val
}
