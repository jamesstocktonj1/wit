package wit

type World struct {
	Name    string
	Imports []string
	Exports []string
}

func (e *Encoder) encodeWorld(w World) {
	e.writeString("world " + w.Name + " {")
	e.writeReturn()
	e.openBlock()
	for _, im := range w.Imports {
		e.writeIndent()
		e.writeString("import " + im + ";")
		e.writeReturn()
	}
	for _, ex := range w.Exports {
		e.writeIndent()
		e.writeString("export " + ex + ";")
		e.writeReturn()
	}
	e.closeBlock()
	e.writeString("}")
	e.writeReturn()
}
