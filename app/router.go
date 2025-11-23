package app

import (
	"foodieland/controller"
	"foodieland/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(recipeController controller.RecipeController ) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/recipes", recipeController.GetAll)
	router.GET("/api/recipe/:id", recipeController.GetById)
	router.POST("/api/recipe", recipeController.Create)
	router.PUT("/api/recipe/:id", recipeController.Update)
	router.PATCH("/api/recipe/:id", recipeController.Patch)
	router.DELETE("/api/recipe/:id", recipeController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}