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

	indentStr string
	returnStr string
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:         bufio.NewWriter(w),
		indentStr: "  ",
		returnStr: "\n",
	}
}

func (e *Encoder) Encode(w Wit) error {
	e.encodeWit(w)
	return e.flush()
}

func (e *Encoder) flush() error {
	if e.err != nil {
		return e.err
	}
	return e.w.Flush()
}

func (e *Encoder) writeString(s string) {
	if e.err != nil {
		return
	}
	_, e.err = e.w.WriteString(s)
}

func (e *Encoder) writeIndent() {
	e.writeString(strings.Repeat(e.indentStr, e.indent))
}

func (e *Encoder) writeReturn() {
	e.writeString(e.returnStr)
}

func (e *Encoder) openBlock() {
	e.indent++
}

func (e *Encoder) closeBlock() {
	e.indent--
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
