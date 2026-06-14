package wit

import (
	"bufio"
	"fmt"
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
	for _, i := range w.Interfaces {
		b.WriteString(i.EncodeWIT())
	}
	for _, w := range w.Worlds {
		b.WriteString(w.EncodeWIT())
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
	b.WriteString("}\n")
	return b.String()
}

func (i Interface) EncodeWIT() string {
	var b strings.Builder
	b.WriteString("interface ")
	b.WriteString(i.Name)
	b.WriteString(" {\n")
	for _, t := range i.TypeDefs {
		b.WriteString(t.EncodeWIT())
	}
	for _, f := range i.Functions {
		b.WriteString(f.EncodeWIT())
	}
	b.WriteString("}\n")
	return b.String()
}

func (f Function) EncodeWIT() string {
	var b strings.Builder
	b.WriteString("  ")
	b.WriteString(f.Name)
	b.WriteString(": func(")
	for i, p := range f.Params {
		b.WriteString(p.EncodeWIT())
		if i < len(f.Params)-2 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")
	if len(f.Results) < 1 {
		b.WriteString(";\n")
		return b.String()
	}
	b.WriteString(" -> ")
	for i, r := range f.Results {
		b.WriteString(r.EncodeWIT())
		if i < len(f.Results)-2 {
			b.WriteString(", ")
		}
	}
	b.WriteString(";\n")
	return b.String()
}

func (p Param) EncodeWIT() string {
	if p.Name == "" {
		return p.Kind.EncodeWIT()
	}
	return fmt.Sprintf("%s: %s", p.Name, p.Kind.EncodeWIT())
}

func (p PrimitiveType) EncodeWIT() string {
	if p.Kind == "" {
		return "_"
	}
	return string(p.Kind)
}

func (l ListType) EncodeWIT() string {
	return fmt.Sprintf("list<%s>", l.Elem.EncodeWIT())
}

func (o OptionType) EncodeWIT() string {
	return fmt.Sprintf("option<%s>", o.Inner.EncodeWIT())
}

func (r ResultType) EncodeWIT() string {
	return fmt.Sprintf("result<%s, %s>", r.Ok.EncodeWIT(), r.Err.EncodeWIT())
}

func (t TupleType) EncodeWIT() string {
	var b strings.Builder
	b.WriteString("tuple<")
	for i, f := range t.Fields {
		b.WriteString(f.EncodeWIT())
		if i < len(t.Fields)-2 {
			b.WriteString(", ")
		}
	}
	b.WriteRune('>')
	return b.String()
}

func (r Record) EncodeWIT() string {
	var b strings.Builder
	b.WriteString("  record ")
	b.WriteString(r.Name)
	b.WriteString(" {\n")
	for i, f := range r.Fields {
		b.WriteString("    ")
		b.WriteString(f.Name)
		b.WriteString(": ")
		b.WriteString(f.Kind.EncodeWIT())
		if i < len(r.Fields)-2 {
			b.WriteString(",\n")
		} else {
			b.WriteRune('\n')
		}
	}
	b.WriteString("  }\n")
	return b.String()
}
