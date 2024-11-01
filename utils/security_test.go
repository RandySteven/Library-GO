package utils_test

import (
	"github.com/RandySteven/Library-GO/utils"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestHashPassword(t *testing.T) {
	t.Run("success to hash password", func(t *testing.T) {
		pass := "pass_1234"

		result := utils.HashPassword(pass)

		assert.NotEqual(t, pass, result)
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("success to compare hash passowrd", func(t *testing.T) {
		pass := "pass_1234"
		hashPass := utils.HashPassword(pass)

		result := utils.ComparePassword(pass, hashPass)

		assert.Equal(t, result, true)
	})

	t.Run("failed to compare password", func(t *testing.T) {
		pass := "pass_1234"
		hashPass := utils.HashPassword("pass_1235")

		result := utils.ComparePassword(pass, hashPass)

		assert.Equal(t, result, false)
	})
}

func TestHashID(t *testing.T) {
	t.Run("success to hash id", func(t *testing.T) {
		id := uint64(1)

		result := utils.HashID(id)

		assert.NotEqual(t, id, result)
	})
}
