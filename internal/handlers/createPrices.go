package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"project_sem/internal/constants"
	"project_sem/internal/db"
	"project_sem/internal/serializers"
	"project_sem/internal/utils"
)

func CreatePrices(repo *db.Repository) http.HandlerFunc {
	const (
		errorResponseBody  = "failed to upload prices"
		successContentType = "application/json"
	)
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			log.Printf("failed to read incoming file: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		defer file.Close()

		formatType := r.URL.Query().Get("type")
		rc, err := unarchiveFile(file, formatType)
		if err != nil {
			log.Printf("failed to unarchive incoming file: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		defer rc.Close()

		stats := PriceStatsStruct{}

		prices, totalCount, err := serializers.DeserializePrices(rc)
		if err != nil {
			log.Printf("failed to parse prices from incoming file: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		stats.TotalCount = totalCount

		for _, price := range prices {
			if err := repo.CreatePrice(price); err != nil {
				stats.DuplicateCount++
			} else {
				stats.TotalItems++
			}
		}

		totalPrice, totalCategories, err := repo.GetTotalPriceAndUniqueCategories()
		if err != nil {
			log.Printf("failed to get total price and unique categories: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		stats.TotalCategories = totalCategories
		stats.TotalPrice = int(totalPrice)

		w.Header().Set("Content-Type", successContentType)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(stats)
	}
}

func unarchiveFile(r io.Reader, fileType string) (io.ReadCloser, error) {
	switch fileType {
	case constants.Tar:
		return utils.UntarFile(r)
	default:
		return utils.UnzipFile(r)
	}
}
