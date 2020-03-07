package command

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Field presents a field in Golang. It includes name & type.
type Field struct {
	Name string
	Type string
}

// Import presents an import in Golang.
type Import struct {
	Path string
	Name string
}

// Endpoint includes details of an endpoint.
type Endpoint struct {
	Name     string
	Method   string
	HTTPPath string
	Params   []Field
	Results  []Field
}

// Service includes details of a service.
type Service struct {
	Endpoints   []Endpoint
	Package     string
	PackageName string
	Imports     []*Import
}

type serviceParser struct {
	pkg         *packages.Package
	f           *ast.File
	serviceType *ast.InterfaceType
	packageName string
}

func (p *serviceParser) parseService() (*Service, error) {
	s := &Service{
		Package:     p.pkg.PkgPath,
		PackageName: p.packageName,
		Endpoints:   p.parseEndpoints(),
		Imports:     extractImports(p.f),
	}

	return s, nil
}

func (p *serviceParser) parseEndpoints() []Endpoint {
	var methods []Endpoint
	for _, method := range p.serviceType.Methods.List {
		fnType, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		params := p.extractFieldsFromAst(fnType.Params.List)
		results := p.extractFieldsFromAst(fnType.Results.List)

		methods = append(methods, Endpoint{
			Name:     method.Names[0].Name,
			HTTPPath: "/" + strings.ToLower(method.Names[0].Name),
			Params:   params,
			Results:  results,
		})
	}

	return methods
}

func (p *serviceParser) getTypeString(expr ast.Expr) string {
	var result string

	switch etype := expr.(type) {
	case *ast.ArrayType:
		result = fmt.Sprintf("[]%s", p.getTypeString(etype.Elt))
	case *ast.MapType:
		result = fmt.Sprintf("map[%s]%s", etype.Key, etype.Value)

	case *ast.SelectorExpr:
		result = fmt.Sprintf("%s.%s", etype.X, etype.Sel)

	case *ast.StarExpr:
		result = fmt.Sprintf("*%s", p.getTypeString(etype.X))

	case *ast.Ident:
		result = fmt.Sprintf("%s.%s", p.packageName, etype.Name)

	default:
		result = fmt.Sprintf("%s", etype)
	}
	return result
}

func (p *serviceParser) extractFieldsFromAst(items []*ast.Field) []Field {
	var output []Field

	for _, item := range items {
		typeStr := p.getTypeString(item.Type)
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
