package main

import (
	"database/sql"
	"go-standart-http/internal/router"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "file:demo.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := initializeDB(db); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	mux := router.NewRouter(db)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initializeDB(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );`
	_, err := db.Exec(query)
	return err
}
