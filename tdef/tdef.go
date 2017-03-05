package tdef

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

type Strategy interface {
	Name() string
	Execute(*Frame, int) string
}

type GameInfo struct {
	Player   int    `json:"Player"`
	Username string `json:"UserName"`
	GameName string `json:"GameName"`
}

type Frame struct {
	W  int `json:"w"`
	H  int `json:"h"`
	P1 struct {
		Owner  int      `json:"owner"`
		Income int      `json:"income"`
		Bits   int      `json:"bits"`
		Towers []string `json:"towers"`
		Troops []struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"troops"`
		MainCore struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"mainCore"`
	} `json:"p1"`
	P2 struct {
		Owner  int      `json:"owner"`
		Income int      `json:"income"`
		Bits   int      `json:"bits"`
		Towers []string `json:"towers"`
		Troops []struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"troops"`
		MainCore struct {
			Owner int `json:"owner"`
			X     int `json:"x"`
			Y     int `json:"y"`
			Maxhp int `json:"maxhp"`
			Hp    int `json:"hp"`
			Enum  int `json:"enum"`
		} `json:"mainCore"`
	} `json:"p2"`
}

type Player struct {
	ServerURL string
	Username  string
	APIKey    string
}

func NewPlayer(credentialFile *string) *Player {
	fp, err := os.Open(*credentialFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			panic(err)
		}
	}()
	buf := make([]byte, 1024)
	n, err := fp.Read(buf)
	if err != nil {
		panic(err)
	}
	if n == 0 {
		return nil
	}
	res := string(buf[:n])
	strres := string(res)
	bits := strings.Split(strres, "\n")
	p := &Player{bits[0], bits[1], bits[2]}
	return p
}

func StartGame(player *Player, strat Strategy) {
	// Open a new web socket
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial("ws://localhost:8080/wsplay", nil)
	handleError(err)
	defer conn.Close()

	// send the api key
	err = conn.WriteMessage(1, []byte(player.APIKey))
	handleError(err)

	// receive acknowledgment
	gameInfo := &GameInfo{}
	_, msg, err := conn.ReadMessage()
	handleError(err)
	json.Unmarshal(msg, gameInfo)
	log.Printf("Match Found. Player: %d, Username: %s, GameName: %s\n", gameInfo.Player,
		gameInfo.Username, gameInfo.GameName)

	//send game inputs
	frame := &Frame{}
	for {
		_, msg, err = conn.ReadMessage()
		_ = msg
		handleError(err)
		json.Unmarshal(msg, frame)
		fmt.Printf("%d, %d\n", frame.P1.MainCore.Hp, frame.P2.MainCore.Hp)
		if frame.P1.MainCore.Hp <= 0 {
			fmt.Println("Player 2 Wins!")
			return
		}
		if frame.P2.MainCore.Hp <= 0 {
			fmt.Println("Player 1 Wins!")
			return
		}
		action := strat.Execute(frame, gameInfo.Player)
		log.Println(action)
		conn.WriteMessage(1, []byte(action))
	}
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
