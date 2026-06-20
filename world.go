package wit

func NewWorld(name string) World {
	return World{Name: name}
}

func (w World) WithImports(imports ...string) World {
	w.Imports = append(w.Imports, imports...)
	return w
}

func (w World) WithExports(exports ...string) World {
	w.Exports = append(w.Exports, exports...)
	return w
}

func (w World) WithDocs(content string) World {
	w.Docs = Docs{Content: content}
	return w
}

type World struct {
	Name    string
	Imports []string
	Exports []string
	Docs    Docs
}

func (e *encoder) encodeWorld(w World) {
	e.encodeDocs(w.Docs)
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
