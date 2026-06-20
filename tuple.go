package wit

func NewTuple(fields ...Type) *Tuple {
	return &Tuple{Fields: fields}
}

type Tuple struct {
	Fields []Type
}

func (t Tuple) witType() {}

func (e *encoder) encodeTuple(t Tuple) {
	e.writeString("tuple<")
	for i, f := range t.Fields {
		e.encodeType(f)
		if i < len(t.Fields)-1 {
			e.writeString(", ")
		}
	}
	e.writeString(">")
}
