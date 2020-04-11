package parser

import (
	"errors"

	"golang.org/x/tools/go/packages"
)

// CRUDParser parses codes to generate a service from a resource name.
type CRUDParser struct{}

// Parse parses codes to return a service.
func (p CRUDParser) Parse(path, resourceName string) (*Data, error) {
	pkgs, err := parsePackages(path)
	if err != nil {
		return nil, err
	}

	return p.parseServiceData(pkgs, resourceName)
}

func (p CRUDParser) parseServiceData(pkgs []*packages.Package, resourceName string) (*Data, error) {
	for _, pkg := range pkgs {
		for _, f := range pkg.Syntax {
			s := &Data{
				Name:        resourceName,
				Package:     pkg.PkgPath,
				PackageName: f.Name.Name,
			}
			return s, nil
		}
	}

	return nil, errors.New("crudParser: no service found")
}
