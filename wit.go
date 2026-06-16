package wit

type Wit struct {
	Package    *Package
	Interfaces []Interface
	Worlds     []World
}

func (e *Encoder) encodeWit(w Wit) {
	if w.Package != nil {
		e.encodePackage(*w.Package)
		e.writeReturn()
	}
	for _, i := range w.Interfaces {
		e.encodeInterface(i)
	}
	for _, w := range w.Worlds {
		e.encodeWorld(w)
	}
}
