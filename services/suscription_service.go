package services

import (
	"errors"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

// CreateSubscription crea una nueva suscripción para un usuario.
func CreateSubscription(subscription models.Subscription) error {
	_, err := utils.DB.Exec("INSERT INTO subscriptions (user_id, plan, price, description, status) VALUES (?, ?, ?, ?, ?)",
		subscription.UserID, subscription.Plan, subscription.Price, subscription.Description, subscription.Status)
	if err != nil {
		return err
	}

	return nil
}

// GetSubscriptionByUserID devuelve la suscripción activa de un usuario.
func GetSubscriptionByUserID(userID int) (models.Subscription, error) {
	var subscription models.Subscription

	err := utils.DB.QueryRow("SELECT id, user_id, plan, price, description, status FROM subscriptions WHERE user_id = ? AND status = ?",
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
