package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Server struct {
	db *sql.DB
}

func (s *Server) getExpense(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func (s *Server) createExpense(w http.ResponseWriter, r *http.Request) {
	var e Expense
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var id int64
	err = s.db.QueryRowContext(r.Context(),
		`INSERT INTO expenses (amount, category, note, spent_at)
     VALUES ($1, $2, $3, $4)
     RETURNING id`,
		e.Amount, e.Category, e.Note, e.SpentAt,
	).Scan(&id)

	e.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func (s *Server) deleteExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	result, err := s.db.ExecContext(r.Context(),
		`DELETE FROM expenses WHERE id = $1`,
		id,
	)

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) updateExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var e Expense
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.db.QueryRowContext(r.Context(),
		`UPDATE expenses
     SET amount = $1, category = $2, note = $3, spent_at = $4
     WHERE id = $5
     RETURNING id, amount, category, note, spent_at`,
		e.Amount, e.Category, e.Note, e.SpentAt, id,
	).Scan(&e.ID, &e.Amount, &e.Category, &e.Note, &e.SpentAt)

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}
