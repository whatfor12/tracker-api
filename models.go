package main

import "time"

type Expense struct {
	ID       int64     `json:"id"`
	Amount   int64     `json:"amount"`
	Category string    `json:"category"`
	Note     string    `json:"note"`
	SpentAt  time.Time `json:"spent_at"`
}