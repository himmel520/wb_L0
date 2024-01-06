package handlers

import (
	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/himmel520/wb_L0/internal/nats"
	"github.com/himmel520/wb_L0/internal/services"
	"github.com/nats-io/nats.go/jetstream"
)

type Handler struct {
	services *services.Service
	nats     *nats.Nats
}

func NewHandler(services *services.Service, nats *nats.Nats) *Handler {
	return &Handler{
		services: services,
		nats:     nats,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
	))
	r.HandleFunc("/order/{id:[0-9]+}", h.handleGetByID()).Methods("GET")
	return r
}

func (h *Handler) InitNatsConsume() jetstream.ConsumeContext {
	cc, _ := h.nats.Consumer.Consume(
		h.GetNatsOrder,
		jetstream.ConsumeErrHandler(func(consumeCtx jetstream.ConsumeContext, err error) {
			log.Println(err)
		}))
	return cc
}
