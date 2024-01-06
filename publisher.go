package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	NatsUrl  = "nats://localhost:4222"
	FilePath = "model.json"
)

func main() {
	jsonFile, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	nc, err := nats.Connect(NatsUrl)
	if err != nil {
		log.Fatal("failed to connect to NATS: ", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal("failed to create JetStream context: ", err)
	}

	if _, err := js.Publish(ctx, "ORDER.new", jsonFile); err != nil {
		log.Printf("failed to publish message: %v", err)
	}

}
