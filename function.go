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
