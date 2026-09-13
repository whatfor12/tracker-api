package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
)

type Server struct {
	db *sql.DB
}

func (s *Server) getExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var e Expense
	err = s.db.QueryRowContext(r.Context(),
		`SELECT id, amount, category, note, spent_at
     FROM expenses
     WHERE id = $1`,
	 id,
	).Scan(&e.ID, &e.Amount, &e.Category, &e.Note, &e.SpentAt)

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
