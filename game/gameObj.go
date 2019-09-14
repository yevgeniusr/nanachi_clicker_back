package game

import (
	"fmt"

	"github.com/google/uuid"
)

//Game ...
type Game struct {
	ID          uuid.UUID
	Players     *[]Player
	MoneyForWin int
	Ended       bool
	StartedTime string
	EndedTime   string
	Winer       *Player
	Workers     *[]Worker
}

func (g *Game) start() {
	fmt.Print("Game started \n")

	sendStartGame(g)
}

func sendStartGame(g *Game) {
	for _, player := range *g.Players {
		_ = player.Ws.WriteJSON(Message{Type: "game_id", Value: g.ID})
		_ = player.Ws.WriteJSON(Message{Type: "game_player", Value: player})
	}
}
