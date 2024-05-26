package services

import (
	"errors"
	"fmt"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
	"time"
)

// ProcessPayment simula el procesamiento de un pago utilizando la API de Stripe.
func ProcessPayment(payment models.Payment) error {
	// Simular el procesamiento del pago con Stripe
	fmt.Println("Procesando pago con Stripe...")
	// Aquí se podría llamar a la API de Stripe para procesar el pago
	// Stripe.Charge(payment.Amount, payment.Description, payment.UserID)
	time.Sleep(time.Second * 2) // Simular una demora en el procesamiento del pago

	// Por ahora, simplemente insertaremos el pago en la base de datos
	_, err := utils.DB.Exec("INSERT INTO payments (user_id, amount, description, status) VALUES (?, ?, ?, ?)",
		payment.UserID, payment.Amount, payment.Description, models.PaymentSuccess)
	if err != nil {
		return err
	}

	fmt.Println("Pago procesado exitosamente con Stripe.")
	return nil
}

// GetPaymentsByUserID devuelve todos los pagos realizados por un usuario específico.
func GetPaymentsByUserID(userID string) ([]models.Payment, error) {
	var payments []models.Payment

	rows, err := utils.DB.Query("SELECT id, user_id, amount, description, status FROM payments WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment models.Payment
		if err := rows.Scan(&payment.ID, &payment.UserID, &payment.Amount, &payment.Description, &payment.Status); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

// UpdatePaymentStatus actualiza el estado de un pago en la base de datos.
func UpdatePaymentStatus(paymentID int) error {
	result, err := utils.DB.Exec("UPDATE payments SET status = ? WHERE id = ?", models.PaymentFailed, paymentID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se encontró ningún pago con el ID proporcionado")
	}

	return nil
}

// GetTotalPaymentsByUserID devuelve la cantidad total de pagos realizados por un usuario específico.
func GetTotalPaymentsByUserID(userID string) (int, error) {
	var totalPayments int

	err := utils.DB.QueryRow("SELECT COUNT(id) FROM payments WHERE user_id = ?", userID).Scan(&totalPayments)
	if err != nil {
		return 0, err
	}

	return totalPayments, nil
}
