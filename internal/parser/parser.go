package parser

import (
	"errors"
	"fmt"
	"go/ast"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	commentWithTagsPrefix = "gokit:"
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
	Tags     map[string]string
}

// Data includes details of a service.
type Data struct {
	Name        string
	Endpoints   []Endpoint
	Package     string
	PackageName string
	Imports     []*Import
}

// Parser is an interface to wrap Parse method.
//go:generate mockery -name=Parser -inpkg -case=underscore
type Parser interface {
	Parse(path, serviceName string) (*Data, error)
}

// DefaultParser is an implementation of parser.
type DefaultParser struct{}

type serviceParser struct {
	serviceName string
	pkg         *packages.Package
	f           *ast.File
	serviceType *ast.InterfaceType
	packageName string
}

// Parse parses codes to return a service.
func (p *DefaultParser) Parse(path, serviceName string) (*Data, error) {
	pkgs, err := parsePackages(path)
	if err != nil {
		return nil, err
	}

	return p.parseServiceData(pkgs, serviceName)
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

func (p *DefaultParser) parseServiceData(pkgs []*packages.Package, serviceName string) (*Data, error) {
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

						if spec.Name.Name != serviceName {
							continue
						}

						p := &serviceParser{
							serviceName: serviceName,
							pkg:         pkg,
							packageName: f.Name.Name,
							serviceType: sType,
							f:           f,
						}

						s := &Data{
							Name:        serviceName,
							Package:     p.pkg.PkgPath,
							PackageName: p.packageName,
							Endpoints:   p.parseEndpoints(),
							Imports:     extractImports(p.f),
						}

						return s, nil
					}
				}
			}
		}
	}

	return nil, errors.New("serviceParser: no service found")
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
		tags := extractTagsFromComment(method.Doc.Text())
		httpPath := "/" + strings.ToLower(method.Names[0].Name)
		if tags["path"] != "" {
			httpPath = tags["path"]
		}

		methods = append(methods, Endpoint{
			Name:     method.Names[0].Name,
			HTTPPath: httpPath,
			Method:   tags["method"],
			Params:   params,
			Results:  results,
			Tags:     tags,
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

func extractTagsFromComment(comment string) map[string]string {
	if !strings.HasPrefix(comment, commentWithTagsPrefix) {
		return nil
	}

	comment = strings.TrimPrefix(comment, commentWithTagsPrefix)
	return getTags(strings.TrimSpace(comment))
}

func getTags(tag string) map[string]string {
	tag = strings.Replace(tag, "`", "", -1)

	out := map[string]string{}

	for tag != "" {
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		value, err := strconv.Unquote(qvalue)
		if err != nil {
			break
		}
		out[name] = value
	}

	return out
}
