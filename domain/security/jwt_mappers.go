package security

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateCSRF(n int) (string, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return "error generete csrf", err
	}

	return base64.URLEncoding.EncodeToString(b), err
}
