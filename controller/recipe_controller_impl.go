package controller

import (
	"encoding/json"
	"foodieland/helper"
	"foodieland/model/web"
	"foodieland/service"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type RecipeControllerImpl struct {
	RecipeService service.RecipeService
}

func NewRecipeControllerImpl (recipeService service.RecipeService) RecipeController {
	return &RecipeControllerImpl{
		RecipeService: recipeService,
	}
}

func (controller RecipeControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	recipeCreateReq := web.RecipeCreateRequest{}

	err := json.NewDecoder(r.Body).Decode(&recipeCreateReq)
	if err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: "Invalid JSON: " + err.Error(),
        })
        return
	}

	recipeResponse, err := controller.RecipeService.Create(r.Context(), recipeCreateReq)
	if err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: err.Error(),
        })
        return
	}

	helper.WriteJSON(w, http.StatusOK, recipeResponse)
}

func (controller RecipeControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := helper.ParseID(params)
	recipeUpdateReq := web.RecipeUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(&recipeUpdateReq)
	if err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: "Invalid JSON: " + err.Error(),
        })
        return
	}
	recipeUpdateReq.Id = id

	recipeResponse,err := controller.RecipeService.Update(r.Context(), id, recipeUpdateReq)
	if err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: err.Error(),
        })
        return
	}
	helper.WriteJSON(w,http.StatusOK,recipeResponse)
}

func (controller RecipeControllerImpl) Patch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := helper.ParseID(params)

	// Parse form-data, maksimal 10 MB
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        helper.WriteJSON(w, http.StatusBadRequest, web.ErrorResponse{Message: "Invalid form data"})
        return
    }

	recipePatchReq := web.RecipePatchRequest{}

	// Ambil file
    file, header, err := r.FormFile("image")
    if err == nil { // file ada
        defer file.Close()
        dst := "./uploads/" + header.Filename
        out, err := os.Create(dst)
        if err != nil {
            helper.WriteJSON(w, http.StatusInternalServerError, web.ErrorResponse{Message: "Cannot save file"})
            return
        }
        defer out.Close()
        io.Copy(out, file)

        // Update field image di request struct
        imagePath := dst
        recipePatchReq.Image = &imagePath
    }

    // Ambil field lain dari form-data
    if title := r.FormValue("title"); title != "" {
        recipePatchReq.Title = &title
    }

	if description := r.FormValue("description"); description != "" {
        recipePatchReq.Description = &description
    }

	if prep_time := r.FormValue("prep_time"); prep_time != "" {
        recipePatchReq.PrepTime = &prep_time
    }

	if cook_time := r.FormValue("cook_time"); cook_time != "" {
        recipePatchReq.CookTime = &cook_time
    }

	if category := r.FormValue("category"); category != "" {
        recipePatchReq.Category = &category
    }

	if nutritionStr := r.FormValue("nutrition"); nutritionStr != "" {
		var nutrition web.NutritionPatchRequest
		if err := json.Unmarshal([]byte(nutritionStr), &nutrition); err != nil {
			helper.WriteJSON(w, http.StatusBadRequest, web.ErrorResponse{Message: "Invalid nutrition JSON"})
			return
		}
		recipePatchReq.Nutrition = &nutrition
	}

	if main_dishStr := r.FormValue("main_dish"); main_dishStr != "" {
		var main_dish []string
		if err := json.Unmarshal([]byte(main_dishStr), &main_dish); err != nil {
			helper.WriteJSON(w, http.StatusBadRequest, web.ErrorResponse{Message: "Invalid main_dish JSON"})
			return
		}
		recipePatchReq.MainDish = &main_dish
	}

	if sauceStr := r.FormValue("sauce"); sauceStr != "" {
		var sauce []string
		if err := json.Unmarshal([]byte(sauceStr), &sauce); err != nil {
			helper.WriteJSON(w, http.StatusBadRequest, web.ErrorResponse{Message: "Invalid sauce JSON"})
			return
		}
		recipePatchReq.Sauce = &sauce
	}

	if directionsStr := r.FormValue("directions"); directionsStr != "" {
		var directions []web.DirectionPatchRequest
		if err := json.Unmarshal([]byte(directionsStr), &directions); err != nil {
			helper.WriteJSON(w, http.StatusBadRequest, web.ErrorResponse{Message: "Invalid directions JSON"})
			return
		}
		recipePatchReq.Directions = &directions
	}
	

	if is_likeStr := r.FormValue("is_like"); is_likeStr != "" {
		is_like := is_likeStr == "true"
        recipePatchReq.IsLike = &is_like
    }

	if writer := r.FormValue("writer"); writer != "" {
        recipePatchReq.Writer = &writer
    }

	recipePatchReq.Id = &id

	recipeResponse,err := controller.RecipeService.Patch(r.Context(), recipePatchReq)
	if err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: err.Error(),
        })
        return
	}
	helper.WriteJSON(w,http.StatusOK,recipeResponse)
}

func (controller RecipeControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := helper.ParseID(params)

	if err :=controller.RecipeService.Delete(r.Context(), id); err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: err.Error(),
        })
        return
	}
	helper.WriteJSON(w,http.StatusOK,nil)
}

func (controller RecipeControllerImpl) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := helper.ParseID(params)
	recipeResponse,err := controller.RecipeService.GetById(r.Context(), id)
	if err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: err.Error(),
        })
        return
	}
	helper.WriteJSON(w,http.StatusOK,recipeResponse)
}

func (controller RecipeControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	recipeResponses,err := controller.RecipeService.GetAll(r.Context())
	if err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: err.Error(),
        })
        return
	}
	helper.WriteJSON(w,http.StatusOK,recipeResponses)
}
