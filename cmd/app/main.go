package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/ChlukasX/give-me-food/internal/models"
)

func main () {
	amount := flag.Int("amount", 7, "The amount of food recipes")

	flag.Parse()

	fmt.Printf("The amount is %d\n", *amount)

	recipes, err := models.GetAll()
	if err != nil {
		panic("Oh no")
	}

	fmt.Println("The Recipes are:")

	for _, recipe := range recipes {
		fmt.Println(recipe)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random Recipes")

	for i := range *amount {
		randomIndex := rand.Intn(len(recipes))
		pick := recipes[randomIndex]
		fmt.Println("nr ", i+1, ": ", pick)
	}
}
