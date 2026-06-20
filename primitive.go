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
