package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ChlukasX/give-me-food/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main () {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	amount := flag.Int("amount", 7, "The amount of food recipes")
	flag.Parse()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    os.Getenv("DB_HOST"),     // localhost
    os.Getenv("DB_PORT"),     // 5432
    os.Getenv("DB_USER"),     // postgres
    os.Getenv("DB_PASSWORD"), // password
    os.Getenv("DB_NAME"),     // myapp
	)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(connStr)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	infoLog.Printf("The amount is %d\n", *amount)

	recipes, err := models.GetAll()
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Println("The Recipes are:")

	for _, recipe := range recipes {
		infoLog.Println(recipe)
	}

	rand.Seed(time.Now().UnixNano())
	infoLog.Println("Random Recipes")

	for i := range *amount {
		randomIndex := rand.Intn(len(recipes))
		pick := recipes[randomIndex]
		infoLog.Println("nr ", i+1, ": ", pick)
	}
}

func openDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
