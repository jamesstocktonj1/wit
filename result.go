package wit

func NewResult(ok, err Type) *Result {
	return &Result{Ok: ok, Err: err}
}

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
