package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) handleGetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		order, err := h.services.Get(id)
		if err != nil {
			log.Printf("[Handler - Redis]: %v", err)

			order, err = h.services.GetByID(id)
			if err != nil {
				log.Printf("[Handler - Postgres]: %v", err)
				http.Error(w, "Failed to retrieve order", http.StatusBadRequest)
				return
			}

			if err := h.services.Set(id, order); err != nil {
				log.Printf("[Handler - Redis]: %v", err)
			}
		}

		w.Write(order)
	}
}
