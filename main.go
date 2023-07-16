package main

import (
	"fmt"
	"os"

	"github.com/sid-008/PubSub/broker"
	"github.com/sid-008/PubSub/publish"
	"github.com/sid-008/PubSub/subscribe"
)

func main() {
	nodetype := os.Args[1]
	switch nodetype {

	case "pub":
		publish.StartPubNode()

	case "broker":
		broker.StartBroker()

	case "sub":
		subscribe.StartPubNode()

	default:
		fmt.Println("Invalid nodetype")
	}
}
