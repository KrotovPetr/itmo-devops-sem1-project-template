package handlers

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"project_sem/internal/archive"
	"project_sem/internal/constants"
	"project_sem/internal/db"
	"project_sem/internal/serializers"
	"strconv"
	"time"

)

type PriceStats struct {
	TotalCount      int `json:"total_count"`
	DuplicateCount  int `json:"duplicates_count"`
	TotalItems      int `json:"total_items"`
	TotalCategories int `json:"total_categories"`
	TotalPrice      int `json:"total_price"`
}

func handleError(w http.ResponseWriter, err error, logMsg string, httpStatus int) bool {
	if err != nil {
		log.Printf("%s: %v\n", logMsg, err)
		http.Error(w, ErrorResponseBody, httpStatus)
		return true
	}
	return false
}

func logErrorAndRespond(w http.ResponseWriter, logMsg string, err error, responseBody string) {
	log.Printf("%s: %v\n", logMsg, err)
	http.Error(w, responseBody, http.StatusInternalServerError)
}

func parseDateParam(param string, defaultValue time.Time) (time.Time, error) {
	if param == "" {
		return defaultValue, nil
	}
	parsedDate, err := time.Parse(constants.DateFormat, param)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}

func parsePriceParam(param string, defaultValue float64) (float64, error) {
	if param == "" {
		return defaultValue, nil
	}
	parsedPrice, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, err
	}
	return parsedPrice, nil
}

func buildFilterParams(r *http.Request) (db.FilterParams, error) {
	params := db.FilterParams{}

	minCreateDate, err := parseDateParam(r.URL.Query().Get("start"), time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		return params, err
	}
	params.MinCreateDate = minCreateDate

	maxCreateDate, err := parseDateParam(r.URL.Query().Get("end"), time.Now())
	if err != nil {
		return params, err
	}
	params.MaxCreateDate = maxCreateDate

	minPrice, err := parsePriceParam(r.URL.Query().Get("min"), 0)
	if err != nil {
		return params, err
	}
	params.MinPrice = minPrice

	maxPrice, err := parsePriceParam(r.URL.Query().Get("max"), math.MaxFloat64)
	if err != nil {
		return params, err
	}
	params.MaxPrice = maxPrice

	return params, nil
}

func GetPrices(repo *db.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := buildFilterParams(r)
		if handleError(w, err, "Error tobuild filter params", http.StatusBadRequest) {
			return
		}

		prices, err := repo.GetPrices(params)
		if handleError(w, err, "Error to load prices", http.StatusInternalServerError) {
			return
		}

		serializedPrices, err := serializers.SerializePrices(prices)
		if handleError(w, err, "Error to serialize prices", http.StatusInternalServerError) {
			return
		}

		err = archive.ZipFile(serializedPrices, w, CSVFileName)
		if handleError(w, err, "Error to archive prices", http.StatusInternalServerError) {
			return
		}

		w.Header().Set("Content-Type", SuccessContentType)
		w.Header().Set("Content-Disposition", SuccessContentDisposition)
	}
}

func CreatePrices(repo *db.Repository) http.HandlerFunc {
	const errorResponseBody = "failed to upload prices"
	const successContentType = "application/json"

	return func(w http.ResponseWriter, r *http.Request) {
		file, err := getFileFromRequest(r)
		if err != nil {
			logErrorAndRespond(w, "failed to read incoming file", err, errorResponseBody)
			return
		}
		defer file.Close()

		formatType := r.URL.Query().Get("type")
		rc, err := unarchiveFile(file, formatType)
		if err != nil {
			logErrorAndRespond(w, "failed to unarchive incoming file", err, errorResponseBody)
			return
		}
		defer rc.Close()

		stats := PriceStats{}
		prices, totalCount, err := serializers.DeserializePrices(rc)
		if err != nil {
			logErrorAndRespond(w, "failed to deserialize prices", err, errorResponseBody)
			return
		}
		stats.TotalCount = totalCount

		transaction, err := repo.Begin()
		if err != nil {
			logErrorAndRespond(w, "failed to begin transaction", err, errorResponseBody)
			return
		}
		defer transaction.Rollback()

		for _, product := range prices {
			err = repo.CreatePrice(product)
			if err == nil {
				stats.TotalItems++
				continue
			}
			logErrorAndRespond(w, "failed to save product", err, errorResponseBody)
			return
		}

		err = transaction.Commit()
		if err != nil {
			logErrorAndRespond(w, "failed to commit transaction", err, errorResponseBody)
			return
		}

		totalPrice, totalCategories, err := repo.GetTotalPriceAndUniqueCategories()
		if err != nil {
			logErrorAndRespond(w, "failed to get total price and unique categories", err, errorResponseBody)
			return
		}
		stats.TotalCategories = totalCategories
		stats.TotalPrice = int(totalPrice)

		w.Header().Set("Content-Type", successContentType)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(stats)
	}
}

func getFileFromRequest(r *http.Request) (io.ReadCloser, error) {
	file, _, err := r.FormFile("file")
	return file, err
}

func unarchiveFile(r io.Reader, fileType string) (io.ReadCloser, error) {
	if fileType == "tar" {
		return archive.UntarFile(r)
	}
	return archive.UnzipFile(r)
}