package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
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

func CompareID(requestId, existId string) bool {
	hash := sha512.New()
	hash.Write([]byte(requestId))
	pass := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return pass == existId
}

func RenameFileWithDateAndUUID(fileName string) string {
	currentDate := time.Now().Format("20060102")

	uniqueID := uuid.New().String()

	extension := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, extension)

	newFileName := fmt.Sprintf("%s_%s_%s%s", baseName, currentDate, uniqueID, extension)

	return newFileName
}
