package game

//CPS - Clicks per second
//MPS - Modey per second

import (
	"github.com/gorilla/websocket"
	skilltrees "github.com/yevheniira/nanachi_hub_backend/skilltrees"
)

// Nationalities
const (
	INDUS = "indus"
	POLAND = "poland"
)

//Player ...
type Player struct {
	Ws *websocket.Conn
	Name string
	Money int
	Workers []Worker
	FrontendSkills []skilltrees.Skill
	BackendSkills []skilltrees.Skill
	MaxCPS int
	App Application
}

//Worker ...
type Worker struct {
	Name string
	Nationality string
	Level int
	Price int
	CPS int
}


//Application ...
type Application struct {
	FrontendCode int
	BackendCode int
	Reliability float32
	Deployed bool
	MPS int
}

//Message ...
type Message struct {
	Type string `json:"type"`
	Value  interface{} `json:"value"`
}
