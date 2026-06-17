package wit

type Interface struct {
	Name      string
	TypeDefs  []Type
	Functions []Function
}

func (e *Encoder) encodeInterface(i Interface) {
	e.writeString("interface " + i.Name + " {")
	e.writeReturn()
	e.openBlock()
	for _, t := range i.TypeDefs {
		e.writeIndent()
		e.encodeType(t)
	}
	for _, f := range i.Functions {
		e.writeIndent()
		e.encodeFunction(f)
	}
	e.closeBlock()
	e.writeString("}")
	e.writeReturn()
}
