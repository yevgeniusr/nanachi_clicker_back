package game

import (
	guuid "github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// MakeGame ...
func MakeGame(wss []*websocket.Conn) {
	game := Game{Players: createPlayers(wss), ID: guuid.New()}
	runGame(game)
}

func createPlayers(wss []*websocket.Conn) *[]Player {
	players := []Player{}

	for _, ws := range wss {
		player := Player{Ws: ws, MaxCPS: 3}
		players = append(players, player)
	}

	return &players
}
