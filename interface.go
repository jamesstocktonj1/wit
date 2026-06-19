package wit

type Interface struct {
	Name      string
	TypeDefs  []Type
	Functions []Function
}

func (e *Encoder) encodeInterface(i Interface) {
	if len(i.TypeDefs) == 0 && len(i.Functions) == 0 {
		e.writeString("interface " + i.Name + " {}")
		return
	}

	e.writeString("interface " + i.Name + " {")
	e.writeReturn()
	e.openBlock()
	for _, t := range i.TypeDefs {
		e.writeIndent()
		e.encodeType(t)
		e.writeReturn()
	}
	for _, f := range i.Functions {
		e.writeIndent()
		e.encodeFunction(f)
		e.writeReturn()
	}
	e.closeBlock()
	e.writeString("}")
}
