package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	g "github.com/PifagorRZ/nanachi_clicker_back/game"
	handlers "github.com/PifagorRZ/nanachi_clicker_back/sockethandlers"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan g.Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleConnection ...
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()
	clients[ws] = true
	handlers.FoundMatchHandler(ws)

	for {
		var msg g.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		msg.User = ws
		// send the new message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {

		// grab next message from the broadcast channel
		msg := <-broadcast

		switch msg.Type {
		case "clicks":
			handlers.ClickHandler(msg)
		}
		// send it out to every client that is currently connected
		// for client := range clients {
		// 	err := client.WriteJSON(msg)
		// 	if err != nil {
		// 		log.Printf("error: %v", err)
		// 		client.Close()
		// 		delete(clients, client)
		// 	}
		// }
	}
}
