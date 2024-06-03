package route

import (
	"ExoPlanet/handler"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/exoplanets", handler.CreateExoplanet).Methods("POST")
	router.HandleFunc("/exoplanets", handler.ListExoplanets).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", handler.GetExoplanetByID).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", handler.UpdateExoplanet).Methods("PUT")
	router.HandleFunc("/exoplanets/{id}", handler.DeleteExoplanet).Methods("DELETE")
	router.HandleFunc("/exoplanets/{id}/fuel", handler.FuelEstimation).Methods("GET")

	return router
}
