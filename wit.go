package wit

func NewWit() *Wit {
	return &Wit{}
}

func (w *Wit) WithPackage(p *Package) *Wit {
	w.Package = p
	return w
}

func (w *Wit) WithInterface(interfaces ...Interface) *Wit {
	w.Interfaces = append(w.Interfaces, interfaces...)
	return w
}

func (w *Wit) WithWorld(worlds ...World) *Wit {
	w.Worlds = append(w.Worlds, worlds...)
	return w
}

type Wit struct {
	Docs       Docs
	Package    *Package
	Interfaces []Interface
	Worlds     []World
}

func (e *Encoder) encodeWit(w Wit) {
	e.encodeDocs(w.Docs)
	first := true
	if w.Package != nil {
		e.encodePackage(*w.Package)
		first = false
	}
	for _, i := range w.Interfaces {
		if !first {
			e.writeReturn()
			e.writeReturn()
		}
		e.encodeInterface(i)
		first = false
	}
	for _, world := range w.Worlds {
		if !first {
			e.writeReturn()
			e.writeReturn()
		}
		e.encodeWorld(world)
		first = false
	}
}
