package app

import (
	"foodieland/controller"
	"foodieland/exception"
	"foodieland/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(recipeController controller.RecipeController ) http.Handler {
	router := httprouter.New()

	router.GET("/api/recipes", recipeController.GetAll)
	router.GET("/api/recipe/:id", recipeController.GetById)
	router.POST("/api/recipe", recipeController.Create)
	router.PUT("/api/recipe/:id", recipeController.Update)
	router.PATCH("/api/recipe/:id", recipeController.Patch)
	router.DELETE("/api/recipe/:id", recipeController.Delete)

	// 18.
	router.ServeFiles("/uploads/*filepath", http.Dir("./uploads"))

	router.PanicHandler = exception.ErrorHandler
	//20.
	return middleware.CORSMiddleware(router)
}