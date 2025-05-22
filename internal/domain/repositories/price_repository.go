package repositories

import (
	"github.com/luizfelipe94/billing-prices/internal/domain/entities"
)

type PriceRepository interface {
	CreatePrice(price entities.Price) error
	ListPrices() ([]entities.Price, error)
	GetPriceByAttributes(product, measure, size string) (*entities.Price, error)
	UpdatePrice(price entities.Price) error
}
