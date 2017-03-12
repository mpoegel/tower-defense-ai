package strategy

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestChoiceDistribution(t *testing.T) {
	source := rand.NewSource(1)
	rr := rand.New(source)
	x := ChoiceDistribution(rr, []int{1, 2, 3, 4, 5}, []float64{0.2, 0.2, 0.2, 0.2, 0.2})
	fmt.Println(x)
}
