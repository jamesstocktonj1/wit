package wit

func NewAlias(name string, kind Type) *Alias {
	return &Alias{
		Name: name,
		Kind: kind,
	}
}

func (a *Alias) WithDocs(content string) *Alias {
	a.Docs = Docs{Content: content}
	return a
}

type Alias struct {
	Name string
	Kind Type
	Docs Docs
}

func (a Alias) witType() {}

func (e *Encoder) encodeAlias(a Alias) {
	e.encodeDocs(a.Docs)
	e.writeIndent()
	e.writeString("type " + a.Name + " = ")
	e.encodeType(a.Kind)
	e.writeString(";")
}
