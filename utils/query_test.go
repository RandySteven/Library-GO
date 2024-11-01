package utils

import (
	"github.com/RandySteven/Library-GO/queries"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryValidation(t *testing.T) {
	t.Run("test query validation with select", func(t *testing.T) {
		var query queries.GoQuery = "SELECT * FROM users"
		queryAction := "SELECT"

		err := QueryValidation(query, queryAction)

		assert.NoError(t, err)
	})

	t.Run("test query validation with insert", func(t *testing.T) {
		var query queries.GoQuery = `INSERT INTO users (name) VALUES (?)`
		queryAction := "INSERT"

		err := QueryValidation(query, queryAction)

		assert.NoError(t, err)
	})
}
