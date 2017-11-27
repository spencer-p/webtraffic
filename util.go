package main

import (
	"crypto/sha1"
	"encoding/base32"
	"strings"
)

func GetId(input string) string {
	// Calculate hash
	hash := sha1.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)

	// Convert to base32
	base32str := strings.ToLower(base32.HexEncoding.EncodeToString(hashBytes))

	return base32str
}
