//go:build wireinject
// +build wireinject

package app

import (
	"database/sql"
	"foodieland/controller"
	"foodieland/repository"
	"foodieland/service"
	"foodieland/validation"

	"github.com/google/wire"
)

func InitializeRecipeController(db *sql.DB) controller.RecipeController {
    wire.Build(
        repository.NewRecipeRepositoryImpl,
        service.NewRecipeServiceImpl,
        controller.NewRecipeControllerImpl,
        validation.NewValidator,
    )
    return nil
}
