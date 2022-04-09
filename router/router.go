package router

import (
	"github.com/Zhang-Yu-Bo/friendly-pancake/controller"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	mRouter := mux.NewRouter()

	// static file server
	// mRouter.PathPrefix("/static/").Handler(
	// 	http.StripPrefix(
	// 		"/static/",
	// 		http.FileServer(http.Dir("")),
	// 	),
	// )

	mRouter.HandleFunc("/", controller.HomePage).Methods("GET")
	mRouter.HandleFunc("/raw/image", controller.RawImage).Methods("GET")

	return mRouter
}
