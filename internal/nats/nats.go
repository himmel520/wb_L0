package nats

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/himmel520/wb_L0/internal/config"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/xeipuuv/gojsonschema"
)

type Nats struct {
	Consumer jetstream.Consumer
	Schema   gojsonschema.JSONLoader
}

func NewNats(cfg *config.Nats) (*Nats, error) {
	consumer, err := NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	loader := gojsonschema.NewReferenceLoader("file:///api/schema.json")

	return &Nats{Consumer: consumer, Schema: loader}, nil
}

func (n *Nats) CompareJsonToSchema(jsonData string) bool {
	loader := gojsonschema.NewStringLoader(jsonData)

	result, err := gojsonschema.Validate(n.Schema, loader)
	if err != nil {
		log.Println("[JsonSchema]: ", err)
		return false
	}

	return result.Valid()
}

func NewConsumer(cfg *config.Nats) (jetstream.Consumer, error) {
	nc, err := nats.Connect(cfg.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %v", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("failed to create JetStream context: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDER",
		Subjects: []string{"ORDER.*"},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create JetStream stream: %v", err)
	}

	c, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create JetStream consumer: %v", err)
	}

	return c, nil
}
