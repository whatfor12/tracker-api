package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func getExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	purchase := Expense{
		ID:       1,
		Amount:   450,
		Category: "Еда",
		Note:     "Обед в кафе",
		SpentAt:  time.Now().Truncate(time.Second),
	}

	jsonBytes, err := json.Marshal(purchase)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}