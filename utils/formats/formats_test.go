package formats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeValid_ValidateFormat(t *testing.T) {
	input := JPEG

	err := ValidateFormat(input)

	assert.Nil(t, err)
}

func TestShouldBeInvalid_ValidateFormat(t *testing.T) {
	input := "pdf"

	err := ValidateFormat(input)

	assert.Error(t, err, "invalid")
}
