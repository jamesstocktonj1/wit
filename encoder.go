package wit

import (
	"bufio"
	"io"
	"strings"
)

type Encoder struct {
	w      *bufio.Writer
	indent int
	err    error
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: bufio.NewWriter(w)}
}

func (e *Encoder) Encode(w Wit) error {
	e.w.WriteString(w.EncodeWIT())
	return e.flush()
}

func (e *Encoder) flush() error {
	if e.err != nil {
		return e.err
	}
	return e.w.Flush()
}

func (w Wit) EncodeWIT() string {
	var b strings.Builder
	if w.Package != nil {
		b.WriteString("package ")
		b.WriteString(w.Package.EncodeWIT())
		b.WriteString(";\n")
	}
	for _, w := range w.Worlds {
		b.WriteString(w.EncodeWIT())
		b.WriteRune('\n')
	}
	return b.String()
}

func (p Package) EncodeWIT() string {
	var b strings.Builder
	b.WriteString(p.Namespace)
	b.WriteRune(':')
	b.WriteString(p.Package)
	if p.Version != "" {
		b.WriteRune('@')
		b.WriteString(p.Version)
	}
	return b.String()
}

func (w World) EncodeWIT() string {
	var b strings.Builder
	b.WriteString("world ")
	b.WriteString(w.Name)
	b.WriteString(" {")
	for _, imp := range w.Imports {
		b.WriteString("\n  import ")
		b.WriteString(imp)
		b.WriteRune(';')
	}
	for _, exp := range w.Exports {
		b.WriteString("\n  export ")
		b.WriteString(exp)
		b.WriteRune(';')
	}
	if len(w.Imports)+len(w.Exports) > 0 {
		b.WriteRune('\n')
	}
	b.WriteRune('}')
	return b.String()
}
