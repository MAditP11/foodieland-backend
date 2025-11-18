package repository

import (
	"context"
	"database/sql"
	"foodieland/model/domain"
)

type RecipeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, recipe domain.Recipe) (domain.Recipe, error)
	Update(ctx context.Context, tx *sql.Tx, recipe domain.Recipe) (domain.Recipe, error)
	Patch(ctx context.Context, tx *sql.Tx, recipeId int, patch domain.RecipePatch ) error
	Delete(ctx context.Context, tx *sql.Tx, recipeId int) error
	GetById(ctx context.Context, tx *sql.Tx, recipeId int) (domain.Recipe, error)
	GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Recipe, error)
}
