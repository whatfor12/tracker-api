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

	http.HandleFunc("/expenses/{id}", server.getExpense)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
