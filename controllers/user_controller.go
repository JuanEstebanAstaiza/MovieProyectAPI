package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/services"
	"github.com/gorilla/mux"
)

func ModifyUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	var updatedUser models.UserCredentials
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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	err = services.RegisterUser(user)
	if err != nil {
		http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parsear los datos del cuerpo de la solicitud
	var user models.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Obtener el usuario con las credenciales proporcionadas
	userInfo, err := services.LoginUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if userInfo == nil {
		// Las credenciales son inválidas
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Si las credenciales son válidas, enviar la información del usuario
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userInfo)
}
