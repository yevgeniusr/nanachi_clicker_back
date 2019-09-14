package game

var games = []Game{}

func runGame(game Game) {
	game.start()
	games = append(games, game)
}