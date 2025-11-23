package web

type RecipePatchRequest struct {
	Id *uint `json:"-"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Img *string `json:"img"`
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