package strategy

import (
	"math/rand"

	"github.com/mpoegel/tower-defense-ai/tdef"
)

func getAffordableTowers(bits int) []*tdef.Tower {
	var towers []*tdef.Tower
	for _, tower := range tdef.AllTowers {
		if tower.Cost <= bits {
			towers = append(towers, tower)
		}
	}
	return towers
}

func getAffordableTroops(bits int) []*tdef.Troop {
	var troops []*tdef.Troop
	for _, troop := range tdef.AllTroops {
		if troop.Cost <= bits {
			troops = append(troops, troop)
		}
	}
	return troops
}

func choice(r *rand.Rand, arr []int) int {
	n := len(arr)
	i := r.Int() % n
	return arr[i]
}

func choiceDistribution(r *rand.Rand, arr []int, probs []float64) int {
	cumsum := make([]float64, len(probs))
	for i := range probs {
		if i == 0 {
			cumsum[i] = probs[0]
		} else if i < len(probs) {
			cumsum[i] = cumsum[i-1] + probs[i]
		} else {
			cumsum[i] = 1
		}
	}
	p := r.Float64()
	low := 0
	high := len(cumsum) - 1
	mid := low + (high-low)/2
	for {
		if low >= high || (cumsum[low] < p && cumsum[low+1] > p) {
			return arr[low+1]
		}
		if cumsum[mid] > p {
			high = mid
		} else {
			low = mid
		}
		mid = low + (high-low)/2
	}
}

func argmax(arr []float64) int {
	maxI := 0
	for i, a := range arr {
		if a > arr[maxI] {
			maxI = i
		}
	}
	return maxI
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
