package wit

func NewOption(inner Type) *Option {
	return &Option{Inner: inner}
}

type Option struct {
	Inner Type
}

func (o Option) witType() {}

func (e *encoder) encodeOption(o Option) {
	e.writeString("option<")
	e.encodeType(o.Inner)
	e.writeString(">")
}
