package main

import (
	"encoding/json"
	"first_web_app/app/server"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	ServerAddress string `json:"web_server"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	fmt.Println(config)
	server.RunWebPortal(config.ServerAddress)
}
