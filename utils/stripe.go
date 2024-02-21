package utils

import (
	"fmt"
)

// StripeClient representa el cliente para interactuar con la API de Stripe.
type StripeClient struct {
	APIKey string
}

// NewStripeClient crea una nueva instancia de StripeClient con la clave de API proporcionada.
func NewStripeClient(apiKey string) *StripeClient {
	return &StripeClient{APIKey: apiKey}
}

// Charge realiza un cargo a una tarjeta de crédito/débito utilizando la API de Stripe.
func (sc *StripeClient) Charge(amount float64, description string, userID int) error {
	// Aquí se implementaría la lógica para realizar un cargo utilizando la API de Stripe.
	// Por simplicidad, se mostrará un mensaje con los detalles del pago.
	fmt.Printf("Cargando %.2f a la tarjeta de usuario %d: %s\n", amount, userID, description)
	return nil
}

// RetrieveTransaction obtiene información sobre una transacción utilizando la API de Stripe.
func (sc *StripeClient) RetrieveTransaction(transactionID string) (map[string]interface{}, error) {
	// Aquí se implementaría la lógica para obtener información sobre una transacción utilizando la API de Stripe.
	// Por simplicidad, se devolverá un mapa con información de ejemplo.
	transaction := map[string]interface{}{
		"id":          transactionID,
		"amount":      100.00,
		"description": "Pago de prueba",
		"status":      "success",
	}
	return transaction, nil
}
