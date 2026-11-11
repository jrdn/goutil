package node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFallbackUnavailable(t *testing.T) {
	id, ok := fallbackID(
		func() (string, bool) { return "", false },
		func() (string, bool) { return "", false },
	)
	assert.False(t, ok)
	assert.Empty(t, id)
}
