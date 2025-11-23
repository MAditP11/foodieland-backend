package controller

import (
	"encoding/json"
	"foodieland/helper"
	"foodieland/model/web"
	"foodieland/service"
	"net/http"

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
	recipePatchReq := web.RecipePatchRequest{}
	if err := json.NewDecoder(r.Body).Decode(&recipePatchReq); err != nil {
		helper.WriteJSON(w,http.StatusBadRequest,web.ErrorResponse{
            Message: "Invalid JSON: " + err.Error(),
        })
        return
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
