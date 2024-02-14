package services

import (
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

func ModifyUserInfo(userID string, updatedUser models.UserCredentials) error {
	err := modifyUserInfoInDB(userID, updatedUser)
	return err
}

func modifyUserInfoInDB(userID string, updatedUser models.UserCredentials) error {
	_, err := utils.DB.Exec("UPDATE users SET email = ?, password = ? WHERE id = ?",
		updatedUser.Email, updatedUser.Password, userID)
	if err != nil {
		return err
	}
	return nil
}

func RegisterUser(user models.UserCredentials) error {
	// Insertar el usuario en la base de datos
	_, err := utils.DB.Exec("INSERT INTO users (nickname, email, password) VALUES (?, ?, ?)", user.Nickname, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(credentials models.UserCredentials) (*models.UserCredentials, error) {
	// Obtener el usuario con el email proporcionado
	var user models.UserCredentials
	err := utils.DB.QueryRow("SELECT id, nickname, email, password FROM users WHERE email = ? AND password = ?", credentials.Email, credentials.Password).Scan(&user.ID, &user.Nickname, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
