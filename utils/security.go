package utils

import (
	"crypto/sha512"
	"encoding/base64"
)

func HashPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	pass := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return pass
}

func ComparePassword(requestPassword, existPassword string) bool {
	hash := sha512.New()
	hash.Write([]byte(requestPassword))
	pass := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return pass == existPassword
}
