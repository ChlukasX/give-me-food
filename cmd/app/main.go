package main

import (
	"flag"
	"fmt"
)

func main () {
	amount := flag.Int("amount", 7, "The amount of food recipes")

	flag.Parse()

	fmt.Printf("The amount is %d", *amount)
}
