package web

import (
	"time"
)

type RecipeResponse struct {
	Id uint  `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Image string `json:"image"`
	PrepTime string `json:"prep_time"`
	CookTime string `json:"cook_time"`
	Category string `json:"category"`
	Nutrition NutritionResponse `json:"nutrition"`
	MainDish []string `json:"main_dish"`
	Sauce []string `json:"sauce"`
	Directions []DirectionResponse `json:"directions"`
	IsLike bool `json:"is_like"`
	Writer string `json:"writer"`
	CreateAt time.Time `json:"create_at"`
}