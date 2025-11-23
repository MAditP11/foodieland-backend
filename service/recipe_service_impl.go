package service

import (
	"context"
	"database/sql"
	"foodieland/helper"
	"foodieland/model/domain"
	"foodieland/model/web"
	"foodieland/repository"

	"github.com/go-playground/validator/v10"
)

type RecipeServiceImpl struct {
	RecipeRepository repository.RecipeRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewRecipeServiceImpl(recipeRepository repository.RecipeRepository, DB *sql.DB, validate *validator.Validate) RecipeService {
	return &RecipeServiceImpl{
		RecipeRepository: recipeRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *RecipeServiceImpl) Create(ctx context.Context, req web.RecipeCreateRequest) (web.RecipeResponse, error) {
	if err := service.Validate.Struct(req); err != nil {
		panic(err)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	recipe := domain.Recipe{
		Name: req.Name,
		Description: req.Description, 
		Img: req.Img, 
		PrepTime:req.PrepTime, 
		CookTime:req.CookTime, 
		Category:req.Category, 
		Nutrition:domain.Nutrition{
			Calories:req.Nutrition.Calories,
			TotalFat:req.Nutrition.TotalFat,
			Protein: req.Nutrition.Protein,
			Carbohydrate:req.Nutrition.Carbohydrate,
			Cholesterol:req.Nutrition.Cholesterol,
			Description:req.Nutrition.Description,
		},
		MainDish:req.MainDish,
		Sauce: req.Sauce,
		Directions: helper.ToDomainDirections(req.Directions), 
		IsLike: req.IsLike,
		Writer:req.Writer,
	}

	recipe, err = service.RecipeRepository.Create(ctx, tx, recipe)
	
	// Return response
    return helper.ToRecipeResponse(recipe),nil
}

func (service *RecipeServiceImpl) Update(ctx context.Context, id uint, req web.RecipeUpdateRequest) (web.RecipeResponse,error) {
	if err := service.Validate.Struct(req); err != nil {
		panic(err)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	recipe, err := service.RecipeRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(err)
	}

	recipe.Name = req.Name
	recipe.Description= req.Description 
	recipe.Img= req.Img 
	recipe.PrepTime=req.PrepTime 
	recipe.CookTime=req.CookTime 
	recipe.Category=req.Category 

	recipe.Nutrition.Calories = req.Nutrition.Calories
	recipe.Nutrition.TotalFat=     req.Nutrition.TotalFat
	recipe.Nutrition.Protein=      req.Nutrition.Protein
	recipe.Nutrition.Carbohydrate= req.Nutrition.Carbohydrate
	recipe.Nutrition.Cholesterol=  req.Nutrition.Cholesterol
	recipe.Nutrition.Description=  req.Nutrition.Description
	recipe.MainDish=req.MainDish
	recipe.Sauce= req.Sauce
	recipe.Directions= helper.ToDomainDirections(req.Directions) 
	recipe.IsLike= req.IsLike
	recipe.Writer=req.Writer
	//Save Update
	recipe, err = service.RecipeRepository.Update(ctx, tx, recipe)

	return helper.ToRecipeResponse(recipe),nil
}

func (service *RecipeServiceImpl) Patch(
    ctx context.Context,
    req web.RecipePatchRequest,
) (web.RecipeResponse, error) {

    tx, err := service.DB.Begin()
    if err != nil {
        return web.RecipeResponse{}, err
    }

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        } else {
            tx.Commit()
        }
    }()

    recipe, err := service.RecipeRepository.GetById(ctx, tx, *req.Id)
    if err != nil {
        return web.RecipeResponse{}, err
    }

    // contoh patch
    if req.Name != nil {
        recipe.Name = *req.Name
    }

    if req.Description != nil {
        recipe.Description = *req.Description
    }

	if req.Img != nil {
        recipe.Img = *req.Img
    }

	if req.PrepTime != nil {
        recipe.PrepTime = *req.PrepTime
    }

	if req.CookTime != nil {
        recipe.CookTime = *req.CookTime
    }

	if req.Category != nil {
        recipe.Category = *req.Category
    }
	if req.Nutrition.Calories != nil {
        recipe.Nutrition.Calories = *req.Nutrition.Calories
    }

	if req.Nutrition.TotalFat != nil {
        recipe.Nutrition.TotalFat = *req.Nutrition.TotalFat
    }

	if req.Nutrition.Protein != nil {
        recipe.Nutrition.Protein = *req.Nutrition.Protein
    }
	if req.Nutrition.Carbohydrate != nil {
        recipe.Nutrition.Carbohydrate = *req.Nutrition.Carbohydrate
    }

	if req.Nutrition.Cholesterol != nil {
        recipe.Nutrition.Cholesterol = *req.Nutrition.Cholesterol
    }
	if req.Nutrition.Description != nil {
        recipe.Nutrition.Description = *req.Nutrition.Description
    }

	if req.MainDish != nil {
        recipe.MainDish = *req.MainDish
    }

	if req.Sauce != nil {
        recipe.Sauce = *req.Sauce
    }

	if req.Directions != nil {
        recipe.Directions = helper.PatchToDirection(*req.Directions) 
    }

    recipe, err = service.RecipeRepository.Update(ctx, tx, recipe)
    if err != nil {
        return web.RecipeResponse{}, err
    }

    return helper.ToRecipeResponse(recipe), nil
}



func (service *RecipeServiceImpl) Delete(ctx context.Context, id uint) error {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	Recipe, err := service.RecipeRepository.GetById(ctx, tx, uint(id))
	if err != nil {
		return err
	}

	service.RecipeRepository.Delete(ctx, tx, Recipe.Id)
	return nil
}

func (service *RecipeServiceImpl) GetById(ctx context.Context, id uint) (web.RecipeResponse,error) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	Recipe, err := service.RecipeRepository.GetById(ctx, tx, id)
	if err != nil {
		panic(err.Error())
	}

	return helper.ToRecipeResponse(Recipe),nil
}

func (service *RecipeServiceImpl) GetAll(ctx context.Context) ([]web.RecipeResponse,error) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	recipe, err := service.RecipeRepository.GetAll(ctx, tx)
	if err != nil {
		panic(err.Error())
	}

	var RecipeResponses []web.RecipeResponse
	for _, Recipe := range recipe {
		RecipeResponses = append(RecipeResponses, helper.ToRecipeResponse(Recipe))
	}
	return RecipeResponses,nil
}
