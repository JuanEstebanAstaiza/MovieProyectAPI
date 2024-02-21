package services

import (
	"database/sql"
	"errors"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/database"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

// AuthenticateUser verifica las credenciales del usuario y retorna true si son válidas, de lo contrario false.
func AuthenticateUser(credentials models.UserCredentials) (bool, error) {
	// Buscar al usuario por email en la base de datos
	user, err := database.FindUserByEmail(credentials.Email)
	if err != nil {
		return false, err
	}

	// Verificar si se encontró un usuario con el email dado
	if user == nil {
		return false, nil
	}

	// Obtener la contraseña encriptada del usuario desde la base de datos
	var storedPassword string
	err = utils.DB.QueryRow("SELECT password FROM users WHERE email = ?", credentials.Email).Scan(&storedPassword)
	if err != nil {
		return false, err
	}

	// Encriptar la contraseña proporcionada con MD5
	encryptedPassword := utils.EncryptPassword(credentials.Password)

	// Verificar si la contraseña encriptada coincide con la contraseña encriptada almacenada en la base de datos
	if storedPassword != encryptedPassword {
		return false, nil
	}

	// Las credenciales son válidas
	return true, nil
}

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

// RegisterUser registra un nuevo usuario en la base de datos.
func RegisterUser(user models.UserCredentials) error {
	// Generar un ID único para el usuario
	userID, err := utils.GenerateUserID()
	if err != nil {
		return err
	}

	// Encriptar la contraseña utilizando MD5
	encryptedPassword := utils.EncryptPassword(user.Password)

	// Insertar el usuario en la base de datos con el ID único generado
	_, err = utils.DB.Exec("INSERT INTO users (id, nickname, email, password) VALUES (?, ?, ?, ?)", userID, user.Nickname, user.Email, encryptedPassword)
	if err != nil {
		return err
	}

	return nil
}
func LoginUser(credentials models.UserCredentials) (*models.UserCredentials, error) {
	// Verificar las credenciales del usuario
	authenticated, err := AuthenticateUser(credentials)
	if err != nil {
		return nil, err
	}

	if !authenticated {
		// Las credenciales son inválidas
		return nil, nil
	}

	// Obtener la contraseña almacenada del usuario con el email proporcionado
	var storedPassword string
	err = utils.DB.QueryRow("SELECT password FROM users WHERE email = ?", credentials.Email).Scan(&storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No se encontró ningún usuario con el email dado
			return nil, nil
		}
		return nil, err
	}

	// Verificar si la contraseña proporcionada coincide con la contraseña almacenada
	if !utils.ComparePasswords(storedPassword, utils.EncryptPassword(credentials.Password)) {
		// Las contraseñas no coinciden
		return nil, nil
	}

	// Obtener el usuario con el email proporcionado
	var user models.UserCredentials
	err = utils.DB.QueryRow("SELECT id, nickname, email FROM users WHERE email = ?", credentials.Email).Scan(&user.ID, &user.Nickname, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No se encontró ningún usuario con el email dado
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
