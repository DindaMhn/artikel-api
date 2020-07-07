package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	t.Run("It should connected", func(t *testing.T) {
		result := Connection()
		assert.NotNil(t, result)
	})
}
