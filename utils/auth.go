package utils

import (
	"github.com/pquerna/otp/totp"
	"time"
)

func GenerateOTP(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}

func VerifyOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}
