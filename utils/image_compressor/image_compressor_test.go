package image_compressor

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidCompression(t *testing.T) {
	input := make([]byte, 0)

	output := CompressImage(input)

	assert.Equal(t, len(output), len(input))
}

func TestValidCompression(t *testing.T) {
	input, _ := ioutil.ReadFile("./test.jpeg")

	output := CompressImage(input)

	assert.Greater(t, len(input), len(output))
}
