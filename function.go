package wit

type Function struct {
	Name    string
	Params  []Param
	Results *Param
}

type Param struct {
	Name string
	Kind Type
}

func (e *Encoder) encodeFunction(f Function) {
	e.writeString(f.Name + ": func(")
	for _, p := range f.Params {
		e.encodeParam(p)
	}
	if f.Results == nil {
		e.writeString(");")
	} else {
		e.writeString(") -> ")
		e.writeString(f.Results.EncodeWIT())
		e.writeString(";")
	}
	e.writeReturn()
}

func (e *Encoder) encodeParam(p Param) {
	if p.Name != "" {
		e.writeString(p.Name + ": ")
	}
	e.writeString(p.Kind.EncodeWIT())
}
