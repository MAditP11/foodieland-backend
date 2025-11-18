package web

import (
	"foodieland/model/domain"
	"time"
)

type RecipeResponse struct {
	Id uint  `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Img string `json:"img"`
	PrepTime string `json:"prep_time"`
	CookTime string `json:"cook_time"`
	Category string `json:"category"`
	Nutrition domain.Nutrition `json:"nutrition"`
	MainDish []string `json:"main_dish"`
	Sauce []string `json:"sauce"`
	Directions []domain.Direction `json:"directions"`
	IsLike bool `json:"is_like"`
	Writer string `json:"writer"`
	CreateAt time.Time `json:"create_at"`
}