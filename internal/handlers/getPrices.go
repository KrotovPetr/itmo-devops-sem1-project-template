package handlers

import (
	"log"
	"math"
	"net/http"
	"project_sem/internal/constants"
	"project_sem/internal/db"
	"project_sem/internal/serializers"
	"project_sem/internal/utils"
	"strconv"
	"time"
)

func GetPrices(repo *db.Repository) http.HandlerFunc {
	const (
		errorResponseBody         = "download error"
		successContentType        = "application/zip"
		successContentDisposition = "attachment; filename=data.zip"
		csvFileName               = "data.csv"
	)
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := buildFilterParams(r)
		if err != nil {
			log.Printf("failed to build filter params: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusBadRequest)
			return
		}
		prices, err := repo.GetPrices(params)
		if err != nil {
			log.Printf("failed to load prices: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		serializedPrices, err := serializers.SerializePrices(prices)
		if err != nil {
			log.Printf("failed to serialize prices: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		if err = utils.ZipFile(serializedPrices, w, csvFileName); err != nil {
			log.Printf("failed to archive prices: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", successContentType)
		w.Header().Set("Content-Disposition", successContentDisposition)
	}
}

func buildFilterParams(r *http.Request) (db.FilterParamsStruct, error) {
	params := db.FilterParamsStruct{}
	if minCreateDate := r.URL.Query().Get("start"); minCreateDate != "" {
		if minDate, err := time.Parse(constants.DateFormat, minCreateDate); err != nil {
			return params, err
		} else {
			params.MinCreateDate = minDate
		}
	} else {
		params.MinCreateDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	if maxCreateDate := r.URL.Query().Get("end"); maxCreateDate != "" {
		if maxDate, err := time.Parse(constants.DateFormat, maxCreateDate); err != nil {
			return params, err
		} else {
			params.MaxCreateDate = maxDate
		}
	} else {
		params.MaxCreateDate = time.Now()
	}
	if minPrice := r.URL.Query().Get("min"); minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err != nil {
			return params, err
		} else {
			params.MinPrice = price
		}
	}
	if maxPrice := r.URL.Query().Get("max"); maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err != nil {
			return params, err
		} else {
			params.MaxPrice = price
		}
	} else {
		params.MaxPrice = math.MaxFloat64
	}
	return params, nil
}
