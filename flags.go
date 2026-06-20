package wit

type Flags struct {
	Name  string
	Cases []Case
	Docs  Docs
}

func (e Flags) witType() {}

func (e *Encoder) encodeFlags(t Flags) {
	e.encodeDocs(t.Docs)
	e.writeIndent()
	e.writeString("flags " + t.Name + " {")
	e.writeReturn()
	e.openBlock()
	for i, c := range t.Cases {
		e.encodeCase(c)
		if i < len(t.Cases)-1 {
			e.writeString(",")
		}
		e.writeReturn()
	}
	e.closeBlock()
	e.writeIndent()
	e.writeString("}")
}
