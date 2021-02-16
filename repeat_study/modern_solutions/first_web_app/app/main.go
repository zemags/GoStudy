package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/zemags/GoStudy/repeat_study/modern_solutions/first_web_app/app/server"
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
