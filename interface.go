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
	first := true
	for _, t := range i.TypeDefs {
		if !first {
			e.writeReturn()
		}
		e.encodeType(t)
		e.writeReturn()
		first = false
	}
	for _, f := range i.Functions {
		if !first {
			e.writeReturn()
		}
		e.encodeFunction(f)
		e.writeReturn()
		first = false
	}
	e.closeBlock()
	e.writeString("}")
}
