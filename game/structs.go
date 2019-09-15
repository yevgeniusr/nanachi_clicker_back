package game

//CPS - Clicks per second
//MPS - Modey per second

import (
	skilltrees "github.com/PifagorRZ/nanachi_clicker_back/skilltrees"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Nationalities
const (
	INDUS  = "indus"
	POLAND = "poland"
)

//Player ...
type Player struct {
	ID             uuid.UUID
	Ws             *websocket.Conn
	Name           string
	Money          int
	Workers        []Worker
	FrontendSkills *[]skilltrees.Skill
	BackendSkills  *[]skilltrees.Skill
	MaxCPS         int
	App            Application
}

//Worker ...
type Worker struct {
	Name        string
	Nationality string
	Level       int
	Price       int
	CPS         int
}

//Application ...
type Application struct {
	FrontendCode       int
	BackendCode        int
	Reliability        float32
	Deployed           bool
	DeployCost         int
	DeployRequirements Click
	MPS                int
}

//Message ...
type Message struct {
	User  *websocket.Conn
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

//Click ...
type Click struct {
	FrontendClicks int
	BackendClicks  int
}
