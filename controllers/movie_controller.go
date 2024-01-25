package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanEstebanAstaiza/MovieProyectAPI/services"
	"github.com/gorilla/mux"
)

func GetMovieDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["movie_id"]

	movieDetails, err := services.GetMovieDetails(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Incrementar el contador de visualizaciones (implementación en services/movie_service.go)
	err = services.IncrementViewCount(movieID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieDetails)
}

func GetMostViewedMovies(w http.ResponseWriter, r *http.Request) {
	// Implementar lógica para obtener las n películas más visualizadas desde services/movie_service.go
}
