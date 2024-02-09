package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DB representa una conexión de base de datos.
var DB *sql.DB

// Config contiene la configuración de la base de datos.
var Config = struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}{
	Username: "root",
	Password: "root",
	Host:     "localhost",
	Port:     3306, // Puerto predeterminado de MySQL
	Database: "awsdeploy",
}

// InitDB inicializa una conexión a la base de datos MySQL.
func InitDB() error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Config.Username, Config.Password, Config.Host, Config.Port, Config.Database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	// Intentar abrir una conexión con la base de datos para verificar si la configuración es válida
	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}
