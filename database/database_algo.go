package database

import (
	"database/sql"
	"errors"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
)

// FindUserByEmail busca un usuario por su dirección de correo electrónico en la base de datos y retorna un puntero al usuario si se encuentra, de lo contrario retorna nil.
func FindUserByEmail(email string) (*models.UserProfile, error) {
	// Abre la conexión a la base de datos (debes tener la conexión configurada previamente).
	db, err := sql.Open("mysql", "usuario:contraseña@tcp(dirección_base_de_datos:puerto)/nombre_base_de_datos")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Consulta a la base de datos para buscar al usuario por su email.
	query := "SELECT id, nickname, email FROM users WHERE email = ?"
	row := db.QueryRow(query, email)

	// Escanear el resultado de la consulta en una estructura de usuario.
	var user models.UserProfile
	err = row.Scan(&user.ID, &user.Nickname, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Si no se encontró ningún usuario con el email dado, devolvemos nil.
			return nil, nil
		}
		// Si hubo otro error durante la consulta, lo devolvemos.
		return nil, err
	}

	// Si se encontró un usuario, lo devolvemos.
	return &user, nil
}
