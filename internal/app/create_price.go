package app

import (
	"context"
	"encoding/json"

	models "github.com/luizfelipe94/billing-prices/internal/domain/entities"
	"github.com/luizfelipe94/billing-prices/internal/domain/repositories"
	"github.com/luizfelipe94/billing-prices/internal/infra"
)

type CreatePriceCommand struct {
	Product string  `json:"product"`
	Measure string  `json:"measure"`
	Size    string  `json:"size"`
	Price   float64 `json:"price"`
}

type CreatePriceHandler struct {
	PriceRepository repositories.PriceRepository
	KafkaProducer   *infra.KafkaProducer
}

func NewCreatePriceHandler(priceRepository repositories.PriceRepository, kafkaProducer *infra.KafkaProducer) *CreatePriceHandler {
	return &CreatePriceHandler{
		PriceRepository: priceRepository,
		KafkaProducer:   kafkaProducer,
	}
}

func (h *CreatePriceHandler) Handle(ctx context.Context, command CreatePriceCommand) error {
	price := models.Price{
		Product: command.Product,
		Measure: command.Measure,
		Size:    command.Size,
		Price:   command.Price,
	}

	if err := h.PriceRepository.CreatePrice(price); err != nil {
		return err
	}

	event, err := json.Marshal(price)
	if err != nil {
		return err
	}

	if err := h.KafkaProducer.Publish(ctx, []byte(price.GetKey()), event); err != nil {
		return err
	}

	return nil
}
