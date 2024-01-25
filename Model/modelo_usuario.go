package Model

type User struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}
