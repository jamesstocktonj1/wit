package wit

func NewRecord(name string, fields ...Field) *Record {
	return &Record{Name: name, Fields: fields}
}

func (r *Record) WithDocs(content string) *Record {
	r.Docs = Docs{Content: content}
	return r
}

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
func NewField(name string, kind Type) Field {
	return Field{Name: name, Kind: kind}
}

func (f Field) WithDocs(content string) Field {
	f.Docs = Docs{Content: content}
	return f
}

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
