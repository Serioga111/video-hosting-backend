package services

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomToken() (string, error) {
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(tokenBytes), nil
}
