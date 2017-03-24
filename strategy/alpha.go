package strategy

import (
	"fmt"
	"math/rand"
	"strconv"

	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/mpoegel/tower-defense-ai/tdef"
)

const alphaStrategySaveFileName = "models/alpha"

// AlphaStrategy is a strategy that uses q-learning to learn the best actions
type AlphaStrategy struct {
	count           int
	weights         []float64
	usedGhandi      bool
	player          int
	lastState       *tdef.Frame
	lastAction      int
	numberOfActions int
	actionIDs       []int
	r               *rand.Rand
	discount        float64
	epsilon         float64
	learningRate    float64
}

// NewAlphaStrategy initializes and returns a new AlphaStrategy given random seed
func NewAlphaStrategy(seed int64, discount, epsilon, learningRate float64) *AlphaStrategy {
	source := rand.NewSource(seed)
	r := rand.New(source)
	numberOfActions := (len(tdef.AllTowers)+len(tdef.AllTroops))*3 + 1
	actionIDs := makeRange(0, numberOfActions)
	weights := make([]float64, numberOfActions)
	if _, err := os.Stat(alphaStrategySaveFileName); os.IsNotExist(err) || !os.IsExist(err) {
		for i := range weights {
			weights[i] = float64(1 / numberOfActions)
		}
		fp, err := os.Create(alphaStrategySaveFileName)
		fp.Close()
		if err != nil {
			log.Println("Could create model file")
		}
	} else {
		b, _ := ioutil.ReadFile(alphaStrategySaveFileName)
		data := strings.Split(string(b), "\n")
		for i, d := range data {
			w, err := strconv.ParseFloat(d, 64)
			if err != nil {
				log.Printf("Error reading model file: '%s'\n", d)
				weights[i] = float64(1 / numberOfActions)
			} else {
				weights[i] = w
			}
		}
	}
	s := AlphaStrategy{0, weights, false, -1, nil, -1, numberOfActions, actionIDs, r, discount,
		epsilon, learningRate}
	return &s
}

// Init updates the player ID, which should be called after a game is found but before Execute is
// called
func (strat *AlphaStrategy) Init(player int) {
	strat.player = player
}

// Name returns the name of the AlphaStrategy as a string
func (strat *AlphaStrategy) Name() string {
	return "Alpha Strategy"
}

// Execute determines the next action to take given the current frame
func (strat *AlphaStrategy) Execute(frame *tdef.Frame, player int) string {
	strat.count++
	a := strat.getAction(frame)
	action := strat.intToAction(a)
	if strat.lastState != nil {
		reward := strat.reward(strat.lastState, strat.lastAction, frame)
		strat.update(strat.lastState, strat.lastAction, frame, reward)
	}
	strat.lastState = frame
	strat.lastAction = a
	return action
}

// Save writes the current model to disk
func (strat *AlphaStrategy) Save() {
	fp, err := os.Open(alphaStrategySaveFileName)
	if err != nil {
		log.Printf("Unable to save model to '%s'\n", alphaStrategySaveFileName)
		return
	}
	defer fp.Close()
	for _, w := range strat.weights {
		_, err = fp.WriteString(strconv.FormatFloat(w, 'f', -1, 64))
		if err != nil {
			log.Printf("Failed to save weight: %f\n", w)
			return
		}
	}
}

// intToAction returns the string action that corresponds to the given integer
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

// getPlayerData returns the player's struct
func (strat *AlphaStrategy) getPlayerData(frame *tdef.Frame) tdef.PlayerData {
	if strat.player == 1 {
		return frame.P1
	}
	return frame.P2
}

// getOpponentData returns the opponent's struct
func (strat *AlphaStrategy) getOpponentData(frame *tdef.Frame) tdef.PlayerData {
	if strat.player == 1 {
		return frame.P2
	}
	return frame.P1
}

func (strat *AlphaStrategy) features(state *tdef.Frame, action int) []int {
	features := make([]int, strat.numberOfActions)
	for i := range features {
		features[i] = 0
	}
	troops := strat.getPlayerData(state).Troops
	for _, troop := range troops {
		e := troop.Enum
		if e < 0 {
			features[strat.numberOfActions-1]++
		} else {
			features[e]++
		}
	}
	towers := strat.getPlayerData(state).Towers
	offset := len(tdef.AllTroops)
	for _, tower := range towers {
		e := tower.Enum
		if e < 0 {
			features[strat.numberOfActions-1]++
		} else {
			features[offset+e]++
		}
	}
	features[action]++
	return features
}

func (strat *AlphaStrategy) getQValue(state *tdef.Frame, action int) float64 {
	features := strat.features(state, action)
	val := 0.0
	for i := range features {
		val += float64(features[i]) * strat.weights[i]
	}
	return val
}

func (strat *AlphaStrategy) computeValueFromQValues(state *tdef.Frame) float64 {
	income := strat.getPlayerData(state).Income
	actions := getAffordableActions(income)
	if len(actions) == 0 {
		return 0.0
	}
	maxQ := strat.getQValue(state, actions[0])
	for i, action := range actions {
		if i == 0 {
			continue
		}
		q := strat.getQValue(state, action)
		if q > maxQ {
			maxQ = q
		}
	}
	return maxQ
}

func (strat *AlphaStrategy) computeActionFromQValues(state *tdef.Frame) int {
	income := strat.getPlayerData(state).Income
	actions := getAffordableActions(income)
	if len(actions) == 0 {
		return -1
	}
	maxQ := strat.getQValue(state, actions[0])
	maxA := actions[0]
	for i, action := range actions {
		if i == 0 {
			continue
		}
		q := strat.getQValue(state, action)
		if q > maxQ {
			maxQ = q
			maxA = action
		}
	}
	return maxA
}

func (strat *AlphaStrategy) getAction(state *tdef.Frame) int {
	income := strat.getPlayerData(state).Income
	actions := getAffordableActions(income)
	if len(actions) == 0 {
		return -1
	}
	p := strat.r.Float64()
	if p < strat.epsilon {
		return choice(strat.r, actions)
	}
	return strat.computeActionFromQValues(state)
}

func (strat *AlphaStrategy) update(state *tdef.Frame, action int, nextState *tdef.Frame,
	reward float64) {
	currVal := strat.getQValue(state, action)
	difference := reward + strat.discount*strat.computeValueFromQValues(nextState) - currVal
	features := strat.features(state, action)
	for i := range strat.weights {
		strat.weights[i] += strat.learningRate * difference * float64(features[i])
	}
}

func (strat *AlphaStrategy) getPolicy(state *tdef.Frame) int {
	return strat.computeActionFromQValues(state)
}

func (strat *AlphaStrategy) getValue(state *tdef.Frame) float64 {
	return strat.computeValueFromQValues(state)
}

func (strat *AlphaStrategy) reward(state *tdef.Frame, action int, nextState *tdef.Frame) float64 {
	pData := strat.getPlayerData(state)
	opData := strat.getOpponentData(state)
	enemyTroops := 0
	enemyTowers := 0
	for _, troop := range opData.Troops {
		if troop.Owner != strat.player {
			enemyTroops++
		}
	}
	for _, tower := range opData.Towers {
		if tower.Owner != strat.player {
			enemyTowers++
		}
	}
	val := float64(pData.MainCore.Hp) +
		float64(pData.Income) +
		-1*float64(opData.MainCore.Hp) +
		-1*float64(enemyTroops) +
		-1*float64(enemyTowers)
	return val
}
