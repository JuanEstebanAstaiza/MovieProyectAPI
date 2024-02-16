package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB representa una conexión de base de datos.
var DB *sql.DB

// InitDB inicializa una conexión a la base de datos MySQL.
func InitDB() error {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Obtener configuraciones de la base de datos desde las variables de entorno
	dbHost := os.Getenv("DB_host")
	dbPortStr := os.Getenv("DB_port")
	dbUser := os.Getenv("DB_user")
	dbPassword := os.Getenv("DB_password")
	dbName := os.Getenv("DB_name")

	// Convertir el puerto de cadena a entero
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return fmt.Errorf("error al convertir el puerto a entero: %v", err)
	}

	// Crear cadena de conexión
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar a la base de datos
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión a la base de datos: %v", err)
	}

	// Intentar abrir una conexión con la base de datos para verificar si la configuración es válida
	if err := db.Ping(); err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	DB = db
	return nil
}
