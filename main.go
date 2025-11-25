package main

import (
	"foodieland/app"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    db := app.NewDB()
	recipeController := app.InitializeRecipeController(db)
	handler := app.NewRouter(recipeController) // sudah ada CORS + static files

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server is running on http://localhost:8080")
	server.ListenAndServe()
}
