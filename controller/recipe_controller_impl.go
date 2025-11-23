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

	_ = json.NewDecoder(r.Body).Decode(&recipeCreateReq)

	recipeResponse, _ := controller.RecipeService.Create(r.Context(), recipeCreateReq)

	helper.WriteJSON(w, 200, recipeResponse)
}

func (controller RecipeControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := helper.ParseID(params)
	recipeUpdateReq := web.RecipeUpdateRequest{}
	_ = json.NewDecoder(r.Body).Decode(&recipeUpdateReq)
	recipeUpdateReq.Id = id

	recipeResponse,_ := controller.RecipeService.Update(r.Context(), id, recipeUpdateReq)
	helper.WriteJSON(w,200,recipeResponse)
}

func (controller RecipeControllerImpl) Patch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := helper.ParseID(params)
	recipePatchReq := web.RecipePatchRequest{}
	_ = json.NewDecoder(r.Body).Decode(&recipePatchReq)
	recipePatchReq.Id = id

	recipeResponse,_ := controller.RecipeService.Patch(r.Context(), id, recipePatchReq)
	helper.WriteJSON(w,200,recipeResponse)
}

func (controller RecipeControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := helper.ParseID(params)

	controller.RecipeService.Delete(r.Context(), id)
	helper.WriteJSON(w,200,nil)
}

func (controller RecipeControllerImpl) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := helper.ParseID(params)
	recipeResponse,_ := controller.RecipeService.GetById(r.Context(), id)
	helper.WriteJSON(w,200,recipeResponse)
}

func (controller RecipeControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	recipeResponses,_ := controller.RecipeService.GetAll(r.Context())
	helper.WriteJSON(w,200,recipeResponses)
}
