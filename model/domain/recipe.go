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
	Img *string
}