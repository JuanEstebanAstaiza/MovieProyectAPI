package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/services"
	"net/http"
	"strconv"
	"strings"
)

// ProcessPaymentHandler maneja la solicitud para procesar un pago.
func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	err = services.ProcessPayment(payment)
	if err != nil {
		http.Error(w, "Error al procesar el pago", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Pago procesado exitosamente")
}

// GetPaymentsByUserIDHandler maneja la solicitud para obtener todos los pagos realizados por un usuario.
func GetPaymentsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromRequest(r)

	payments, err := services.GetPaymentsByUserID(userID)
	if err != nil {
		http.Error(w, "Error al obtener los pagos del usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}

// UpdatePaymentStatusHandler maneja la solicitud para actualizar el estado de un pago.
func UpdatePaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	paymentID := getPaymentIDFromRequest(r)

	var updateData struct {
		Status models.PaymentStatus `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
		return
	}

	err = services.UpdatePaymentStatus(paymentID, updateData.Status)
	if err != nil {
		http.Error(w, "Error al actualizar el estado del pago", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Estado del pago actualizado exitosamente")
}

// GetTotalPaymentsByUserIDHandler maneja la solicitud para obtener la cantidad total de pagos realizados por un usuario.
func GetTotalPaymentsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromRequest(r)

	totalPayments, err := services.GetTotalPaymentsByUserID(userID)
	if err != nil {
		http.Error(w, "Error al obtener la cantidad total de pagos del usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"total_payments": %d}`, totalPayments)
}

// Función auxiliar para obtener el ID de usuario de la solicitud.
func getUserIDFromRequest(r *http.Request) int {
	// Supongamos que el ID de usuario se pasa como parte de la ruta en el formato "/api/user/{user_id}"
	parts := strings.Split(r.URL.Path, "/")
	userIDStr := parts[len(parts)-1]

	userID, _ := strconv.Atoi(userIDStr)
	return userID
}

// Función auxiliar para obtener el ID de pago de la solicitud.
func getPaymentIDFromRequest(r *http.Request) int {
	// Supongamos que el ID de pago se pasa como parte de la ruta en el formato "/api/payment/{payment_id}"
	parts := strings.Split(r.URL.Path, "/")
	paymentIDStr := parts[len(parts)-1]

	paymentID, _ := strconv.Atoi(paymentIDStr)
	return paymentID
}
