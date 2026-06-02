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
	return nil
}
