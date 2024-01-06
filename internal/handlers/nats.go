package handlers

import (
	"log"

	"github.com/nats-io/nats.go/jetstream"
)

func (h *Handler) GetNatsOrder(msg jetstream.Msg) {
	order := msg.Data()

	if !h.nats.CompareJsonToSchema(string(order)) {
		log.Println("[JetStream]: invalid json data")
		msg.Ack()
		return
	}

	id, err := h.services.Create(order)
	if err != nil {
		log.Printf("[JetStream]: failed to create order: %v", err)
		msg.Nak()
		return
	}

	if err := h.services.Set(id, order); err != nil {
		log.Printf("[JetStream]: failed to set order data: %v", err)
	}

	msg.Ack()
}
