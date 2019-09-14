package sockethandlers

import (
	g "github.com/PifagorRZ/nanachi_clicker_back/game"
)

//ClickHandler ...
func ClickHandler(value g.Message) {
	g.HandleClick(value)
}