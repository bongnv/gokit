package parser

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
)

func TestEntityParser_parseServiceData(t *testing.T) {
	mockPkgs := []*packages.Package{
		&packages.Package{
			Syntax: []*ast.File{
				{
					Decls: []ast.Decl{
						&ast.GenDecl{
							Specs: []ast.Spec{
								&ast.TypeSpec{
									Type: &ast.StructType{},
									Name: &ast.Ident{
										Name: "Resource",
									},
								},
							},
						},
					},
					Name: &ast.Ident{
						Name: "service",
					},
				},
			},
			PkgPath: "github.com/demo",
		},
	}
	p := EntityParser{}
	d, err := p.parseServiceData(mockPkgs, "Resource")
	require.NoError(t, err)
	require.Equal(t, "Resource", d.Name)
	require.Equal(t, "service", d.PackageName)
	require.Equal(t, "github.com/demo", d.Package)
}
