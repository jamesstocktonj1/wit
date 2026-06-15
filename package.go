package wit

import (
	"fmt"
)

// Package
//
// package documentation:example;
// package documentation:example@1.0.1;
type Package struct {
	Namespace string
	Package   string
	Interface []string
	Version   string
}

func (p Package) String() string {
	pkg := fmt.Sprintf("%s:%s", p.Namespace, p.Package)
	if p.Version != "" {
		pkg += fmt.Sprintf("@%s", p.Version)
	}
	return pkg
}

func (e *Encoder) encodePackage(p Package) {
	e.writeString("package " + p.String() + ";")
}
