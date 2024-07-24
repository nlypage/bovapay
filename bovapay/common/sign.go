package common

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateSignature(body string, apiKey string) string {
	hasher := sha1.New()
	hasher.Write([]byte(apiKey + body))
	sha := hasher.Sum(nil)
	return hex.EncodeToString(sha)
}
