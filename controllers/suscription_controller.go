package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/services"
	"net/http"
)

// CreateSubscriptionHandler maneja la solicitud para crear una nueva suscripción.
func CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var subscription models.Subscription
	err := json.NewDecoder(r.Body).Decode(&subscription)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	err = services.CreateSubscription(subscription)
	if err != nil {
		http.Error(w, "Error al crear la suscripción", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Suscripción creada exitosamente")
}

// GetSubscriptionByUserIDHandler maneja la solicitud para obtener la suscripción activa de un usuario.
func GetSubscriptionByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromRequest(r)

	subscription, err := services.GetSubscriptionByUserID(userID)
	if err != nil {
		http.Error(w, "Error al obtener la suscripción del usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscription)
}

// CancelSubscriptionHandler maneja la solicitud para cancelar la suscripción activa de un usuario.
func CancelSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromRequest(r)

	err := services.CancelSubscription(userID)
	if err != nil {
		http.Error(w, "Error al cancelar la suscripción del usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Suscripción cancelada exitosamente")
}
