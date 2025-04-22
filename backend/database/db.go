package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := "user=riya password=riyaghosalrg@123 dbname=myappdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Database connection established")
}

func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}
