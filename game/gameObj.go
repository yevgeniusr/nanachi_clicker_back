package game

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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

	g.SendForPlayers(&Message{Type: "game_tick", Value: g}, nil)
	g.sendPlayers()
}

//UserHere ...
func (g *Game) UserHere(user *websocket.Conn) bool {
	for _, player := range *g.Players {
		if player.Ws == user {
			return true
		}
	}

	return false
}

//SendForPlayers ...
func (g *Game) SendForPlayers(m *Message, user *websocket.Conn) {
	for _, player := range *g.Players {
		if player.Ws != user {
			_ = player.Ws.WriteJSON(m)
		}
	}
}

//SendForPlayer ...
func (g *Game) SendForPlayer(m *Message, user *websocket.Conn) {
		_ = user.WriteJSON(m)
}

func (g *Game) sendPlayers() {
	for _, player := range *g.Players {
		g.SendForPlayer(&Message{Type: "game_player", Value: player}, player.Ws)

		for _, player2 := range *g.Players {
			if player2.Ws != player.Ws {
				g.SendForPlayer(&Message{Type: "game_enemy", Value: player2}, player.Ws)
			}
		}
	}
}

func (g * Game) handleClick(*websocket.Conn, interface {}) {

}

func (g *Game) startIntervals() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				g.SendForPlayers(&Message{Type: "game_tick", Value: g}, nil)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
