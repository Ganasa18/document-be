package helper

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
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

func CamelToSnake(input string) string {
	regex := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := regex.ReplaceAllString(input, "${1}_${2}")
	return strings.ToLower(snake)
}
