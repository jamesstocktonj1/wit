package wit

import (
	"bufio"
	"fmt"
	"io"
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
	e.encodePackage(w.Package)
	return e.flush()
}

func (e *Encoder) flush() error {
	if e.err != nil {
		return e.err
	}
	return e.w.Flush()
}

func (e *Encoder) encodePackage(p Package) {
	e.w.WriteString(fmt.Sprintf("package %s:%s", p.Namespace, p.Package))
	if p.Version != "" {
		e.w.WriteRune('@')
		e.w.WriteString(p.Version)
	}
	e.w.WriteRune(';')
}
