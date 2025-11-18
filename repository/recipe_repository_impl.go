package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"foodieland/helper"
	"foodieland/model/domain"
)

type RecipeRepositoryImpl struct{

}

func NewRecipeRepository() *RecipeRepositoryImpl {
	return &RecipeRepositoryImpl{}
}

func (repository *RecipeRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, recipe domain.Recipe) (domain.Recipe, error) {
	SQL := `insert into recipe (
        name, description, img, prep_time, cook_time, category,
        calories, total_fat, protein, carbohydrate, cholesterol,
        description_nutrition, main_dish, sauce, directions,
        islike, writer, create_at
    ) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	mainDishJson, _ := json.Marshal(recipe.MainDish)
	sauceJson, _ := json.Marshal(recipe.Sauce)
	directionJson, _ := json.Marshal(recipe.Directions)
	
	result, err := tx.ExecContext(ctx, SQL,
		recipe.Name, recipe.Description, recipe.Img,
		recipe.PrepTime, recipe.CookTime, recipe.Category,
		recipe.Nutrition.Calories, recipe.Nutrition.TotalFat, recipe.Nutrition.Protein,
		recipe.Nutrition.Carbohydrate, recipe.Nutrition.Cholesterol, recipe.Nutrition.Description,
		string(mainDishJson), string(sauceJson), string(directionJson),
		recipe.IsLike, recipe.Writer, recipe.CreateAt,
	)

	if err != nil {
		return recipe, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return recipe, err
	}

	recipe.Id = int(id)
	return recipe, nil
}


func (repository *RecipeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, recipe domain.Recipe) (domain.Recipe, error) {
	SQL := `
        UPDATE recipe SET 
            name = ?, description = ?, img = ?, prep_time = ?, cook_time = ?, 
            category = ?, calories = ?, total_fat = ?, protein = ?, carbohydrate = ?, 
            cholesterol = ?, description_nutrition = ?, main_dish = ?, sauce = ?, 
            directions = ?, islike = ?, writer = ?
        WHERE id = ?
    `

	mainDishJson, _ := json.Marshal(recipe.MainDish)
	sauceJson, _ := json.Marshal(recipe.Sauce)
	directionJson, _ := json.Marshal(recipe.Directions)
	
	_, err := tx.ExecContext(
		ctx,
		SQL,
		recipe.Name,
		recipe.Description,
		recipe.Img,
		recipe.PrepTime,
		recipe.CookTime,
		recipe.Category,
		recipe.Nutrition.Calories,
		recipe.Nutrition.TotalFat,
		recipe.Nutrition.Protein,
		recipe.Nutrition.Carbohydrate,
		recipe.Nutrition.Cholesterol,
		recipe.Nutrition.Description,
		string(mainDishJson),
		string(sauceJson),
		string(directionJson),
		recipe.IsLike,
		recipe.Writer,
		recipe.Id,
	)

	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (repo *RecipeRepositoryImpl) Patch(ctx context.Context, tx *sql.Tx, recipeId int, patch domain.RecipePatch) error {
	query := "UPDATE recipe SET "
	params := []interface{}{}
	first := true

	if patch.Name != nil {
		query += addComma(first) + "name = ?"
		params = append(params, *patch.Name)
		first = false
	}

	if patch.Description != nil {
		query += addComma(first) + "description = ?"
		params = append(params, *patch.Description)
		first = false
	}

	if patch.Img != nil {
		query += addComma(first) + "img = ?"
		params = append(params, *patch.Img)
		first = false
	}

	if patch.PrepTime != nil {
		query += addComma(first) + "prep_time = ?"
		params = append(params, *patch.PrepTime)
		first = false
	}

	if patch.CookTime != nil {
		query += addComma(first) + "cook_time = ?"
		params = append(params, *patch.CookTime)
		first = false
	}

	if patch.Category != nil {
		query += addComma(first) + "category = ?"
		params = append(params, *patch.Category)
		first = false
	}

	if patch.Nutrition != nil {
		if patch.Nutrition.Calories != nil {
			query += addComma(first) + "calories = ?"
			params = append(params, *patch.Nutrition.Calories)
			first = false
		}
		if patch.Nutrition.TotalFat != nil {
			query += addComma(first) + "total_fat = ?"
			params = append(params, *patch.Nutrition.TotalFat)
			first = false
		}

		if patch.Nutrition.Carbohydrate != nil {
			query += addComma(first) + "carbohydrate = ?"
			params = append(params, *patch.Nutrition.Carbohydrate)
			first = false
		}

		if patch.Nutrition.Cholesterol != nil {
			query += addComma(first) + "cholesterol = ?"
			params = append(params, *patch.Nutrition.Cholesterol)
			first = false
		}

		if patch.Nutrition.Description != nil {
			query += addComma(first) + "description_nutrition = ?"
			params = append(params, *patch.Nutrition.Description)
			first = false
		}
	}

	if patch.MainDish != nil {
		jsonData,_ := json.Marshal(*patch.MainDish)
		query += addComma(first) + "main_dish = ?"
		params = append(params, string(jsonData))
		first = false
	}

	if patch.Sauce != nil {
		jsonData,_ := json.Marshal(*patch.Sauce)
		query += addComma(first) + "sauce = ?"
		params = append(params, string(jsonData))
		first = false
	}

	if patch.Directions != nil {
		jsonData,_ := json.Marshal(*patch.Directions)
		query += addComma(first) + "directions = ?"
		params = append(params, string(jsonData))
		first = false
	}

	if patch.IsLike != nil {
		query += addComma(first) + "is_like = ?"
		params = append(params, *patch.IsLike)
		first = false
	}

	if patch.Writer != nil {
		query += addComma(first) + "writer = ?"
		params = append(params, *patch.Writer)
		first = false
	}

	query += " WHERE id = ?"
	params = append(params, recipeId)

	_, err := tx.ExecContext(ctx, query, params...)
	return err
}

func addComma(first bool) string {
	if first {
		return ""
	}
	return ", "
}



func (repo *RecipeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, recipeId int) error {
	SQL := "delete from recipe where id = ?"
	_, err := tx.ExecContext(ctx, SQL, recipeId)
	return err
}

func (repo *RecipeRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, recipeId int) (domain.Recipe, error) {
	SQL := "select id,name,description,img,prep_time,cook_time,category,calories,total_fat,protein,carbohydrate,cholesterol,description_nutrition,main_dish,sauce,directions,islike,writer,create_at from recipe where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, recipeId)
	if err != nil {
		return domain.Recipe{}, err
	}
	defer rows.Close()
	recipe := domain.Recipe{}
	if rows.Next() {
		var mainDishStr string
		var sauceStr string
		var directionStr string
		if err := rows.Scan(&recipe.Id,&recipe.Name,&recipe.Description,&recipe.Img,&recipe.PrepTime,&recipe.CookTime,&recipe.Category,&recipe.Nutrition.Calories,&recipe.Nutrition.TotalFat,&recipe.Nutrition.Protein,&recipe.Nutrition.Carbohydrate,&recipe.Nutrition.Cholesterol,&recipe.Nutrition.Description,&mainDishStr,&sauceStr,&directionStr,&recipe.IsLike,&recipe.Writer,&recipe.CreateAt); err != nil{
			return recipe,err
		}
		if err := json.Unmarshal([]byte(mainDishStr), &recipe.MainDish); err != nil {
			return recipe,err
		}
		if err:= json.Unmarshal([]byte(sauceStr), &recipe.Sauce); err != nil {
			return recipe,err
		}
		if err := json.Unmarshal([]byte(directionStr), &recipe.Directions); err != nil {
			return recipe,err
		}
		
		return recipe, nil
	} 
		return recipe, errors.New("recipe is not found")
	
}

func (repo *RecipeRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Recipe,error) {
	SQL := "select id,name,description,img,prep_time,cook_time,category,calories,total_fat,protein,carbohydrate,cholesterol,description_nutrition,main_dish,sauce,directions,islike,writer,create_at from recipe"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)

	defer rows.Close()
	var recipes []domain.Recipe
	for rows.Next() {
		recipe := domain.Recipe{}
		var mainDishStr,sauceStr,directionStr string
		
		if err := rows.Scan(&recipe.Id,&recipe.Name,&recipe.Description,&recipe.Img,&recipe.PrepTime,&recipe.CookTime,&recipe.Category,&recipe.Nutrition.Calories,&recipe.Nutrition.TotalFat,&recipe.Nutrition.Protein,&recipe.Nutrition.Carbohydrate,&recipe.Nutrition.Cholesterol,&recipe.Nutrition.Description,&mainDishStr,&sauceStr,&directionStr,&recipe.IsLike,&recipe.Writer,&recipe.CreateAt); err != nil {
			return nil, err
		}
		if err:=helper.ScanJson(mainDishStr,&recipe.MainDish); err != nil {
			return nil,err
		}
		if err := helper.ScanJson(sauceStr,&recipe.Sauce); err != nil {
			return nil,err
		}
		if err := helper.ScanJson(directionStr,&recipe.Directions); err != nil {
			return nil,err
		}
		recipes = append(recipes, recipe)
	}

	return recipes,nil
}