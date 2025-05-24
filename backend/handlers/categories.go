package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"expense-tracker/database"
	"expense-tracker/types"
	"github.com/gorilla/mux"
)

// GetCategories List categories as a tree
func GetCategories(w http.ResponseWriter, r *http.Request) {
	cats, err := database.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tree := buildTree(cats)
	_ = json.NewEncoder(w).Encode(tree)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		ParentID *int   `json:"parentId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Name == "" {
		http.Error(w, "Invalid input", 400)
		return
	}
	id, err := database.CreateCategory(input.Name, input.ParentID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", 400)
		return
	}
	var input struct {
		Name     *string `json:"name"`
		ParentID *int    `json:"parentId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", 400)
		return
	}
	err = database.UpdateCategory(id, input.Name, input.ParentID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(204)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", 400)
		return
	}
	err = database.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(204)
}

func buildTree(cats []types.Category) []map[string]interface{} {
	idToNode := map[int]map[string]interface{}{}
	var roots []map[string]interface{}

	for _, c := range cats {
		idToNode[c.ID] = map[string]interface{}{
			"id":       c.ID,
			"name":     c.Name,
			"parentId": c.ParentID,
			"children": []map[string]interface{}{},
		}
	}
	for _, c := range cats {
		if c.ParentID != nil {
			parent := idToNode[*c.ParentID]
			parent["children"] = append(parent["children"].([]map[string]interface{}), idToNode[c.ID])
		} else {
			roots = append(roots, idToNode[c.ID])
		}
	}
	return roots
}
