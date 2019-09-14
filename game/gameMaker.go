package game

import (
	guuid "github.com/google/uuid"
	"github.com/gorilla/websocket"
	skilltrees "github.com/PifagorRZ/nanachi_clicker_back/skilltrees"
)

// MakeGame ...
func MakeGame(wss []*websocket.Conn) {
	game := Game{Players: createPlayers(wss), ID: guuid.New()}
	RunGame(game)
}

func createPlayers(wss []*websocket.Conn) *[]Player {
	players := []Player{}

	for _, ws := range wss {
		player := Player{Ws: ws, MaxCPS: 3, ID: guuid.New(), 
			FrontendSkills: skilltrees.ParseSkills("frontend"), 
			BackendSkills: skilltrees.ParseSkills("backend")}
		players = append(players, player)
	}

	return &players 
}
