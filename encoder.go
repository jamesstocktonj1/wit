package wit

import (
	"bufio"
	"io"
	"strings"
)

type encoder struct {
	w      *bufio.Writer
	indent int
	err    error

	indentStr string
	returnStr string
}

func NewEncoder(w io.Writer) *encoder {
	return &encoder{
		w:         bufio.NewWriter(w),
		indentStr: "  ",
		returnStr: "\n",
	}
}

func (e *encoder) Encode(w Wit) error {
	e.encodeWit(w)
	return e.flush()
}

func (e *encoder) flush() error {
	if e.err != nil {
		return e.err
	}
	return e.w.Flush()
}

func (e *encoder) writeString(s string) {
	if e.err != nil {
		return
	}
	_, e.err = e.w.WriteString(s)
}

func (e *encoder) writeIndent() {
	e.writeString(strings.Repeat(e.indentStr, e.indent))
}

func (e *encoder) writeReturn() {
	e.writeString(e.returnStr)
}

func (e *encoder) openBlock() {
	e.indent++
}

func (e *encoder) closeBlock() {
	e.indent--
}
