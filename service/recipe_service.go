package service

import (
	"context"
	"foodieland/model/web"
)

type RecipeService interface {
	Create(ctx context.Context, req web.RecipeCreateRequest) (web.RecipeResponse, error)
	Update(ctx context.Context, id uint, req web.RecipeUpdateRequest) (web.RecipeResponse, error) // PUT
	Patch(ctx context.Context, id uint, req web.RecipePatchRequest) (web.RecipeResponse, error) // PATCH
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (web.RecipeResponse, error)
	GetAll(ctx context.Context) ([]web.RecipeResponse, error)
}
