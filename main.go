package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	// module "github.com/PifagorRZ/nanachi_hub_back/module"
)

// ServerConfig ...
type ServerConfig struct {
	HTTPport string `json:"HTTPport"`
	TCPport  string `json:"TCPport"`
}

func main() {
	r := GetRouters()

	config := parseConfig(os.Args[1])

	fmt.Printf("Server started on port: %v \n", config.HTTPport)
	defer fmt.Print("Server stoped")

	r.HandleFunc("/ws", HandleConnection)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	go handleMessages()
	http.ListenAndServe(":"+config.HTTPport, r)
}

func parseConfig(path string) *ServerConfig {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error with config path: %v\n", path)
	}

	config := ServerConfig{}

	_ = json.Unmarshal(configFile, &config)
	return &config
}
