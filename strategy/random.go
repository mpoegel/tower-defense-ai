package strategy

import (
	"math/rand"

	"fmt"

	"github.com/mpoegel/tower-defense-ai/tdef"
)

type RandomStrategy struct {
	count int64
	rr    *rand.Rand
}

func NewRandomStrategy(seed int64) *RandomStrategy {
	source := rand.NewSource(seed)
	rr := rand.New(source)
	s := RandomStrategy{0, rr}
	return &s
}

func (strat *RandomStrategy) Name() string {
	return "Random Strategy"
}

func (strat *RandomStrategy) Execute(frame *tdef.Frame, player int) string {
	// choose an action at random
	// one of every 5 actions will be a random tower placement
	var playerData tdef.PlayerData
	if player == 1 {
		playerData = frame.P1
	} else {
		playerData = frame.P2
	}
	var action string
	if playerData.Bits < 2000 {
		action = ""
	} else if strat.count%5 == 0 {
		// choose a random tower
		enum := (strat.rr.Int() % 10) + 50
		var plot int
		if player == 1 {
			plot = Choice(strat.rr, []int{0, 22, 44}) + strat.rr.Int()%10
		} else {
			plot = Choice(strat.rr, []int{11, 33, 56}) + strat.rr.Int()%10
		}
		if plot < 10 {
			action = fmt.Sprintf("b%d 0%d", enum, plot)
		} else {
			action = fmt.Sprintf("b%d %d", enum, plot)
		}
	} else {
		// choose a random troop
		lane := strat.rr.Int() % 3
		enum := strat.rr.Int() % 11
		var enumStr string
		if enum < 10 {
			enumStr = fmt.Sprintf("0%d", enum)
		} else {
			enumStr = fmt.Sprintf("%d", enum)
		}
		action = fmt.Sprintf("b0%d %s", lane, enumStr)
	}
	strat.count++
	return action
}

func Choice(r *rand.Rand, arr []int) int {
	n := len(arr)
	i := r.Int() % n
	return arr[i]
}
