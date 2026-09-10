package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Expense struct {
	ID       int64     `json:"id"`
	Amount   int64     `json:"amount"`  // в копейках
	Category string    `json:"category"`
	Note     string    `json:"note"`
	SpentAt  time.Time `json:"spent_at"`
}

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

func main() {
	http.HandleFunc("/expenses/{id}", getExpense)

	log.Fatal(http.ListenAndServe(":8080", nil))
}