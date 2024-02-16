package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateUniqueMovieID genera un ID único de 8 dígitos para una película.
func GenerateUniqueMovieID() (string, error) {
	// Inicializar la semilla del generador de números aleatorios con el tiempo actual
	rand.Seed(time.Now().UnixNano())

	// Generar un número aleatorio de 8 dígitos
	randomID := rand.Intn(100000000) // Genera un número entre 0 y 99999999

	// Convertir el número generado a una cadena
	uniqueID := fmt.Sprintf("%08d", randomID)

	return uniqueID, nil
}
