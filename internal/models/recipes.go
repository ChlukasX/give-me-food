package models

import (
	"encoding/json"
	"os"
)

type RecipeType string

const (
	Main RecipeType = "main"
	Side RecipeType = "side"
)

type BlockType string

const (
	Protein BlockType = "protein"
	Sauce BlockType = "sauce"
	Aromatics BlockType = "aromatics"
	Vegtable BlockType = "vegtable"
	Garnish BlockType = "garnish"
	Seasonings BlockType = "seasoning"
)

type RecipeComponent struct {
	ID int
	Name string
	Type BlockType
	Ingredients []Ingredient
}

type Recipe struct {
	ID int
	Name string
	Type RecipeType
	Instructions string
	Component []RecipeComponent
}

func GetAll() ([]Recipe, error) {
	var recipes []Recipe

	recipejson := "data/recipes.json"

	data, err := os.ReadFile(recipejson)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
