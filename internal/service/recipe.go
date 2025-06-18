package service

import (
	"fmt"
	"math/rand"

	"github.com/ChlukasX/give-me-food/internal/models"
	"slices"
)

type RecipeService struct {
	recipeModel models.RecipeModelInterface
	usedRecipes map[int]bool
}

func NewRecipeService(recipeModel models.RecipeModelInterface) *RecipeService {
	return &RecipeService{
		recipeModel: recipeModel,
		usedRecipes: make(map[int]bool),
	}
}

func (s *RecipeService) RecommendUnique(count int) ([]models.Recipe, error){
	allRecipes, err := s.recipeModel.GetAll()
	if err != nil {
		return nil, err
	}

	availableRecipes := []models.Recipe{}

	for _, recipe := range allRecipes {
		if !s.usedRecipes[recipe.ID] {
			availableRecipes = append(availableRecipes, *recipe)
		}
	}

	if len(availableRecipes) < count {
		return nil, fmt.Errorf("not enough unique or unused recipes availabe for amount. Have %d, need %d", len(availableRecipes), count)
	}

	recommended := []models.Recipe{}
	for i := 0; i < count && len(availableRecipes) > 0; i++ {
		randomIndex := rand.Intn(len(availableRecipes))
		selectedRecipe := availableRecipes[randomIndex]

		recommended = append(recommended, selectedRecipe)
		s.usedRecipes[selectedRecipe.ID] = true

		availableRecipes = slices.Delete(availableRecipes, randomIndex, randomIndex+1)
	}

	return recommended, nil
}
