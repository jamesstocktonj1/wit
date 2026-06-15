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
		switch td := t.(type) {
		case *Record:
			e.encodeRecord(*td)
		default:
			e.writeString(td.EncodeWIT())
		}
	}
	e.closeBlock()
	e.writeString("}")
	e.writeReturn()
}
