package sockethandlers

import (
	"github.com/gorilla/websocket"
	game "github.com/yevheniira/nanachi_hub_backend/game"
)

var waiters = []*websocket.Conn{}

// FoundMatchHandler ...
func FoundMatchHandler(s *websocket.Conn) error  {
	waiters = append(waiters, s)

	if len(waiters) >= 1 {
		game.MakeGame(waiters[:1])
		waiters = []*websocket.Conn{}
	}

	return nil
}