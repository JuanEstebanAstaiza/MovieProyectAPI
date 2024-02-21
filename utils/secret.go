package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptPassword encripta una contraseña utilizando el algoritmo MD5
func EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	encryptedPassword := hex.EncodeToString(hash.Sum(nil))
	return encryptedPassword
}

func ComparePasswords(storedHash string, providedHash string) bool {
	return storedHash == providedHash
}
