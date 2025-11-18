package app

import (
	"github.com/julienschmidt/httprouter"
)

func NewRouter(recipeController recipeController ) *httprouter.Router {
	router := httprouter.New()
}