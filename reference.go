package wit

type Reference string

func (r Reference) witType() {}

func NewReference(ref string) *Reference {
	r := Reference(ref)
	return &r
}

func (e *encoder) encodeReference(ref Reference) {
	e.writeString(string(ref))
}
