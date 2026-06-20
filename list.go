package wit

func NewList(elem Type) *List {
	return &List{Elem: elem}
}

type List struct {
	Elem Type
}

func (l List) witType() {}

func (e *encoder) encodeList(l List) {
	e.writeString("list<")
	e.encodeType(l.Elem)
	e.writeString(">")
}
