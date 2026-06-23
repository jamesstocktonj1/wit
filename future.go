package wit

func NewFuture(inner Type) *Future {
	return &Future{Inner: inner}
}

type Future struct {
	Inner Type
}

func (f Future) witType() {}

func (e *encoder) encodeFuture(f Future) {
	e.writeString("future<")
	e.encodeType(f.Inner)
	e.writeString(">")
}
