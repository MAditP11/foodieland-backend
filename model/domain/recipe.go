package domain

import "time"

type Nutrition struct {
	Calories string 
	TotalFat string 
	Protein string 
	Carbohydrate string 
	Cholesterol string 
	Description string 
}

type Direction struct {
	Step uint
	Description string
	Image string
}

type Recipe struct {
	Id int
	Title string
	Description string
	Image string
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
	Title        *string
	Description *string
	Image         *string
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
	Calories *string 
	TotalFat *string 
	Protein *string 
	Carbohydrate *string 
	Cholesterol *string 
	Description *string 
}

type DirectionPatch struct {
	Step *uint
	Description *string
	Image *string
}