package router

import (
	"net/http"

	"github.com/Zhang-Yu-Bo/friendly-pancake/controller"
	"github.com/Zhang-Yu-Bo/friendly-pancake/model/utility"
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
	mRouter.HandleFunc("/", limitRate(controller.HomePage, true, utility.Page)).Methods("GET")
	mRouter.HandleFunc("/favicon.ico", limitRate(controller.FaviconIco, true, utility.Json)).Methods("GET")

	mRouter.HandleFunc("/raw/code/image", limitRate(controller.ShowRawImage, true, utility.QRCode)).Methods("GET")
	mRouter.HandleFunc("/show/error/{message}", limitRate(controller.ShowMessagePage, true, utility.Page)).Methods("GET")

	mRouter.HandleFunc("/code", limitRate(controller.ShowCodeContent, true, utility.Json)).Methods("GET")
	mRouter.HandleFunc("/code", limitRate(controller.UploadCode, false, utility.Json)).Methods("POST")

	mRouter.HandleFunc("/test", controller.TestPage).Methods("GET")

	return mRouter
}
