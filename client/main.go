package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrr/fastws"
)

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

	var message []byte

	for {
		_, message, err = websocket.ReadMessage(message[:0])

		if err != nil {
			if err != fastws.EOF {
				fmt.Fprintf(os.Stderr, "Error reading message: %s\n", err)
			}

			break
		}

		text := strings.ToLower(string(message))

		if text != "ping" {
			fmt.Printf("Received message: %s\n", text)
		}

		switch text {
		case "lock":
			LockComputer()
		}
	}
}
