package web

type NutritionRequest struct {
    Calories     string    `json:"calories" validate:"required"`
    TotalFat     string    `json:"total_fat" validate:"required"`
    Protein      string    `json:"protein" validate:"required"`
    Carbohydrate string    `json:"carbohydrate" validate:"required"`
    Cholesterol  string    `json:"cholesterol" validate:"required"`
    Description  string    `json:"description" validate:"required"`
}

type NutritionPatchRequest struct {
    Calories     *string    `json:"calories"`
    TotalFat     *string    `json:"total_fat"`
    Protein      *string    `json:"protein"`
    Carbohydrate *string    `json:"carbohydrate"`
    Cholesterol  *string    `json:"cholesterol"`
    Description  *string    `json:"description"`
}
type NutritionResponse struct {
    Calories     string    `json:"calories"`
    TotalFat     string    `json:"total_fat"`
    Protein      string    `json:"protein"`
    Carbohydrate string    `json:"carbohydrate"`
    Cholesterol  string    `json:"cholesterol"`
    Description  string    `json:"description"`
}
