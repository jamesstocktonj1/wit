package wit

func NewVariant(name string, cases ...Field) *Variant {
	return &Variant{Name: name, Cases: cases}
}

func (v *Variant) WithDocs(content string) *Variant {
	v.Docs = Docs{Content: content}
	return v
}

type Variant struct {
	Name  string
	Cases []Field
	Docs  Docs
}

func (v Variant) witType() {}

func (e *encoder) encodeVariant(v Variant) {
	e.encodeDocs(v.Docs)
	e.writeIndent()
	e.writeString("variant " + v.Name + " {")
	e.writeReturn()
	e.openBlock()
	for i, c := range v.Cases {
		e.encodeVariantCase(c)
		if i < len(v.Cases)-1 {
			e.writeString(",")
		}
		e.writeReturn()
	}
	e.closeBlock()
	e.writeIndent()
	e.writeString("}")
}

func (e *encoder) encodeVariantCase(f Field) {
	e.encodeDocs(f.Docs)
	e.writeIndent()
	e.writeString(f.Name)
	if f.Kind != nil {
		e.writeString("(")
		e.encodeType(f.Kind)
		e.writeString(")")
	}
}
