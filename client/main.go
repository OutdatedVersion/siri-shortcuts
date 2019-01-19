package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dgrr/fastws"
)

type Message struct {
	Action  string `json:"action"`
	Payload string `json:"payload"`
}

func main() {
	eventHubURI := os.Getenv("EVENT_HUB_URI")

	if eventHubURI == "" {
		fmt.Println("Missing EVENT_HUB_URI environment variable")
		return
	}

	fmt.Printf("Using event hub URI: %s\n", eventHubURI)

	websocket, err := fastws.Dial(eventHubURI)

	if err != nil {
		fmt.Printf("Failed to open WebSocket: %s", err)
		return
	}

	fmt.Println("Opened socket to event hub")

	var frame []byte

	for {
		_, frame, err = websocket.ReadMessage(frame[:0])

		if err != nil {
			if err != fastws.EOF {
				fmt.Fprintf(os.Stderr, "Error reading message: %s\n", err)
			}

			break
		}

		message := Message{}
		var payload map[string]interface{}

		json.Unmarshal(frame, &message)

		if message.Payload != "" {
			json.Unmarshal([]byte(message.Payload), &payload)
		}

		message.Action = strings.ToLower(message.Action)

		if message.Action != "" {
			fmt.Printf("Received message: %s\n", message.Action)
		}

		switch message.Action {
		case "lock":
			LockComputer()

		case "shutdown":
			ShutdownComputer()
		}
	}
}
