package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB() {
	var err error
	connStr := "host=172.17.0.1 user=riya password=riyaghosalrg@123 dbname=myappdb sslmode=disable"
	db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = createUsersTable()
	if err != nil {
		log.Fatal("error while creating the users table", err)
	}

	err = createExpensesTable()
	if err != nil {
		log.Fatal("error while creating the expenses table", err)
	}

	log.Println("Database connection established")
}

func createUsersTable() error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			total_income FLOAT DEFAULT 0,
			password VARCHAR(50) NOT NULL
		)
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating table %v", err)
	}

	return nil
}
func createExpensesTable() error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS expenses (
			id SERIAL PRIMARY KEY,
			category VARCHAR(50) NOT NULL,
			create_date DATE DEFAULT CURRENT_DATE,
            amount INT NOT NULL,
			user_id INT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating table %v", err)
	}

	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}

func GetDB() *sqlx.DB {
	return db
}
