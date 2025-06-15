package models

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
	Type ComponentType
	Ingredients []Ingredient
}
