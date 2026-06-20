package wit

type Option struct {
	Inner Type
}

func (o Option) witType() {}

func (e *Encoder) encodeOption(o Option) {
	e.writeString("option<")
	e.encodeType(o.Inner)
	e.writeString(">")
}
