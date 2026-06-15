package wit

// Record
type Record struct {
	Name   string
	Fields []Field
}

func (r Record) witType() {}

func (r Record) EncodeWIT() string { return "" }

func (e *Encoder) encodeRecord(r Record) {
	e.writeString("record " + r.Name + " {")
	e.writeReturn()
	e.openBlock()
	for _, f := range r.Fields {
		e.writeIndent()
		e.writeString(f.Name)
		e.writeString(": ")
		e.writeString(f.Kind.EncodeWIT())
		e.writeReturn()
	}
	e.closeBlock()
	e.writeIndent()
	e.writeString("}")
	e.writeReturn()
}

// Field
type Field struct {
	Name string
	Kind Type
}
