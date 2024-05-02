package main

import (
	"encoding/json"
	"fmt"
	"github.com/hugolgst/rich-go/client"
	"os"
	"time"
)

type Config struct {
	ClientID  string `json:"ClientID"`
	ButtonURL string `json:"ButtonURL"`
}

var config Config

func LoadConfig() {
	file, err := os.Open("config.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)

	if err != nil {
		panic(err)
	}
}

func main() {
	LoadConfig()
	err := client.Login(config.ClientID)
	if err != nil {
		panic(err)
	}

	_t := time.Now()

	button1 := &client.Button{
		Label: "Smash!",
		Url:   config.ButtonURL,
	}

	err = client.SetActivity(client.Activity{
		State:      "State: Single",
		Details:    "Do you want to be my Valentine's date",
		LargeImage: "heart",
		LargeText:  "Love is in the air",
		SmallImage: "",
		SmallText:  "",
		Party: &client.Party{
			ID:         "-1",
			Players:    1,
			MaxPlayers: 2,
		},
		Timestamps: &client.Timestamps{
			Start: &_t,
		},
		Buttons: []*client.Button{button1},
	})

	if err != nil {
		panic(err)
	}

	fmt.Scanln()
}
