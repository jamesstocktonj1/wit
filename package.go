package wit

import (
	"fmt"
)

func NewPackage(namespace, pkg string) *Package {
	return &Package{Namespace: namespace, Package: pkg}
}

func (p *Package) WithVersion(version string) *Package {
	p.Version = version
	return p
}

func (p *Package) WithInterface(interfaces ...string) *Package {
	p.Interface = append(p.Interface, interfaces...)
	return p
}

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
