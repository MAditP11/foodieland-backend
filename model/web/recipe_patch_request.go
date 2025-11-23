package web

type RecipePatchRequest struct {
	Id *uint `json:"-"`
	Title *string `json:"title"`
	Description *string `json:"description"`
	Image *string `json:"image"`
	PrepTime *string `json:"prep_time"`
	CookTime *string `json:"cook_time"`
	Category *string `json:"category"`
	Nutrition *NutritionPatchRequest `json:"nutrition"`
	MainDish *[]string `json:"main_dish"`
	Sauce *[]string `json:"sauce"`
	Directions *[]DirectionPatchRequest `json:"directions"`
	IsLike *bool `json:"is_like"`
	Writer *string `json:"writer"`
}