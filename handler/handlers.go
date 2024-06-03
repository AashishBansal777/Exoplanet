package handler

import (
	"ExoPlanet/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var exoplanets = make(map[int]model.ExoPlanet)
var idCounter = 1

func CreateExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet model.ExoPlanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if exoplanet.Type == "Terrestrial" && exoplanet.Mass == 0 {
		http.Error(w, "Please enter Mass because it is required for Terrestrial planets", http.StatusBadRequest)
		return
	}

	exoplanet.ID = idCounter
	idCounter++
	exoplanets[exoplanet.ID] = exoplanet

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func ListExoplanets(w http.ResponseWriter, r *http.Request) {
	exoplanetList := []model.ExoPlanet{}
	for _, exoplanet := range exoplanets {
		exoplanetList = append(exoplanetList, exoplanet)
	}
	json.NewEncoder(w).Encode(exoplanetList)
}

func GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID. Please enter valid id", http.StatusBadRequest)
		return
	}

	exoplanet, found := exoplanets[id]
	if !found {
		http.Error(w, "offoo Exoplanet not found...", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID. Please enter valid Id", http.StatusBadRequest)
		return
	}

	var exoplanet model.ExoPlanet
	err = json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if exoplanet.Type == "Terrestrial" && exoplanet.Mass == 0 {
		http.Error(w, "Please enter Mass because it is required for Terrestrial planets", http.StatusBadRequest)
		return
	}

	exoplanet.ID = id
	exoplanets[id] = exoplanet

	json.NewEncoder(w).Encode(exoplanet)
}

func DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, found := exoplanets[id]
	if !found {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	delete(exoplanets, id)
	w.WriteHeader(http.StatusNoContent)
}

func FuelEstimation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID.Please enter valid id", http.StatusBadRequest)
		return
	}

	exoplanet, found := exoplanets[id]
	if !found {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	crewCapacity, err := strconv.Atoi(r.URL.Query().Get("crew"))
	if err != nil || crewCapacity <= 0 {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}

	var gravity float64
	if exoplanet.Type == "GasGiant" {
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else if exoplanet.Type == "Terrestrial" {
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}

	fuel := (exoplanet.DistanceFromEarth / (gravity * gravity)) * float64(crewCapacity)

	json.NewEncoder(w).Encode(map[string]float64{"fuel_estimation": fuel})
}
