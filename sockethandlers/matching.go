package sockethandlers

import (
	"github.com/gorilla/websocket"
	game "github.com/PifagorRZ/nanachi_clicker_back/game"
)

var waiters = []*websocket.Conn{}

// FoundMatchHandler ...
func FoundMatchHandler(s *websocket.Conn) error {
	waiters = append(waiters, s)

	if len(waiters) >= 2 {
		game.MakeGame(waiters[:2])
		waiters = []*websocket.Conn{}
	}

	return nil
}

