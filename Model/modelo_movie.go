package Model

type Movie struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Overview         string `json:"overview"`
	ReleaseDate      string `json:"release_date"`
	OriginalLanguage string `json:"original_language"`
}

type MovieComment struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	MovieID    int    `json:"movie_id"`
	Comment    string `json:"comment"`
	CreateTime string `json:"create_time"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
