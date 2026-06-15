package wit

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEncoder(t *testing.T) {
	enc := NewEncoder(io.Discard)
	assert.NotNil(t, enc)
}
