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
	{
		name:      "list-types",
		witStruct: listTypes(),
		file:      "list-types.wit",
	},
	{
		name:      "option-result",
		witStruct: optionResult(),
		file:      "option-result.wit",
	},
	{
		name:      "enum-flags",
		witStruct: enumFlags(),
		file:      "enum-flags.wit",
	},
	{
		name:      "multi-interface",
		witStruct: multiInterface(),
		file:      "multi-interface.wit",
	},
	{
		name:      "multi-world",
		witStruct: multiWorld(),
		file:      "multi-world.wit",
	},
	{
		name:      "nested-types",
		witStruct: nestedTypes(),
		file:      "nested-types.wit",
	},
	{
		name:      "tuple",
		witStruct: tupleTypes(),
		file:      "tuple.wit",
	},
	{
		name:      "alias",
		witStruct: aliasTypes(),
		file:      "alias.wit",
	},
	{
		name:      "complex-record",
		witStruct: complexRecord(),
		file:      "complex-record.wit",
	},
	{
		name:      "wasi-cli",
		witStruct: wasiCli(),
		file:      "wasi-cli.wit",
	},
	{
		name:      "kv-store",
		witStruct: kvStore(),
		file:      "kv-store.wit",
	},
	{
		name:      "docs",
		witStruct: docs(),
		file:      "docs.wit",
	},
	{
		name:      "variant",
		witStruct: variantTypes(),
		file:      "variant.wit",
	},
	{
		name:      "inline-interface",
		witStruct: inlineInterface(),
		file:      "inline-interface.wit",
	},
	{
		name:      "wasip3-types",
		witStruct: p3Types(),
		file:      "p3-types.wit",
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
