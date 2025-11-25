// helper/convert.go
package helper

import (
	"foodieland/model/domain"
	"foodieland/model/web"
)

// ToRecipePatch converts web.RecipePatchRequest -> domain.RecipePatch
func ToRecipePatch(req web.RecipePatchRequest) domain.RecipePatch {
	var p domain.RecipePatch

	if req.Title != nil {
		p.Title = req.Title
	}
	if req.Description != nil {
		p.Description = req.Description
	}
	if req.Image != nil {
		p.Image = req.Image
	}
	if req.PrepTime != nil {
		p.PrepTime = req.PrepTime
	}
	if req.CookTime != nil {
		p.CookTime = req.CookTime
	}
	if req.Category != nil {
		p.Category = req.Category
	}
	// Nutrition
	if req.Nutrition != nil {
		np := domain.NutritionPatch{}
		if req.Nutrition.Calories != nil {
			np.Calories = req.Nutrition.Calories
		}
		if req.Nutrition.TotalFat != nil {
			np.TotalFat = req.Nutrition.TotalFat
		}
		if req.Nutrition.Protein != nil {
			np.Protein = req.Nutrition.Protein
		}
		if req.Nutrition.Carbohydrate != nil {
			np.Carbohydrate = req.Nutrition.Carbohydrate
		}
		if req.Nutrition.Cholesterol != nil {
			np.Cholesterol = req.Nutrition.Cholesterol
		}
		if req.Nutrition.Description != nil {
			np.Description = req.Nutrition.Description
		}
		p.Nutrition = &np
	}
	// MainDish (slice of string)
	if req.MainDish != nil {
		p.MainDish = req.MainDish
	}
	// Sauce (slice of string)
	if req.Sauce != nil {
		p.Sauce = req.Sauce
	}
	// Directions (slice of DirectionPatch)
	if req.Directions != nil {
		var dirs []domain.DirectionPatch
		for _, d := range *req.Directions {
			dp := domain.DirectionPatch{}
			if d.Step != nil {
				dp.Step = d.Step
			}
			if d.Description != nil {
				dp.Description = d.Description
			}
			if d.Image != nil {
				dp.Image = d.Image
			}
			dirs = append(dirs, dp)
		}
		p.Directions = &dirs
	}
	if req.IsLike != nil {
		p.IsLike = req.IsLike
	}
	if req.Writer != nil {
		p.Writer = req.Writer
	}

	return p
}
