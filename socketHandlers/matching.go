package sockethandlers

import (
	"github.com/gorilla/websocket"
	game "github.com/PifagorRZ/nanachi_clicker_back/game"
)

var waiters = []*websocket.Conn{}

// FoundMatchHandler ...
func FoundMatchHandler(s *websocket.Conn) error {
	waiters = append(waiters, s)

	if len(waiters) >= 1 {
		game.MakeGame(waiters[:1])
		waiters = []*websocket.Conn{}
	}

	return nil
}
