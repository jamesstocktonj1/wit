package wit

func NewFunction(name string, params ...Param) Function {
	return Function{Name: name, Params: params}
}

func (f Function) WithResult(result Type) Function {
	f.Result = result
	return f
}

func (f Function) WithDocs(content string) Function {
	f.Docs = Docs{Content: content}
	return f
}

func (f Function) SetAsync() Function {
	f.Async = true
	return f
}

type Function struct {
	Name   string
	Params []Param
	Result Type
	Async  bool
	Docs   Docs
}

func NewParam(name string, kind Type) Param {
	return Param{Name: name, Kind: kind}
}

type Param struct {
	Name string
	Kind Type
}

func (e *encoder) encodeFunction(f Function) {
	e.encodeDocs(f.Docs)
	e.writeIndent()
	e.writeString(f.Name + ": ")
	if f.Async {
		e.writeString("async ")
	}
	e.writeString("func(")
	for i, p := range f.Params {
		e.encodeParam(p)
		if i < len(f.Params)-1 {
			e.writeString(", ")
		}
	}
	if f.Result == nil {
		e.writeString(");")
	} else {
		e.writeString(") -> ")
		e.encodeType(f.Result)
		e.writeString(";")
	}
}

func (e *encoder) encodeParam(p Param) {
	if p.Name != "" {
		e.writeString(p.Name + ": ")
	}
	e.encodeType(p.Kind)
}
