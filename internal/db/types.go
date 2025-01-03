package db

import "time"

type PriceStruct struct {
	ID         int
	Name       string
	Category   string
	Price      float64
	CreateDate time.Time
}

type FilterParamsStruct struct {
	MinPrice      float64
	MaxPrice      float64
	MinCreateDate time.Time
	MaxCreateDate time.Time
}
