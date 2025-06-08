package models

import (
	"encoding/json"
	"os"
)

type Recipe struct {
	ID int
	Name string
}

func GetAll() ([]Recipe, error) {
	var recipes []Recipe

	recipejson := "data/recipes.json"

	data, err := os.ReadFile(recipejson)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &recipes)

	return recipes, nil
}
