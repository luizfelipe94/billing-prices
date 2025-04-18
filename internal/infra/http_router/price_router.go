package http_router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/luizfelipe94/billing-prices/internal/app"
	"github.com/luizfelipe94/billing-prices/internal/domain/repositories"
	"github.com/luizfelipe94/billing-prices/internal/infra"
)

type PriceRouter struct {
	createPriceHandler app.CreatePriceHandler
}

func NewPriceRouter(repository repositories.PriceRepository, db *sql.DB, kafkaProducer *infra.KafkaProducer) *PriceRouter {
	return &PriceRouter{
		createPriceHandler: *app.NewCreatePriceHandler(repository, kafkaProducer),
	}
}

func (h *PriceRouter) CreatePrice(w http.ResponseWriter, r *http.Request) {
	var price app.CreatePriceCommand
	if err := json.NewDecoder(r.Body).Decode(&price); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.createPriceHandler.Handle(r.Context(), price); err != nil {
		http.Error(w, "Failed to create price", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(price)
}
