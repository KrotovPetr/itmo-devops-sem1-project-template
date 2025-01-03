package db

import (
	"fmt"
	"project_sem/internal/constants"
)

func (r *Repository) CreatePrice(price PriceStruct) error {
	query := fmt.Sprintf("INSERT INTO prices (id, name, category, price, create_date) VALUES (%d, '%s', '%s', %f, '%s')",
		price.ID, price.Name, price.Category, price.Price, price.CreateDate.Format(constants.DateFormat))
	_, err := r.db.Exec(query)
	return err
}

func (r *Repository) GetTotalPriceAndUniqueCategories() (float64, int, error) {
	var sum float64
	var count int
	err := r.db.QueryRow("SELECT SUM(price), COUNT(DISTINCT category) FROM prices").Scan(&sum, &count)
	if err != nil {
		return 0, 0, err
	}
	return sum, count, nil
}

func (r *Repository) GetPrices(params FilterParamsStruct) ([]PriceStruct, error) {
	var prices []PriceStruct
	query := fmt.Sprintf("SELECT * FROM prices WHERE price BETWEEN %f AND %f AND create_date BETWEEN '%s' AND '%s'",
		params.MinPrice, params.MaxPrice, params.MinCreateDate.Format(constants.DateFormat), params.MaxCreateDate.Format(constants.DateFormat))
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p PriceStruct
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.CreateDate); err != nil {
			return nil, err
		}
		prices = append(prices, p)
	}
	return prices, nil
}
