package integration

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/jamesstocktonj1/wit"
	"github.com/stretchr/testify/assert"
)

var testMatrix = []struct {
	name      string
	witStruct wit.Wit
	file      string
}{
	{
		name:      "basic",
		witStruct: basic(),
		file:      "basic.wit",
	},
	{
		name:      "empty",
		witStruct: empty(),
		file:      "empty.wit",
	},
}

func TestIntegration(t *testing.T) {
	for _, tt := range testMatrix {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)

			enc := wit.NewEncoder(buf)
			err := enc.Encode(tt.witStruct)
			assert.NoError(t, err)

			exp, err := os.ReadFile(filepath.Join("testdata", tt.file))
			assert.NoError(t, err)

			assert.Equal(t, string(exp), buf.String())
		})
	}
}
