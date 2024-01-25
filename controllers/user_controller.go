package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/services"
	"github.com/gorilla/mux"
)

func ModifyUserInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	var updatedUser models.UserProfile
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.ModifyUserInfo(userID, updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
