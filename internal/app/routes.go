package app
import (
	"net/http"
	"project_sem/internal/db"
	"project_sem/internal/handlers"
)

func NewServerRouter(repo *db.Repository) *http.ServeMux {
    mux := http.NewServeMux()
    registerRoute := func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
        mux.HandleFunc(pattern, handler)
    }
    
    registerRoute("GET /api/v0/prices", handlers.GetPrices(repo))
    registerRoute("POST /api/v0/prices", handlers.CreatePrices(repo))
    
    return mux
}