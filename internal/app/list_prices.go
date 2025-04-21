package app

import (
	"context"

	models "github.com/luizfelipe94/billing-prices/internal/domain/entities"
	"github.com/luizfelipe94/billing-prices/internal/domain/repositories"
)

type ListPricesHandler struct {
	PriceRepository repositories.PriceRepository
}

func NewListPricesHandler(priceRepository repositories.PriceRepository) *ListPricesHandler {
	return &ListPricesHandler{
		PriceRepository: priceRepository,
	}
}

func (h *ListPricesHandler) Handle(ctx context.Context) ([]models.Price, error) {
	prices, err := h.PriceRepository.ListPrices()
	if err != nil {
		return nil, err
	}
	return prices, nil
}
