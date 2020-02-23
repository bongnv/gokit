package command

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"path"

	"github.com/google/subcommands"
	"golang.org/x/tools/go/packages"
)

const (
	endpointsTemplateName = "endpoints"
	serverTemplateName    = "server"
)

var (
	internalFolder    = "internal"
	endpointsFileName = path.Join(internalFolder, "endpoint", "z_endpoints.go")
	serverFileName    = path.Join(internalFolder, "server", "z_server.go")
)

type task interface {
	do() error
}

type genCmd struct {
	path string
}

func (*genCmd) Name() string     { return "gen" }
func (*genCmd) Synopsis() string { return "Generate go-kit codes." }
func (*genCmd) Usage() string {
	return `print [-directory rootDir]:
  Generate go-kit codes.
`
}

func (c *genCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "directory", ".", "root path of a go-kit project")
}

func (c *genCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.do(); err != nil {
		fmt.Println("Executed with err:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *genCmd) do() error {
	s, err := c.parseSource()
	if err != nil {
		return err
	}

	writer := &fileWriter{}
	tasks := []task{
		&fileGenerator{
			filePath:     path.Join(c.path, endpointsFileName),
			templateName: endpointsTemplateName,
			service:      s,
			writer:       writer,
		},
		&fileGenerator{
			filePath:     path.Join(c.path, serverFileName),
			templateName: serverTemplateName,
			service:      s,
			writer:       writer,
		},
	}

	for _, t := range tasks {
		if err := t.do(); err != nil {
			return err
		}
	}

	return nil
}

func (c *genCmd) parseSource() (*Service, error) {
	pkgs, err := parsePackages(c.path)
	if err != nil {
		return nil, err
	}

	return c.parseServiceData(pkgs)
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

func (c *genCmd) parseServiceData(pkgs []*packages.Package) (*Service, error) {
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

						p := &serviceParser{
							f:           f,
							packageName: f.Name.Name,
							serviceType: sType,
							pkg:         pkg,
						}

						return p.parseService()
					}
				}
			}
		}
	}

	return nil, errors.New("serviceParser: no service found")
}
