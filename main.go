package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/expenses/{id}", getExpense)

	log.Fatal(http.ListenAndServe(":8080", nil))
}