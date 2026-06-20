package wit

import "strings"

type Docs struct {
	Content string
}

func (e *Encoder) encodeDocs(d Docs) {
	if d.Content == "" {
		return
	}
	content := strings.Split(d.Content, "\n")
	for _, line := range content {
		e.writeIndent()
		e.writeString("/// " + line)
		e.writeReturn()
	}
}
