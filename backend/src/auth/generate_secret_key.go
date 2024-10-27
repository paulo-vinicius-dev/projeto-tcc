package auth

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateJWTSecretKey ...
func GenerateJWTSecretKey() (string, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
