package router

import (
	"github.com/Vansh0100/E-Library/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router:=mux.NewRouter()
	router.HandleFunc("/",controller.HomePage).Methods("GET")
	router.HandleFunc("/login",controller.VerifyLogin).Methods("GET")
	router.HandleFunc("/signup",controller.RegisterUser).Methods("POST")

	return router
}