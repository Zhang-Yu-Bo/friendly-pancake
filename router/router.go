package router

import (
	"github.com/Zhang-Yu-Bo/friendly-pancake/controller"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	mRouter := mux.NewRouter()

	mRouter.HandleFunc("/", controller.HomePage).Methods("GET")

	return mRouter
}
