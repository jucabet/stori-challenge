package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {

	t.Run("generate code", func(t *testing.T) {
		code := "1234"
		tempGenerateCode := GenerateRandomCode(4)
		assert.NotEmpty(t, code, tempGenerateCode)
		assert.NotNil(t, tempGenerateCode)

	})
}
