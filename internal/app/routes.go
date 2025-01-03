package app

import (
	"net/http"
	"project_sem/internal/constants"
	"project_sem/internal/db"
	"project_sem/internal/handlers"
)

func registerGetRoutes(mux *http.ServeMux, repo *db.Repository) {
	mux.HandleFunc(constants.RoutePrices, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == constants.MethodGet {
			handlers.GetPrices(repo)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func registerPostRoutes(mux *http.ServeMux, repo *db.Repository) {
	mux.HandleFunc(constants.RoutePrices, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == constants.MethodPost {
			handlers.CreatePrices(repo)(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func NewServerRouter(repo *db.Repository) *http.ServeMux {
	mux := http.NewServeMux()

	registerGetRoutes(mux, repo)
	registerPostRoutes(mux, repo)

	return mux
}
