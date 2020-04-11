package parser

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
)

func TestCRUDParser_parseServiceData(t *testing.T) {
	t.Run("happy-path", func(t *testing.T) {
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
		p := CRUDParser{}
		d, err := p.parseServiceData(mockPkgs, "Resource")
		require.NoError(t, err)
		require.Equal(t, "Resource", d.Name)
		require.Equal(t, "service", d.PackageName)
		require.Equal(t, "github.com/demo", d.Package)
	})

	t.Run("not-found", func(t *testing.T) {
		p := CRUDParser{}
		_, err := p.parseServiceData(nil, "Resource")
		require.Error(t, err)
	})
}
