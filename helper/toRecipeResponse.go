package helper

import (
	"foodieland/model/domain"
	"foodieland/model/web"
)

func ToNutritionResponse(nutrition domain.Nutrition) web.NutritionResponse {
    return web.NutritionResponse{
        Calories:     nutrition.Calories,
        TotalFat:     nutrition.TotalFat,
        Carbohydrate: nutrition.Carbohydrate,
        Cholesterol:  nutrition.Cholesterol,
        Description:  nutrition.Description,
    }
}

func ToDirectionResponse(direction domain.Direction) web.DirectionResponse {
    return web.DirectionResponse{
        Step:        direction.Step,
        Description: direction.Description,
        Image:         direction.Image,
    }
}

func ToDirectionsResponse(list []domain.Direction) []web.DirectionResponse {
    responses := make([]web.DirectionResponse, 0, len(list))
    for _, d := range list {
        responses = append(responses, ToDirectionResponse(d))
    }
    return responses
}



func ToRecipeResponse(recipe domain.Recipe) web.RecipeResponse {
	return web.RecipeResponse{
		Id:   uint(recipe.Id),
        Title: recipe.Title,
        Description: recipe.Description,
        Image: recipe.Image,
        PrepTime: recipe.PrepTime,
        CookTime: recipe.CookTime,
        Category: recipe.Category,
        Nutrition: ToNutritionResponse(recipe.Nutrition),
        MainDish: recipe.MainDish,
        Sauce: recipe.Sauce,
        Directions: ToDirectionsResponse(recipe.Directions),
        IsLike: recipe.IsLike,
        Writer: recipe.Writer,
	}
}