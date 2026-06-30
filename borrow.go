package wit

func NewBorrow(inner Type) *Borrow {
	return &Borrow{Inner: inner}
}

type Borrow struct {
	Inner Type
}

func (b Borrow) witType() {}

func (e *encoder) encodeBorrow(b Borrow) {
	e.writeString("borrow<")
	e.encodeType(b.Inner)
	e.writeString(">")
}
