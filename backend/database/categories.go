package database

import (
	"errors"
	"expense-tracker/types"
	"strconv"
	"time"
)

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
