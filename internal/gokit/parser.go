package parser

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	interfaceName = "Service"
)

type Field struct {
	Name string
	Type string
}

type Import struct {
	Path string
	Name string
}

// Endpoint includes details of an endpoint.
type Endpoint struct {
	Name    string
	Params  []Field
	Results []Field
}

// Service includes details of a service.
type Service struct {
	Endpoints   []Endpoint
	Package     string
	PackageName string
	Imports     []*Import
}

// Parse ...
func Parse(path string) (*Service, error) {
	pkgs, err := parsePackages(path)
	if err != nil {
		return nil, err
	}

	return parseServiceData(pkgs)
}

func parsePackages(path string) ([]*packages.Package, error) {
	parseMode := packages.NeedName |
		packages.NeedFiles |
		packages.NeedImports |
		packages.NeedDeps |
		packages.NeedCompiledGoFiles |
		packages.NeedTypes |
		packages.NeedSyntax |
		packages.NeedTypesInfo

	return packages.Load(
		&packages.Config{
			Mode: parseMode,
		},
		path,
	)
}

func parseServiceData(pkgs []*packages.Package) (*Service, error) {
	for _, pkg := range pkgs {
		for _, f := range pkg.Syntax {
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range decl.Specs {
						spec, ok := spec.(*ast.TypeSpec)
						if !ok {
							continue
						}

						sType, ok := spec.Type.(*ast.InterfaceType)
						if !ok {
							continue
						}

						if spec.Name.Name != interfaceName {
							continue
						}

						d := &Service{
							Package:     pkg.PkgPath,
							PackageName: f.Name.Name,
							Endpoints:   parseEndpointsFrom(sType),
							Imports:     extractImports(f),
						}

						return d, nil
					}
				}
			}
		}
	}

	return nil, errors.New("parser: no service found")
}

func parseEndpointsFrom(i *ast.InterfaceType) []Endpoint {
	var methods []Endpoint
	for _, method := range i.Methods.List {
		fnType, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		params := extractFieldsFromAst(fnType.Params.List)
		results := extractFieldsFromAst(fnType.Results.List)

		methods = append(methods, Endpoint{
			Name:    method.Names[0].Name,
			Params:  params,
			Results: results,
		})
	}

	return methods
}

func getTypeString(expr ast.Expr) string {
	var result string

	switch etype := expr.(type) {
	case *ast.ArrayType:
		result = fmt.Sprintf("[]%s", getTypeString(etype.Elt))
	case *ast.MapType:
		result = fmt.Sprintf("map[%s]%s", etype.Key, etype.Value)

	case *ast.SelectorExpr:
		result = fmt.Sprintf("%s.%s", etype.X, etype.Sel)

	case *ast.StarExpr:
		result = fmt.Sprintf("*%s", getTypeString(etype.X))

	default:
		result = fmt.Sprintf("%s", etype)
	}
	return result
}

func extractFieldsFromAst(items []*ast.Field) []Field {
	var output []Field

	for _, item := range items {
		typeStr := getTypeString(item.Type)
		name := ""

		//  nil if anonymous field
		if len(item.Names) > 0 {
			name = item.Names[0].Name
		}

		output = append(output, Field{
			Type: typeStr,
			Name: name,
		})
	}

	return output
}

func extractImports(f *ast.File) []*Import {
	var res []*Import
	for _, impt := range f.Imports {
		name := ""
		if impt.Name != nil {
			name = impt.Name.Name
		}

		res = append(res, &Import{
			Path: strings.Trim(impt.Path.Value, `"`),
			Name: name,
		})
	}

	return res
}
