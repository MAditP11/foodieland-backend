package main

import (
	"foodieland/app"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    db := app.NewDB()
	recipeController := app.InitializeRecipeController(db)
	router := app.NewRouter(recipeController)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
