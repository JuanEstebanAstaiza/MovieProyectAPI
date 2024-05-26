package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/services"
	"github.com/gorilla/mux"
	"net/http"
)

// CreateSubscriptionHandler maneja la solicitud para crear una nueva suscripción.
func CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var subscription models.Subscription
	err := json.NewDecoder(r.Body).Decode(&subscription)
	fmt.Println(subscription)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	sub := models.Subscription{}
	sub, err = services.CreateSubscription(subscription)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear la suscripción: %e", err), http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(sub); err != nil {
		http.Error(w, fmt.Sprintf("Error al crear la suscripción: %e", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetSubscriptionByUserIDHandler maneja la solicitud para obtener la suscripción activa de un usuario.
func GetSubscriptionByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	subscription, err := services.GetSubscriptionByUserID(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener la suscripción del usuario: %e", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscription)
}

// CancelSubscriptionHandler maneja la solicitud para cancelar la suscripción activa de un usuario.
func CancelSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	err := services.CancelSubscription(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al cancelar la suscripción del usuario: %e", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Suscripción cancelada exitosamente")
}
