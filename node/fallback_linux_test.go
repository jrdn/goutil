package node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMACFallback(t *testing.T) {
	val, ok := mac()
	assert.Equal(t, true, ok)
	assert.NotEmpty(t, val)
}
