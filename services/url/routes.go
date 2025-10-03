package url

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler holds the HTTP routes and logic
type Handler struct{}

// constructor
func NewHandler() *Handler {
	return &Handler{}
}

// attach routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/metrics", h.handleMetrics).Methods("POST")
}

// actual HTTP handler

// hadle metrics for batch urls

func (h *Handler) handleMetrics(w http.ResponseWriter, r *http.Request) {
	var urls []string

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&urls); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if len(urls) == 0 {
		http.Error(w, "no URLs provided", http.StatusBadRequest)
		return
	}

	// collect metrics
	// results := make([]string, 0, len(urls))
	// for _, url := range urls {
	// 	results = append(results, h.GetMetrics(url))
	// }
	results := h.GetBatchMetrics(urls)

	// return as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
