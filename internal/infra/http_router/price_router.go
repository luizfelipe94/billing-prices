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
	createPriceHandler *app.CreatePriceHandler
	listPriceHandler   *app.ListPricesHandler
}

func NewPriceRouter(repository repositories.PriceRepository, db *sql.DB, kafkaProducer *infra.KafkaProducer) *PriceRouter {
	return &PriceRouter{
		createPriceHandler: app.NewCreatePriceHandler(repository, kafkaProducer),
		listPriceHandler:   app.NewListPricesHandler(repository),
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

func (h *PriceRouter) ListPrices(w http.ResponseWriter, r *http.Request) {
	prices, err := h.listPriceHandler.Handle(r.Context())
	if err != nil {
		http.Error(w, "Failed to list prices", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)
}
