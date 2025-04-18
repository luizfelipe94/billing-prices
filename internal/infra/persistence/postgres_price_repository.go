package persistence

import (
	"database/sql"
	"fmt"

	"github.com/luizfelipe94/billing-prices/internal/domain/entities"
)

type PostgresPriceRepository struct {
	DB *sql.DB
}

func NewPostgresPriceRepository(db *sql.DB) *PostgresPriceRepository {
	return &PostgresPriceRepository{DB: db}
}

func (r *PostgresPriceRepository) CreatePrice(price entities.Price) error {
	query := "INSERT INTO prices (product, measure, size, price) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, price.Product, price.Measure, price.Size, price.Price)
	if err != nil {
		return fmt.Errorf("failed to create price: %w", err)
	}
	return nil
}
