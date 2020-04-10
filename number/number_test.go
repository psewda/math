package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, Sum(5, 5))
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, 20, Multiply(5, 4))
}
