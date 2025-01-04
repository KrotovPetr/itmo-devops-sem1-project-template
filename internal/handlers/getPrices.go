package handlers
import (
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
	"project_sem/internal/db"
	"project_sem/internal/serializers"
	"project_sem/internal/utils"
)
func GetPrices(repo *db.Repository) http.HandlerFunc {
	const errorResponseBody = "failed to load prices"
	const successContentType = "application/zip"
	const sucessContentDisposition = "attachment; filename=data.zip"
	const csvFileName = "data.csv"
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
		err = utils.ZipFile(serializedPrices, w, csvFileName)
		if err != nil {
			log.Printf("failed to archive prices: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", successContentType)
		w.Header().Set("Content-Disposition", sucessContentDisposition)
	}
}
func buildFilterParams(r *http.Request) (db.FilterParams, error) {
	params := db.FilterParams{}
	minCreateDate := r.URL.Query().Get("start")
	if minCreateDate != "" {
		minDate, err := time.Parse("2006-01-02", minCreateDate)
		if err != nil {
			return params, err
		}
		params.MinCreateDate = minDate
	} else {
		params.MinCreateDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	maxCreateDate := r.URL.Query().Get("end")
	if maxCreateDate != "" {
		maxDate, err := time.Parse("2006-01-02", maxCreateDate)
		if err != nil {
			return params, err
		}
		params.MaxCreateDate = maxDate
	} else {
		params.MaxCreateDate = time.Now()
	}
	minPrice := r.URL.Query().Get("min")
	if minPrice != "" {
		price, err := strconv.ParseFloat(minPrice, 64)
		if err != nil {
			return params, err
		}
		params.MinPrice = price
	}
	maxPrice := r.URL.Query().Get("max")
	if maxPrice != "" {
		price, err := strconv.ParseFloat(maxPrice, 64)
		if err != nil {
			return params, err
		}
		params.MaxPrice = price
	} else {
		params.MaxPrice = math.MaxFloat64
	}
	return params, nil
}