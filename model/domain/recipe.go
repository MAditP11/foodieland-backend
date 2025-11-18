package domain

import "time"

type Nutrition struct {
	Calories string `json:"calories" validate:"required"`
	TotalFat string `json:"total_fat" validate:"required"`
	Protein string `json:"protein" validate:"required"`
	Carbohydrate string `json:"carbohydrate" validate:"required"`
	Cholesterol string `json:"cholesterol" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type Direction struct {
	Step uint
	Description string
	Img string
}

type Recipe struct {
	Id int
	Name string
	Description string
	Img string
	PrepTime string
	CookTime string
	Category string
	Nutrition Nutrition
	MainDish []string
	Sauce []string
	Directions []Direction
	IsLike bool
	Writer string
	CreateAt time.Time
}

type RecipePatch struct {
	Name        *string
	Description *string
	Img         *string
	PrepTime    *string
	CookTime    *string
	Category    *string
	Nutrition   *NutritionPatch
	MainDish    *[]string
	Sauce       *[]string
	Directions  *[]DirectionPatch
	IsLike      *bool
	Writer      *string
}

type NutritionPatch struct {
	Calories *string `json:"calories"`
	TotalFat *string `json:"total_fat"`
	Protein *string `json:"protein"`
	Carbohydrate *string `json:"carbohydrate"`
	Cholesterol *string `json:"cholesterol"`
	Description *string `json:"description"`
}

type DirectionPatch struct {
	Step *uint
	Description *string
	Img *string
}