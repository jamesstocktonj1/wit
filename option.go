package wit

type Option struct {
	Inner Type
}

func (o Option) witType() {}
