package sockethandlers

import (
	game "github.com/PifagorRZ/nanachi_clicker_back/game"
	"github.com/gorilla/websocket"
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
