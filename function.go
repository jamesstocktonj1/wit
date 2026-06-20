package wit

func NewFunction(name string, result *Param, params ...Param) *Function {
	return &Function{Name: name, Params: params, Results: result}
}

func (f *Function) WithDocs(content string) *Function {
	f.Docs = Docs{Content: content}
	return f
}

type Function struct {
	Name    string
	Params  []Param
	Results *Param
	Docs    Docs
}

func NewParam(name string, kind Type) Param {
	return Param{Name: name, Kind: kind}
}

type Param struct {
	Name string
	Kind Type
}

func (e *Encoder) encodeFunction(f Function) {
	e.encodeDocs(f.Docs)
	e.writeIndent()
	e.writeString(f.Name + ": func(")
	for i, p := range f.Params {
		e.encodeParam(p)
		if i < len(f.Params)-1 {
			e.writeString(", ")
		}
	}
	if f.Results == nil {
		e.writeString(");")
	} else {
		e.writeString(") -> ")
		e.encodeParam(*f.Results)
		e.writeString(";")
	}
}

func (e *Encoder) encodeParam(p Param) {
	if p.Name != "" {
		e.writeString(p.Name + ": ")
	}
	e.encodeType(p.Kind)
}
