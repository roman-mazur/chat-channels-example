package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/roman-mazur/chat-channels-example/pkg/sdk"
)

var (
	target  = flag.String("target", "http://localhost:8080", "Target server address")
	timeout = flag.Duration("timeout", 5*time.Second, "Timeout for scenarios execution")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	client := &sdk.Client{BaseUrl: *target}

	fmt.Println("=== Scenario 1 ===")
	channels, err := client.ListChannels(ctx)
	if err != nil {
		log.Fatal("Cannot list channels: ", err)
	}
	fmt.Print("Available channels: ")
	fmt.Println(channels)

	fmt.Println("=== Scenario 2 ===")
	err = client.CreateChannel(ctx, "my-new-channel")
	if err != nil {
		log.Fatal("Cannot create channel: ", err)
	}
	fmt.Println("Created a new channel")
	channels, err = client.ListChannels(ctx)
	if err != nil {
		log.Fatal("Cannot list channels: ", err)
	}
	fmt.Print("New channels list: ")
	fmt.Println(channels)
}
