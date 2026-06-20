package wit

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
