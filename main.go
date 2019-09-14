package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	// module "github.com/yevheniira/nanachi_hub_backend/module"
)

// ServerConfig ...
type ServerConfig struct {
	Port string `json:"port"`
}

func main() {
	r := GetRouters()

	config := parseConfig(os.Args[1])

	fmt.Printf("Server started on port: %v \n", config.Port)
	defer fmt.Print("Server stoped")

	http.ListenAndServe(":" + config.Port, r)
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