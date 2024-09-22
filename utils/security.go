package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"
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

func HashID(id uint64) string {
	id *= rand.Uint64()
	formated := fmt.Sprintf("%016x", id)
	hash := sha512.New()
	hash.Write([]byte(formated))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
