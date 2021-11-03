package router

import (
	"github.com/sfeir-cloud-iot/bicycle-api/handlers"

	"github.com/gorilla/mux"
)

func createBicycleRouter(router *mux.Router) {
	bicycleRouter := router.PathPrefix("/bicycle").Subrouter()

	bicycleRouter.
		HandleFunc("", handlers.GetAllBicyleData).
		Methods("GET")

	bicycleRouter.
		HandleFunc("/speed", handlers.GetBicycleCurrentSpeed).
		Methods("GET")

	bicycleRouter.
		HandleFunc("/config", handlers.GetConfig).
		Methods("GET")
}
