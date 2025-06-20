package models

import (
	"database/sql"
)

type ComponentModelInterface interface {
	Insert(name string, componentType ComponentType) error
}

type ComponentType string

const (
	Protein ComponentType = "protein"
	Sauce ComponentType = "sauce"
	Aromatics ComponentType = "aromatics"
	Vegtable ComponentType = "vegtable"
	Garnish ComponentType = "garnish"
	Seasonings ComponentType = "seasoning"
)

type RecipeComponent struct {
	ID int
	Name string
	ComponentType ComponentType
	Ingredients []Ingredient
}

type ComponentModel struct {
	DB *sql.DB
}

func (m *ComponentModel) Insert(name string, componentType ComponentType) error {
	stmt := `INSERT INTO component(name, component_type)
	VALUES(?, ?)`

	// TODO: handle component type const before Insert

	_, err := m.DB.Exec(stmt, name, componentType)
	if err != nil {
		return err
	}

	return nil
}
