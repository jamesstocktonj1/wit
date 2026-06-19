package wit

type Wit struct {
	Package    *Package
	Interfaces []Interface
	Worlds     []World
}

func (e *Encoder) encodeWit(w Wit) {
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
