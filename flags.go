package wit

func NewFlags(name string, cases ...Case) *Flags {
	return &Flags{Name: name, Cases: cases}
}

func (f *Flags) WithDocs(content string) *Flags {
	f.Docs = Docs{Content: content}
	return f
}

type Flags struct {
	Name  string
	Cases []Case
	Docs  Docs
}

func (e Flags) witType() {}

func (e *encoder) encodeFlags(t Flags) {
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
