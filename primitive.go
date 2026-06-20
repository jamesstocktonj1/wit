package wit

import "fmt"

type Type interface {
	witType()
}

func (e *Encoder) encodeType(t Type) {
	switch tp := t.(type) {
	case *List:
		e.writeString("list<")
		e.encodeType(tp.Elem)
		e.writeString(">")
	case *Option:
		e.writeString("option<")
		e.encodeType(tp.Inner)
		e.writeString(">")
	case *Result:
		e.encodeResult(*tp)
	case *Tuple:
		e.encodeTuple(*tp)
	case *Variant:
		e.encodeVariant(*tp)
	case *Enum:
		e.encodeEnum(*tp)
	case *Flags:
		e.encodeFlags(*tp)
	case *Alias:
		e.encodeAlias(*tp)
	case *Primitive:
		e.writeString(string(*tp))
	case *Reference:
		e.encodeReference(*tp)
	case *Record:
		e.encodeRecord(*tp)
	default:
		e.err = fmt.Errorf("unknown type - %+v", t)
	}
}

type Primitive string

const (
	Bool       Primitive = "bool"
	Signed8    Primitive = "s8"
	Signed16   Primitive = "s16"
	Signed32   Primitive = "s32"
	Signed64   Primitive = "s64"
	Unsigned8  Primitive = "u8"
	Unsigned16 Primitive = "u16"
	Unsigned32 Primitive = "u32"
	Unsigned64 Primitive = "u64"
	Float32    Primitive = "f32"
	Float64    Primitive = "f64"
	Char       Primitive = "char"
	String     Primitive = "string"
)

func (p Primitive) witType() {}

func NewPrimitive(p Primitive) *Primitive {
	return &p
}

type Reference string

func (r Reference) witType() {}

func NewReference(ref string) *Reference {
	r := Reference(ref)
	return &r
}

func (e *Encoder) encodeReference(ref Reference) {
	e.writeString(string(ref))
}

type List struct {
	Elem Type
}

func (l List) witType() {}

type Option struct {
	Inner Type
}

func (o Option) witType() {}

type Result struct {
	Ok  Type
	Err Type
}

func (r Result) witType() {}

// TODO: allow double nil / err nil cases
func (e *Encoder) encodeResult(r Result) {
	e.writeString("result<")
	e.encodeType(r.Ok)
	e.writeString(", ")
	e.encodeType(r.Err)
	e.writeString(">")
}

type Tuple struct {
	Fields []Type
}

func (t Tuple) witType() {}

func (e *Encoder) encodeTuple(t Tuple) {
	e.writeString("tuple<")
	for i, f := range t.Fields {
		e.encodeType(f)
		if i < len(t.Fields)-1 {
			e.writeString(", ")
		}
	}
	e.writeString(">")
}

type Variant struct {
	Name  string
	Cases []Field
	Docs  Docs
}

func (v Variant) witType() {}

func (e *Encoder) encodeVariant(v Variant) {
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

func (e *Encoder) encodeVariantCase(f Field) {
	e.encodeDocs(f.Docs)
	e.writeIndent()
	e.writeString(f.Name)
	if f.Kind != nil {
		e.writeString("(")
		e.encodeType(f.Kind)
		e.writeString(")")
	}
}

type Enum struct {
	Name  string
	Cases []Case
	Docs  Docs
}

type Case struct {
	Name string
	Docs Docs
}

func (e Enum) witType() {}

func (e *Encoder) encodeEnum(t Enum) {
	e.encodeDocs(t.Docs)
	e.writeIndent()
	e.writeString("enum " + t.Name + " {")
	e.writeReturn()
	e.openBlock()
	for i, c := range t.Cases {
		e.encodeCase(c)
		if i < len(t.Cases)-1 {
			e.writeString(",")
		}
		e.writeReturn()
	}
	e.closeBlock()
	e.writeIndent()
	e.writeString("}")
}

func (e *Encoder) encodeCase(c Case) {
	e.encodeDocs(c.Docs)
	e.writeIndent()
	e.writeString(c.Name)
}

type Flags struct {
	Name  string
	Cases []Case
	Docs  Docs
}

func (e Flags) witType() {}

func (e *Encoder) encodeFlags(t Flags) {
	e.encodeDocs(t.Docs)
	e.writeIndent()
	e.writeString("flags " + t.Name + " {")
	e.writeReturn()
	e.openBlock()
	for i, c := range t.Cases {
		e.encodeCase(c)
		if i < len(t.Cases)-1 {
			e.writeString(",")
		}
		e.writeReturn()
	}
	e.closeBlock()
	e.writeIndent()
	e.writeString("}")
}

type Alias struct {
	Name string
	Kind Type
	Docs Docs
}

func (a Alias) witType() {}

func (e *Encoder) encodeAlias(a Alias) {
	e.encodeDocs(a.Docs)
	e.writeIndent()
	e.writeString("type " + a.Name + " = ")
	e.encodeType(a.Kind)
	e.writeString(";")
}
