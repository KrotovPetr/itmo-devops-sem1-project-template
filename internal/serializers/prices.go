package serializers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"project_sem/internal/constants"
	"project_sem/internal/db"
	"strconv"
	"time"
)

func SerializePrices(prices []db.Price) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	csvWriter := csv.NewWriter(&buffer)
	defer csvWriter.Flush()

	if err := csvWriter.Write([]string{"id", "name", "category", "price", "create_date"}); err != nil {
		return nil, fmt.Errorf("failed to write header: %w", err)
	}

	for _, price := range prices {
		record := formatPriceRecord(price)
		if err := csvWriter.Write(record); err != nil {
			return nil, fmt.Errorf("failed to write record: %w", err)
		}
	}

	return &buffer, nil
}

func DeserializePrices(r io.Reader) ([]db.Price, int, error) {
	csvReader := csv.NewReader(r)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read records: %w", err)
	}

	if len(records) <= 1 {
		return nil, 0, fmt.Errorf("no data records found")
	}

	prices := make([]db.Price, 0, len(records)-1)

	for idx, record := range records {
		if idx == 0 {
			continue
		}

		price, err := parsePriceRecord(record)
		if err != nil {
			log.Printf("price conversion failed: %v\n", err)
			continue
		}

		prices = append(prices, price)
	}

	return prices, len(records) - 1, nil
}

func formatPriceRecord(price db.Price) []string {
	return []string{
		fmt.Sprintf("%d", price.ID),
		price.Name,
		price.Category,
		fmt.Sprintf("%.2f", price.Price),
		price.CreateDate.Format(constants.DateFormat),
	}
}

func parsePriceRecord(record []string) (db.Price, error) {
	if len(record) != 5 {
		return db.Price{}, fmt.Errorf("invalid record format: %v", record)
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return db.Price{}, fmt.Errorf("invalid id %q: %w", record[0], err)
	}

	name := record[1]
	if name == "" {
		return db.Price{}, fmt.Errorf("name cannot be empty")
	}

	category := record[2]
	if category == "" {
		return db.Price{}, fmt.Errorf("category cannot be empty")
	}

	cost, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return db.Price{}, fmt.Errorf("invalid cost %q: %w", record[3], err)
	}

	createDate, err := time.Parse(constants.DateFormat, record[4])
	if err != nil {
		return db.Price{}, fmt.Errorf("invalid creation date %q: %w", record[4], err)
	}

	return db.Price{
		ID:         id,
		Name:       name,
		Category:   category,
		Price:      cost,
		CreateDate: createDate,
	}, nil
}
