package models

import (
	"database/sql"
)

type RecipeModelInterface interface {
	Insert(name, recipeType, description string) (error)
	GetAll() ([]*Recipe, error)
}

type Recipe struct {
	ID int
	Name string
	RecipeType string
	Instructions string
	Component []RecipeComponent
}

type RecipeModel struct {
	DB *sql.DB
}

func (m *RecipeModel) Insert(name, recipeType, description string) (error) {
	stmt := `INSERT INTO recipes (name, recipe_type, description)
	VALUES(?, ?, ?)`

	_, err := m.DB.Exec(stmt, name, recipeType, description)
	if err != nil {
		return err
	}

	return nil
}
func (m *RecipeModel) GetAll() ([]*Recipe, error) {
	stmt := `SELECT r.id, r.name, r.recipe_type, r.instructions from recipes r;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := []*Recipe{}

	for rows.Next() {
		r := &Recipe{}

		err = rows.Scan(&r.ID, &r.Name, &r.RecipeType, &r.Instructions)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}
