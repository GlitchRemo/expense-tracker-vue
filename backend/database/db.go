package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func InitDB() {
	var err error
	connStr := "host=localhost user=riya password=riyaghosalrg@123 dbname=myappdb sslmode=disable"
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

	err = createCategoriesTable()
	if err != nil {
		log.Fatal("error while creating the categories table", err)
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
		return fmt.Errorf("error creating users table %v", err)
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
		return fmt.Errorf("error creating expenses table %v", err)
	}

	return nil
}

func createCategoriesTable() error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS categories(
    		id SERIAL PRIMARY KEY,
    		name VARCHAR(255) NOT NULL,
    		parent_id INT NULL,
    		created_at TIMESTAMP DEFAULT NOW(),
    		updated_at TIMESTAMP DEFAULT NOW(),
    		CONSTRAINT fk_parent FOREIGN KEY (parent_id) REFERENCES categories (id) ON DELETE CASCADE
			);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating categories table %v", err)
	}

	return nil
}

func CloseDB() {
	if db != nil {
		_ = db.Close()
		log.Println("Database connection closed")
	}
}

func GetDB() *sqlx.DB {
	return db
}
