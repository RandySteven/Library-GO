package utils

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func ResizeImage(inputPath string, outputPath string, width uint, height uint) error {
	log.Println(inputPath)
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output image: %v", err)
	}
	defer outFile.Close()

	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(outFile, resizedImg, nil)
		if err != nil {
			return fmt.Errorf("failed to encode jpeg image: %v", err)
		}
	case "png":
		err = png.Encode(outFile, resizedImg)
		if err != nil {
			return fmt.Errorf("failed to encode png image: %v", err)
		}
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}

	return nil
}

func SeparateStringIntoUint64Arr(str string, sep string) []uint64 {
	strArr := strings.Split(str, sep)
	resArr := make([]uint64, len(strArr))
	for i, s := range strArr {
		resArr[i], _ = strconv.ParseUint(s, 10, 64)
	}
	return resArr
}

func GenerateBorrowReference(length uint64) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
