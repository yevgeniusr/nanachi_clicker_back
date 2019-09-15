package game

var games = []Game{}

//RunGame ...
func RunGame(game Game) {
	game.start()
	go game.startIntervals()
	games = append(games, game)
}

//HandleClick ...
func HandleClick(msg Message) {
	for _, game := range games {
		if game.UserHere(msg.User) {
			// game.handleClick(msg.User, msg.Value)
			game.SendForPlayers(&Message{Type: "game_click", Value: "click"}, msg.User)
		}
	}
}