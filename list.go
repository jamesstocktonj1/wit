package wit

type List struct {
	Elem Type
}

func (l List) witType() {}

func (e *Encoder) encodeList(l List) {
	e.writeString("list<")
	e.encodeType(l.Elem)
	e.writeString(">")
}
