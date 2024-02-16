package database

import (
	"database/sql"
	"errors"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

func FindUserByEmail(email string) (*models.UserProfile, error) {
	// Consulta a la base de datos para buscar al usuario por su email.
	query := "SELECT id, nickname, email FROM users WHERE email = ?"
	row := utils.DB.QueryRow(query, email)

	// Escanear el resultado de la consulta en una estructura de usuario.
	var user models.UserProfile
	err := row.Scan(&user.ID, &user.Nickname, &user.Email)
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
