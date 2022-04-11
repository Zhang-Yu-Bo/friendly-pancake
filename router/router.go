package router

import (
	"net/http"

	"github.com/Zhang-Yu-Bo/friendly-pancake/controller"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	mRouter := mux.NewRouter()

	// static file server
	mRouter.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static/")),
		),
	)

	mRouter.HandleFunc("/", controller.HomePage).Methods("GET")
	mRouter.HandleFunc("/raw/code/image", controller.ShowRawImage).Methods("GET")
	mRouter.HandleFunc("/test", controller.TestPage).Methods("GET")

	mRouter.HandleFunc("/code", controller.UploadCode).Methods("POST")

	return mRouter
}
