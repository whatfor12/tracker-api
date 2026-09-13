package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connURL := os.Getenv("DATABASE_URL")
	if connURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx := context.Background()

	db, err := connectDB(ctx, connURL)
	if err != nil {
		log.Fatal("Couldn't connect to DB", err)
	}
	defer db.Close()

	log.Println("Connected to DB")

	server := &Server{db: db}

	http.HandleFunc("GET /expenses/{id}", server.getExpense)
	http.HandleFunc("POST /expenses", server.createExpense)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
