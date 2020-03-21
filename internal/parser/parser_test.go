package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func getASTFromSrc(src string) *ast.File {
	fs := token.NewFileSet()
	srcAST, _ := parser.ParseFile(fs, "", src, parser.ParseComments)
	return srcAST
}

func Test_parseEndpointsFrom(t *testing.T) {
	src := `
package hello

import "context"

// Service is a simple interface for a service.
type Service interface {
	//gokit: path:"/say-hello"
	Hello(ctx context.Context, p Person) error
}

// Person presents a single person.
type Person struct {
	Name string ` + "`" + `json:"name"` + "`" + `
}
`

	srcAST := getASTFromSrc(src)
	decl := srcAST.Decls[1].(*ast.GenDecl)
	spec := decl.Specs[0].(*ast.TypeSpec)
	p := &serviceParser{
		serviceType: spec.Type.(*ast.InterfaceType),
		f:           srcAST,
		packageName: srcAST.Name.Name,
	}
	endpoints := p.parseEndpoints()
	require.Len(t, endpoints, 1)
	require.Equal(t, "/say-hello", endpoints[0].HTTPPath)
}
