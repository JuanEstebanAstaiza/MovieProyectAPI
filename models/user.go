package models

type UserProfile struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

// UserCredentials representa las credenciales de inicio de sesión de un usuario
type UserCredentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
