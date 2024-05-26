package services

import (
	"errors"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

// CreateSubscription crea una nueva suscripción para un usuario.
func CreateSubscription(subscription models.Subscription) (models.Subscription, error) {
	// Verificar si ya hay una suscripción activa para este usuario
	var activeSubCount int
	err := utils.DB.QueryRow("SELECT COUNT(*) FROM subscriptions WHERE user_id = ? AND status = ?", subscription.UserID, models.SubscriptionActive).Scan(&activeSubCount)
	if err != nil {
		return models.Subscription{}, err
	}

	if activeSubCount > 0 {
		// Ya hay una suscripción activa para este usuario, devolver un error
		return models.Subscription{}, errors.New("Ya existe una suscripción activa para este usuario")
	}

	// Insertar la nueva suscripción
	result, err := utils.DB.Exec("INSERT INTO subscriptions (user_id, plan, price, description, status) VALUES (?, ?, ?, ?, ?)",
		subscription.UserID, subscription.Plan, subscription.Price, subscription.Description, models.SubscriptionActive)
	if err != nil {
		return models.Subscription{}, err
	}

	// Obtener el ID de la nueva suscripción
	newSubscriptionID, err := result.LastInsertId()
	if err != nil {
		return models.Subscription{}, err
	}

	// Obtener los detalles completos de la nueva suscripción
	var newSubscription models.Subscription
	err = utils.DB.QueryRow("SELECT id, user_id, plan, price, description, status FROM subscriptions WHERE id = ?", newSubscriptionID).Scan(&newSubscription.ID, &newSubscription.UserID, &newSubscription.Plan, &newSubscription.Price, &newSubscription.Description, &newSubscription.Status)
	if err != nil {
		return models.Subscription{}, err
	}

	return newSubscription, nil
}

// GetSubscriptionByUserID devuelve la suscripción activa de un usuario.
func GetSubscriptionByUserID(userID string) (models.Subscription, error) {
	var subscription models.Subscription
	err := utils.DB.QueryRow("SELECT * FROM subscriptions WHERE user_id = ? AND status = ?",
		userID, models.SubscriptionActive).Scan(&subscription.ID, &subscription.UserID, &subscription.Plan,
		&subscription.Price, &subscription.Description, &subscription.Status)
	if err != nil {
		return models.Subscription{}, err
	}

	return subscription, nil
}

// CancelSubscription cancela la suscripción activa de un usuario cambiando su estado a inactivo.
func CancelSubscription(userID string) error {
	result, err := utils.DB.Exec("UPDATE subscriptions SET status = ? WHERE user_id = ? AND status = ?",
		models.SubscriptionInactive, userID, models.SubscriptionActive)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se encontró ninguna suscripción activa para el usuario")
	}

	return nil
}
