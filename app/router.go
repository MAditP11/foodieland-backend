package app

import (
	"foodieland/controller"
	"foodieland/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(recipeController controller.RecipeController ) *httprouter.Router {
	router := httprouter.New()

	router.GET("api/recipes", recipeController.GetAll)
	router.GET("api/recipe/:recipeId", recipeController.GetById)
	router.POST("api/recipe", recipeController.Create)
	router.PUT("api/recipe/:recipeId", recipeController.Update)
	router.PATCH("api/recipe/:recipeId", recipeController.Patch)
	router.DELETE("api/recipe/:recipeId", recipeController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}