package utils

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"reflect"
	"strings"
)

func ToPayload(data interface{}) []byte {
	payload, _ := json.Marshal(data)

	return payload
}

func ExtractWords[T comparable](text string, delimiter string) []T {
	words := strings.Fields(text)
	var result []T
	for _, word := range words {
		if strings.Contains(word, delimiter) {
			result = append(result, reflect.ValueOf(word[len(delimiter):]).Interface().(T))
		}
	}
	return result
}

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	code := strings.ToUpper(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b))
	return code[:length], nil
}
