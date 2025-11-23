package web

type RecipeUpdateRequest struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Img string `json:"img" validate:"required"`
	PrepTime string `json:"prep_time" validate:"required"`
	CookTime string `json:"cook_time" validate:"required"`
	Category string `json:"category" validate:"required"`
	Nutrition NutritionRequest `json:"nutrition" validate:"required"`
	MainDish []string `json:"main_dish" validate:"required"`
	Sauce []string `json:"sauce" validate:"required"`
	Directions []DirectionRequest `json:"directions" validate:"required"`
	IsLike bool `json:"is_like" validate:"required"`
	Writer string `json:"writer" validate:"required"`
}