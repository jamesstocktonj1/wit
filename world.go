package wit

type World struct {
	Name    string
	Imports []string
	Exports []string
}

func (e *Encoder) encodeWorld(w World) {
	if len(w.Imports) == 0 && len(w.Exports) == 0 {
		e.writeString("world " + w.Name + " {}")
		return
	}
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
}
