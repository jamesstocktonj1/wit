package wit

import (
	"bufio"
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
	e.w.WriteRune('\n')
	e.encodeWorlds(w.Worlds)
	return e.flush()
}

func (e *Encoder) flush() error {
	if e.err != nil {
		return e.err
	}
	return e.w.Flush()
}

func (e *Encoder) encodePackage(p Package) {
	e.w.WriteString("package ")
	e.w.WriteString(p.Namespace)
	e.w.WriteRune(':')
	e.w.WriteString(p.Package)
	if p.Version != "" {
		e.w.WriteRune('@')
		e.w.WriteString(p.Version)
	}
	e.w.WriteRune(';')
}

func (e *Encoder) encodeWorlds(w []World) {
	for i, n := range w {
		e.encodeWorld(n)
		if i != len(w)-1 {
			e.w.WriteRune('\n')
		}
	}
}

func (e *Encoder) encodeWorld(w World) {
	e.w.WriteString("world ")
	e.w.WriteString(w.Name)
	e.w.WriteString(" {")
	for _, imp := range w.Imports {
		e.w.WriteString("\n  import ")
		e.w.WriteString(imp)
		e.w.WriteRune(';')
	}
	for _, exp := range w.Exports {
		e.w.WriteString("\n  export ")
		e.w.WriteString(exp)
		e.w.WriteRune(';')
	}
	if len(w.Imports)+len(w.Exports) > 0 {
		e.w.WriteRune('\n')
	}
	e.w.WriteRune('}')
}
