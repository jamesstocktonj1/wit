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
	for i, f := range r.Fields {
		e.writeIndent()
		e.encodeField(f)
		if i < len(r.Fields)-1 {
			e.writeString(",")
		}
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

func (e *Encoder) encodeField(f Field) {
	e.writeString(f.Name + ": ")
	e.encodeType(f.Kind)
}
