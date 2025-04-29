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
		CREATE TABLE IF NOT EXISTS users(
    		id SERIAL NOT NULL,
    		username varchar(50) NOT NULL,
    		total_income double precision DEFAULT 0,
    		password varchar(50) NOT NULL,
    		PRIMARY KEY(id)
		);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating table %v", err)
	}

	return nil
}
func createExpensesTable() error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS expenses(
    		id SERIAL NOT NULL,
    		category varchar(50) NOT NULL,
    		create_date date DEFAULT CURRENT_DATE,
    		amount integer NOT NULL,
    		note varchar(255),
    		user_id integer NOT NULL,
    		PRIMARY KEY(id),
    		CONSTRAINT expenses_user_id_fkey FOREIGN key(user_id) REFERENCES users(id)
		);
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
