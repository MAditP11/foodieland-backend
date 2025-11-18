package web

import (
	"foodieland/model/domain"
)

type RecipePatchRequest struct {
	Name *string `json:"name"`
	Description *string `json:"description"`
	Img *string `json:"img"`
	PrepTime *string `json:"prep_time"`
	CookTime *string `json:"cook_time"`
	Category *string `json:"category"`
	Nutrition *domain.NutritionPatch `json:"nutrition"`
	MainDish *[]string `json:"main_dish"`
	Sauce *[]string `json:"sauce"`
	Directions *[]domain.DirectionPatch `json:"directions"`
	IsLike *bool `json:"is_like"`
	Writer *string `json:"writer"`
}