package wit

func NewEnum(name string, cases ...Case) *Enum {
	return &Enum{Name: name, Cases: cases}
}

func (e *Enum) WithDocs(content string) *Enum {
	e.Docs = Docs{Content: content}
	return e
}

type Enum struct {
	Name  string
	Cases []Case
	Docs  Docs
}

func NewCase(name string) Case {
	return Case{Name: name}
}

func (c Case) WithDocs(content string) Case {
	c.Docs = Docs{Content: content}
	return c
}

type Case struct {
	Name string
	Docs Docs
}

func (e Enum) witType() {}

func (e *Encoder) encodeEnum(t Enum) {
	e.encodeDocs(t.Docs)
	e.writeIndent()
	e.writeString("enum " + t.Name + " {")
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

func (e *Encoder) encodeCase(c Case) {
	e.encodeDocs(c.Docs)
	e.writeIndent()
	e.writeString(c.Name)
}
