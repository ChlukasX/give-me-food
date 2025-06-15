package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ChlukasX/give-me-food/internal/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	recipes models.RecipeModelInterface
}

func main () {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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

	app := &Application{
		errorLog: errorLog,
		infoLog: infoLog,
		recipes: &models.RecipeModel{DB: db},
	}

	recipes, err := app.recipes.GetAll()
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Print(recipes)
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
