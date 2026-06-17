package wit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageString(t *testing.T) {
	testMatrix := []struct {
		name string
		pkg  Package
		exp  string
	}{
		{
			name: "simple package",
			pkg: Package{
				Namespace: "foo",
				Package:   "bar",
				Version:   "",
			},
			exp: "foo:bar",
		},
		{
			name: "package - with version",
			pkg: Package{
				Namespace: "bar",
				Package:   "bat",
				Version:   "0.1.2",
			},
			exp: "bar:bat@0.1.2",
		},
		{
			name: "package - with complex version",
			pkg: Package{
				Namespace: "foo",
				Package:   "bat",
				Version:   "3.1.2-rc5",
			},
			exp: "foo:bat@3.1.2-rc5",
		},
	}

	for _, tt := range testMatrix {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.pkg.String()
			assert.Equal(t, tt.exp, res)
		})
	}
}
