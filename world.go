package wit

import "fmt"

func NewWorld(name string) World {
	return World{Name: name}
}

func (w World) WithInclude(includes ...Package) World {
	w.Includes = append(w.Includes, includes...)
	return w
}

func (w World) WithImports(imports ...Importable) World {
	w.Imports = append(w.Imports, imports...)
	return w
}

func (w World) WithExports(exports ...Importable) World {
	w.Exports = append(w.Exports, exports...)
	return w
}

func (w World) WithDocs(content string) World {
	w.Docs = Docs{Content: content}
	return w
}

type World struct {
	Name     string
	Includes []Package
	Imports  []Importable
	Exports  []Importable
	Docs     Docs
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
	for _, in := range w.Includes {
		e.writeIndent()
		e.writeString("include ")
		e.encodePackage(in)
		e.writeString(";")
		e.writeReturn()
	}
	for _, im := range w.Imports {
		e.writeIndent()
		e.writeString("import ")
		e.encodeImport(im)
		e.writeReturn()
	}
	for _, ex := range w.Exports {
		e.writeIndent()
		e.writeString("export ")
		e.encodeImport(ex)
		e.writeReturn()
	}
	e.closeBlock()
	e.writeString("}")
}

type Importable interface {
	witImportable()
}

func (e *encoder) encodeImport(i Importable) {
	switch t := i.(type) {
	case Package:
		e.writeString(t.String())
		e.writeString(";")
	case Interface:
		e.encodeInlineInterface(t)
	default:
		e.err = fmt.Errorf("unknown importable type - %+v", t)
	}
}
