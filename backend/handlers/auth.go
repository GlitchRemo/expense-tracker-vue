package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"expense-tracker/database"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	var userID int
	query := `SELECT id FROM users WHERE username = $1 AND password = $2`
	err := db.QueryRow(query, req.Username, req.Password).Scan(&userID)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		UserID:   userID,
		Username: req.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
