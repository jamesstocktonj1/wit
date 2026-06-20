package wit

// Record
type Record struct {
	Name   string
	Fields []Field
	Docs   Docs
}

func (r Record) witType() {}

func (r Record) EncodeWIT() string { return "" }

func (e *Encoder) encodeRecord(r Record) {
	e.encodeDocs(r.Docs)
	e.writeIndent()
	e.writeString("record " + r.Name + " {")
	e.writeReturn()
	e.openBlock()
	for i, f := range r.Fields {
		e.encodeField(f)
		if i < len(r.Fields)-1 {
			e.writeString(",")
		}
		e.writeReturn()
	}
	e.closeBlock()
	e.writeIndent()
	e.writeString("}")
}

// Field
type Field struct {
	Name string
	Kind Type
	Docs Docs
}

func (e *Encoder) encodeField(f Field) {
	e.encodeDocs(f.Docs)
	e.writeIndent()
	e.writeString(f.Name + ": ")
	e.encodeType(f.Kind)
}
