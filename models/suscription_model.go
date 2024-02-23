package models

type Subscription struct {
	ID          int     `json:"id"`
	UserID      string  `json:"user_id"`
	Plan        string  `json:"plan"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}

type SubscriptionStatus string

const (
	SubscriptionActive   SubscriptionStatus = "active"
	SubscriptionInactive SubscriptionStatus = "inactive"
)
