package wit

func NewResource(name string, functions ...Function) *Resource {
	return &Resource{Name: name, Functions: functions}
}

func (r *Resource) WithDocs(content string) *Resource {
	r.Docs = Docs{Content: content}
	return r
}

type Resource struct {
	Name      string
	Functions []Function
	Docs      Docs
}

func (r Resource) witType() {}

func (e *encoder) encodeResource(r Resource) {
	e.encodeDocs(r.Docs)
	e.writeIndent()
	e.writeString("resource " + r.Name + " {")
	e.writeReturn()
	e.openBlock()
	for _, f := range r.Functions {
		e.encodeFunction(f)
		e.writeReturn()
	}
	e.closeBlock()
	e.writeIndent()
	e.writeString("}")
}
