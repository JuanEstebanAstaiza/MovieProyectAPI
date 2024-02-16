package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/pquerna/otp/totp"
	"time"
)

func GenerateOTP(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}

func VerifyOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}

// GenerateUserID genera un ID único para los usuarios.
func GenerateUserID() (string, error) {
	// Generar un UUID (identificador único universal) v4
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", fmt.Errorf("error al generar UUID: %v", err)
	}
	// La versión 4 de UUID tiene los bits de la posición 6 y 7 establecidos en 01
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
