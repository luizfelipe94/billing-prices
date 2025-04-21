package persistence

import (
	"database/sql"

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
	return err
}

func (r *PostgresPriceRepository) ListPrices() ([]entities.Price, error) {
	query := "SELECT product, measure, size, price FROM prices"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prices []entities.Price
	for rows.Next() {
		var price entities.Price
		if err := rows.Scan(&price.Product, &price.Measure, &price.Size, &price.Price); err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	return prices, nil
}
