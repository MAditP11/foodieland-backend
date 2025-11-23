package web

import (
	"time"
)

type RecipeCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image string `json:"image" validate:"required"`
	PrepTime string `json:"prep_time" validate:"required"`
	CookTime string `json:"cook_time" validate:"required"`
	Category string `json:"category" validate:"required"`
	Nutrition NutritionRequest `json:"nutrition" validate:"required"`
	MainDish []string `json:"main_dish" validate:"required"`
	Sauce []string `json:"sauce" validate:"required"`
	Directions []DirectionRequest `json:"directions" validate:"required"`
	IsLike bool `json:"is_like"`
	Writer string `json:"writer" validate:"required"`
	CreateAt time.Time `json:"create_at"`

}