package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/employee", h.CreateEmployee).Methods("POST")
	return router
}
func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name       string  `json:"name"`
		PositionID int     `json:"position_id"`
		Salary     float64 `json:"salary"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := h.db.Exec(`
		INSERT INTO employees (name, position_id, salary)
		VALUES ($1, $2, $3)
	`, req.Name, req.PositionID, req.Salary)

	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
