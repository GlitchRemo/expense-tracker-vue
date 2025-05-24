package database

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"expense-tracker/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func GetCategories() ([]types.Category, error) {
	db := GetDB()
	var cats []types.Category
	err := db.Select(&cats, "SELECT * FROM categories ORDER BY parent_id NULLS FIRST, id")
	return cats, err
}

func CreateCategory(name string, parentID *int) (int, error) {
	db := GetDB()
	var id int
	err := db.QueryRow(
		"INSERT INTO categories (name, parent_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		name, parentID, time.Now(), time.Now(),
	).Scan(&id)
	return id, err
}

func UpdateCategory(id int, name *string, parentID *int) error {
	db := GetDB()
	query := "UPDATE categories SET "
	var args []interface{}
	argIdx := 1

	if name != nil {
		query += "name = $" + strconv.Itoa(argIdx) + ", "
		args = append(args, *name)
		argIdx++
	}
	if parentID != nil {
		query += "parent_id = $" + strconv.Itoa(argIdx) + ", "
		args = append(args, *parentID)
		argIdx++
	}
	if len(args) == 0 {
		return errors.New("no fields to update")
	}
	query += "updated_at = $" + strconv.Itoa(argIdx) + " WHERE id = $" + strconv.Itoa(argIdx+1)
	args = append(args, time.Now(), id)

	_, err := db.Exec(query, args...)
	return err
}

func DeleteCategory(id int) error {
	db := GetDB()
	_, err := db.Exec("DELETE FROM categories WHERE id=$1", id)
	return err
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
