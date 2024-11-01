package utils_test

import (
	"github.com/RandySteven/Library-GO/utils"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerateBorrowReference(t *testing.T) {
	t.Run("success to generate borrow reference id", func(t *testing.T) {
		length := uint64(16)

		referenceID := utils.GenerateBorrowReference(length)

		assert.Equal(t, int(length), len(referenceID))
	})

}
