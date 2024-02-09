package models

type UserProfile struct {
	ID       *int    `json:"id"`
	Nickname *string `json:"nickname"`
	Email    *string `json:"email"`
}

// UserCredentials representa las credenciales de inicio de sesión de un usuario
type UserCredentials struct {
	Password *string `json:"password"`
	Nickname *string `json:"nickname`
	Email    *string `json:"email`
	ID       *int    `json:"id"`
}
