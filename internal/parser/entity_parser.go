package parser

import (
	"errors"
	"go/ast"

	"golang.org/x/tools/go/packages"
)

// EntityParser parses codes to retrieve details of an entity.
type EntityParser struct{}

// Parse parses codes to return details of an entity.
func (p EntityParser) Parse(path, entityName string) (*Data, error) {
	pkgs, err := parsePackages(path)
	if err != nil {
		return nil, err
	}

	return p.parseServiceData(pkgs, entityName)
}

func (p EntityParser) parseServiceData(pkgs []*packages.Package, entityName string) (*Data, error) {
	for _, pkg := range pkgs {
		for _, f := range pkg.Syntax {
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range decl.Specs {
						spec, ok := spec.(*ast.TypeSpec)
						if !ok {
							continue
						}

						if spec.Name.Name != entityName {
							continue
						}

						_, ok = spec.Type.(*ast.StructType)
						if !ok {
							continue
						}

						s := &Data{
							Name:        entityName,
							Package:     pkg.PkgPath,
							PackageName: f.Name.Name,
						}

						return s, nil
					}
				}
			}
		}
	}

	return nil, errors.New("serviceParser: no service found")
}
