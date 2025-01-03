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

func SerializePrices(prices []db.PriceStruct) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	csvWriter := csv.NewWriter(&buffer)
	defer csvWriter.Flush()

	csvWriter.Write([]string{"id", "name", "category", "price", "create_date"})
	for _, price := range prices {
		record := []string{
			fmt.Sprintf("%d", price.ID),
			price.Name,
			price.Category,
			fmt.Sprintf("%.2f", price.Price),
			price.CreateDate.Format(constants.DateFormat),
		}
		if err := csvWriter.Write(record); err != nil {
			return nil, fmt.Errorf("error writing CSV record: %w", err)
		}
	}

	return &buffer, nil
}

func DeserializePrices(r io.Reader) ([]db.PriceStruct, int, error) {
	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, 0, fmt.Errorf("error reading CSV data: %w", err)
	}

	var prices []db.PriceStruct
	totalItems := len(records) - 1

	for idx, record := range records {
		if idx == 0 {
			continue
		}
		price, err := validatePrice(record)
		if err != nil {
			log.Printf("invalid price record: %v, error: %v", record, err)
			continue
		}
		prices = append(prices, price)
	}

	return prices, totalItems, nil
}

func validatePrice(record []string) (db.PriceStruct, error) {
	if len(record) != 5 {
		return db.PriceStruct{}, fmt.Errorf("expected 5 fields, got %d: %v", len(record), record)
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return db.PriceStruct{}, fmt.Errorf("invalid ID \"%v\": %w", record[0], err)
	}

	name := record[1]
	if name == "" {
		return db.PriceStruct{}, fmt.Errorf("name field is empty")
	}

	category := record[2]
	if category == "" {
		return db.PriceStruct{}, fmt.Errorf("category field is empty")
	}

	cost, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return db.PriceStruct{}, fmt.Errorf("invalid price \"%v\": %w", record[3], err)
	}

	createDate, err := time.Parse(constants.DateFormat, record[4])
	if err != nil {
		return db.PriceStruct{}, fmt.Errorf("invalid date \"%v\": %w", record[4], err)
	}

	return db.PriceStruct{
		ID:         id,
		Name:       name,
		Category:   category,
		Price:      cost,
		CreateDate: createDate,
	}, nil
}
