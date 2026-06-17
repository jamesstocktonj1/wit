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
