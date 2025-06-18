package models

import (
	"database/sql"
)

type RecipeModelInterface interface {
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
