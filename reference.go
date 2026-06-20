package wit

type Reference string

func (r Reference) witType() {}

func NewReference(ref string) *Reference {
	r := Reference(ref)
	return &r
}

func (e *Encoder) encodeReference(ref Reference) {
	e.writeString(string(ref))
}
