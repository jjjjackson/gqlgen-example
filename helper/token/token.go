package token

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateToken(size int) string {
	// TODO: make sure there are at least 1 number, 1 uppercast 1 and lowercaset in string
	result := make([]rune, size)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
