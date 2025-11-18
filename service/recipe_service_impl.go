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

func NewRecipeServiceImpl(recipeRepository repository.RecipeRepository, DB *sql.DB, validate *validator.Validate) *RecipeServiceImpl {
	return &RecipeServiceImpl{
		RecipeRepository: recipeRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *RecipeServiceImpl) Create(ctx context.Context, req web.RecipeCreateRequest) web.RecipeResponse {
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
		Directions: req.Directions, 
		IsLike: req.IsLike,
		Writer:req.Writer,
		CreateAt:   req.CreateAt,
	}

	recipe, err = service.RecipeRepository.Create(ctx, tx, recipe)
	
	// Return response
    return web.RecipeResponse{
        Id:   uint(recipe.Id),
        Name: recipe.Name,
        Description: recipe.Description,
        Img: recipe.Img,
        PrepTime: recipe.PrepTime,
        CookTime: recipe.CookTime,
        Category: recipe.Category,
        Nutrition: recipe.Nutrition,
        MainDish: recipe.MainDish,
        Sauce: recipe.Sauce,
        Directions: recipe.Directions,
        IsLike: recipe.IsLike,
        Writer: recipe.Writer,
        CreateAt: recipe.CreateAt,
    }
}

func (service *RecipeServiceImpl) Update(ctx context.Context, req web.RecipeUpdateRequest) web.RecipeResponse {
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

	recipe, err := service.RecipeRepository.GetById(ctx, tx, req.Id)
	if err != nil {
		panic(err)
	}

	recipe.Name = req.Name
	recipe.Description= req.Description 
	recipe.Img= req.Img 
	recipe.PrepTime=req.PrepTime 
	recipe.CookTime=req.CookTime 
	recipe.Category=req.Category 

	recipe.Nutrition = domain.Nutrition{
		Calories:     req.Nutrition.Calories,
		TotalFat:     req.Nutrition.TotalFat,
		Protein:      req.Nutrition.Protein,
		Carbohydrate: req.Nutrition.Carbohydrate,
		Cholesterol:  req.Nutrition.Cholesterol,
		Description:  req.Nutrition.Description,
	}
	recipe.MainDish=req.MainDish
	recipe.Sauce= req.Sauce
	recipe.Directions= req.Directions 
	recipe.IsLike= req.IsLike
	recipe.Writer=req.Writer
	//Save Update
	recipe, err = service.RecipeRepository.Update(ctx, tx, recipe)

	return web.RecipeResponse{
        Id:   uint(recipe.Id),
        Name: recipe.Name,
        Description: recipe.Description,
        Img: recipe.Img,
        PrepTime: recipe.PrepTime,
        CookTime: recipe.CookTime,
        Category: recipe.Category,
        Nutrition: recipe.Nutrition,
        MainDish: recipe.MainDish,
        Sauce: recipe.Sauce,
        Directions: recipe.Directions,
        IsLike: recipe.IsLike,
        Writer: recipe.Writer,
    }
}

func (service *RecipeServiceImpl) Delete(ctx context.Context, RecipeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	Recipe, err := service.RecipeRepository.GetById(ctx, tx, RecipeId)
	if err != nil {
		panic(exception.NotFoundErrHandler(err.Error()))
	}

	service.RecipeRepository.Delete(ctx, tx, Recipe)
}

func (service *RecipeServiceImpl) GetById(ctx context.Context, RecipeId int) web.RecipeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	Recipe, err := service.RecipeRepository.GetById(ctx, tx, RecipeId)
	if err != nil {
		panic(exception.NotFoundErrHandler(err.Error()))
	}

	return helper.ToRecipeResponse(Recipe)
}

func (service *RecipeServiceImpl) GetAll(ctx context.Context) []web.RecipeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	categories := service.RecipeRepository.GetAll(ctx, tx)

	var RecipeResponses []web.RecipeResponse
	for _, Recipe := range categories {
		RecipeResponses = append(RecipeResponses, helper.ToRecipeResponse(Recipe))
	}
	return RecipeResponses
}
