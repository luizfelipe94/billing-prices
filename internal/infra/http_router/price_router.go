package http_router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/luizfelipe94/billing-prices/internal/app"
	"github.com/luizfelipe94/billing-prices/internal/domain/repositories"
	"github.com/luizfelipe94/billing-prices/internal/infra"
	"k8s.io/client-go/kubernetes"
)

type PriceRouter struct {
	createPriceHandler *app.CreatePriceHandler
	listPriceHandler   *app.ListPricesHandler
	turnOnGenerateDataHandler *app.TurnOnGenerateDataHandler
}

func NewPriceRouter(repository repositories.PriceRepository, db *sql.DB, kafkaProducer *infra.KafkaProducer, k8sClient *kubernetes.Clientset) *PriceRouter {
	return &PriceRouter{
		createPriceHandler: app.NewCreatePriceHandler(repository, kafkaProducer),
		listPriceHandler:   app.NewListPricesHandler(repository),
		turnOnGenerateDataHandler: app.NewTurnOnGenerateDataHandler(k8sClient),
	}
}

func (h *PriceRouter) CreatePrice(w http.ResponseWriter, r *http.Request) {
	var price app.CreatePriceCommand
	if err := json.NewDecoder(r.Body).Decode(&price); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.createPriceHandler.Handle(r.Context(), price); err != nil {
		if err.Error() == "price for product already exists" {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
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

func (h *PriceRouter) TurnOnGenerateData(w http.ResponseWriter, r *http.Request) {
	var command app.GenerateDataCommand
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.turnOnGenerateDataHandler.Handle(r.Context(), command); err != nil {
		http.Error(w, "Failed to turn on data generation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "data generation started"})
}