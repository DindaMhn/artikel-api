package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRouter(t *testing.T) {
	t.Run("It should connected", func(t *testing.T) {
		result := CreateRouter()
		assert.NotNil(t, result)
	})
}
