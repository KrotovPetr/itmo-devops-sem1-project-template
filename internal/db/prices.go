package db

import (
	"fmt"
	"project_sem/internal/constants"
	"time"
)

type Price struct {
	ID         int
	Name       string
	Category   string
	Price      float64
	CreateDate time.Time
}

type FilterParams struct {
	MinPrice      float64
	MaxPrice      float64
	MinCreateDate time.Time
	MaxCreateDate time.Time
}

func (r *Repository) GetPrices(params FilterParams) ([]Price, error) {
	prices := make([]Price, 0)
	statement := fmt.Sprintf(
		"SELECT id, name, category, price, create_date FROM prices WHERE price >= %f AND price <= %f AND create_date BETWEEN '%s' AND '%s'",
		params.MinPrice, params.MaxPrice, params.MinCreateDate.Format(constants.DateFormat), params.MaxCreateDate.Format(constants.DateFormat),
	)
	rows, err := r.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var price Price
		if err = rows.Scan(&price.ID, &price.Name, &price.Category, &price.Price, &price.CreateDate); err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}

func (r *Repository) CreatePrice(price Price) error {
	statement := fmt.Sprintf(
		"INSERT INTO prices (name, category, price, create_date) VALUES ('%s', '%s', %f, '%s')",
		price.Name, price.Category, price.Price, price.CreateDate.Format(constants.DateFormat),
	)
	_, err := r.db.Exec(statement)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTotalPriceAndUniqueCategories() (float64, int, error) {
	var totalPrice float64
	var totalCategories int
	err := r.db.QueryRow("SELECT SUM(price), COUNT(DISTINCT category) FROM prices").Scan(&totalPrice, &totalCategories)
	if err != nil {
		return 0, 0, err
	}
	return totalPrice, totalCategories, nil
}
