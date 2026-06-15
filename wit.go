package wit

type Wit struct {
	Package    *Package
	Interfaces []Interface
	Worlds     []World
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

// Functions
type Function struct {
	Name    string
	Params  []Param
	Results []Param
}

type Param struct {
	Name string
	Kind Type
}

// Interfaces
type Interface struct {
	Name      string
	TypeDefs  []Type
	Functions []Function
}

// Worlds
type World struct {
	Name    string
	Imports []string
	Exports []string
}

// Packages
//
// package documentation:example;
// package documentation:example@1.0.1;
type Package struct {
	Namespace string
	Package   string
	Interface []string
	Version   string
}
