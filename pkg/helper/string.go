package helper

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(length int) (string, error) {
	randomBytes := make([]byte, length/2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := hex.EncodeToString(randomBytes)
	return randomString[:length], nil
}
