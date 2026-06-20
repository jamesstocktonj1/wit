package wit

func NewInterface(name string, typeDefs ...Type) Interface {
	return Interface{Name: name, TypeDefs: typeDefs}
}

func (i Interface) WithFunctions(funcs ...Function) Interface {
	i.Functions = append(i.Functions, funcs...)
	return i
}

func (i Interface) WithDocs(content string) Interface {
	i.Docs = Docs{Content: content}
	return i
}

type Interface struct {
	Name      string
	TypeDefs  []Type
	Functions []Function
	Docs      Docs
}

func (e *encoder) encodeInterface(i Interface) {
	e.encodeDocs(i.Docs)
	if len(i.TypeDefs) == 0 && len(i.Functions) == 0 {
		e.writeString("interface " + i.Name + " {}")
		return
	}

	e.writeString("interface " + i.Name + " {")
	e.writeReturn()
	e.openBlock()
	for _, t := range i.TypeDefs {
		e.encodeType(t)
		e.writeReturn()
	}
	for _, f := range i.Functions {
		e.encodeFunction(f)
		e.writeReturn()
	}
	e.closeBlock()
	e.writeString("}")
}
