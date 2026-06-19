package wit

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEncoder(t *testing.T) {
	enc := NewEncoder(io.Discard)
	assert.NotNil(t, enc)
}

func TestEncoder(t *testing.T) {
	t.Run("basic write", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		enc := NewEncoder(buf)
		enc.writeString("foo")

		err := enc.flush()
		assert.NoError(t, err)

		assert.Equal(t, "foo", buf.String())
	})

	t.Run("write with indent - no block", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		enc := NewEncoder(buf)
		enc.writeIndent()
		enc.writeString("foo")

		err := enc.flush()
		assert.NoError(t, err)

		assert.Equal(t, "foo", buf.String())
	})

	t.Run("write with indent - with block", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		enc := NewEncoder(buf)
		enc.openBlock()
		enc.writeIndent()
		enc.writeString("foo")
		enc.closeBlock()

		err := enc.flush()
		assert.NoError(t, err)

		assert.Equal(t, "  foo", buf.String())
	})

	t.Run("write with return", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		enc := NewEncoder(buf)
		enc.writeString("foo")
		enc.writeReturn()

		err := enc.flush()
		assert.NoError(t, err)

		assert.Equal(t, "foo\n", buf.String())
	})

	t.Run("write with block", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)

		enc := NewEncoder(buf)
		enc.writeString("record foo {")
		enc.writeReturn()

		enc.openBlock()

		enc.writeIndent()
		enc.writeString("bar: str,")
		enc.writeReturn()

		enc.writeIndent()
		enc.writeString("bat: s8")
		enc.writeReturn()

		enc.closeBlock()

		enc.writeString("}")

		err := enc.flush()
		assert.NoError(t, err)

		exp := "record foo {\n"
		exp += "  bar: str,\n"
		exp += "  bat: s8\n"
		exp += "}"
		assert.Equal(t, exp, buf.String())
	})
}
