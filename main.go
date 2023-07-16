package main

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/sid-008/PubSub/broker"
	"github.com/sid-008/PubSub/publish"
	"github.com/sid-008/PubSub/subscribe"
)

func startServer() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

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
