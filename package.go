package wit

import (
	"fmt"
)

func NewPackage(namespace, pkg string) Package {
	return Package{Namespace: namespace, Package: pkg}
}

func NewInterfaceReference(name string) Package {
	return Package{Interface: []string{name}}
}

func (p Package) WithVersion(version string) Package {
	p.Version = version
	return p
}

func (p Package) WithInterface(interfaces ...string) Package {
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

func (p Package) witImportable() {}

func (p Package) String() string {
	hasInterface := len(p.Interface) > 0
	if p.Namespace == "" || p.Package == "" {
		if hasInterface {
			return p.Interface[0]
		}
		return ""
	}

	s := fmt.Sprintf("%s:%s", p.Namespace, p.Package)
	if hasInterface {
		s += "/" + p.Interface[0]
	}
	if p.Version != "" {
		s += "@" + p.Version
	}
	return s
}

func (e *encoder) encodePackage(p Package) {
	e.writeString("package " + p.String() + ";")
}
