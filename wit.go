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

// Types
type Type interface {
	witType()
	EncodeWIT() string
}

// Primitive Types
type PrimitiveType struct {
	Kind PrimKind
}

func (t PrimitiveType) witType() {}

type PrimKind string

const (
	PrimBool       PrimKind = "bool"
	PrimSigned8    PrimKind = "s8"
	PrimSigned16   PrimKind = "s16"
	PrimSigned32   PrimKind = "s32"
	PrimSigned64   PrimKind = "s64"
	PrimUnsigned8  PrimKind = "u8"
	PrimUnsigned16 PrimKind = "u16"
	PrimUnsigned32 PrimKind = "u32"
	PrimUnsigned64 PrimKind = "u64"
	PrimFloat32    PrimKind = "f32"
	PrimFloat64    PrimKind = "f64"
	PrimChar       PrimKind = "char"
	PrimString     PrimKind = "string"
)

// Lists
type ListType struct {
	Elem Type
}

func (t ListType) witType() {}

// Options
type OptionType struct {
	Inner Type
}

func (t OptionType) witType() {}

// Results
type ResultType struct {
	Ok  Type
	Err Type
}

func (t ResultType) witType() {}

// Tuples
type TupleType struct {
	Fields []Type
}

func (t TupleType) witType() {}

// Variants
type Variant struct {
	Cases []Case
}

func (t Variant) witType() {}

type Case struct {
	Name string
	Kind Type
}

// Enums
type Enum struct {
	Cases []string
}

func (t Enum) witType() {}

// Resources
// type Resources struct

// Flags
type Flags struct {
	Labels []string
}

func (t Flags) witType() {}

// Type Aliases
type Alias struct {
	Name string
	Kind Type
}

func (t Alias) witType() {}
