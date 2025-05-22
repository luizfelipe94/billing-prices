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

func (r *PostgresPriceRepository) GetPriceByAttributes(product, measure, size string) (*entities.Price, error) {
	query := "SELECT product, measure, size, price FROM prices WHERE product = $1 AND measure = $2 AND size = $3"
	row := r.DB.QueryRow(query, product, measure, size)

	var price entities.Price
	if err := row.Scan(&price.Product, &price.Measure, &price.Size, &price.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &price, nil
}

func (r *PostgresPriceRepository) UpdatePrice(price entities.Price) error {
	query := "UPDATE prices SET price = $1 WHERE product = $2 AND measure = $3 AND size = $4"
	_, err := r.DB.Exec(query, price.Price, price.Product, price.Measure, price.Size)
	return err
}