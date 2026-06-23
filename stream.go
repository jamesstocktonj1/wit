package wit

func NewStream(elem Type) *Stream {
	return &Stream{Elem: elem}
}

type Stream struct {
	Elem Type
}

func (s Stream) witType() {}

func (e *encoder) encodeStream(s Stream) {
	e.writeString("stream<")
	e.encodeType(s.Elem)
	e.writeString(">")
}
