package command

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"path"
	"strings"

	"github.com/bongnv/gokit/internal/task"
	"github.com/google/subcommands"
	"golang.org/x/tools/go/packages"
)

const (
	endpointsTemplateName = "endpoints"
	serverTemplateName    = "server"
	defaultInterfaceName  = "Service"
)

var (
	internalFolder = "internal"
)

type genCmd struct {
	path          string
	interfaceName string
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
	f.StringVar(&c.interfaceName, "interface", defaultInterfaceName, "service interface")
}

func (c *genCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.Do(); err != nil {
		fmt.Println("Executed with err:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *genCmd) Do() error {
	s, err := c.parseSource()
	if err != nil {
		return err
	}

	writer := &fileWriter{}
	tasks := task.Group{
		&fileGenerator{
			filePath:     c.getFilePath("endpoint", endpointsTemplateName),
			templateName: endpointsTemplateName,
			service:      s,
			writer:       writer,
		},
		&fileGenerator{
			filePath:     c.getFilePath("server", serverTemplateName),
			templateName: serverTemplateName,
			service:      s,
			writer:       writer,
		},
	}

	return tasks.Do()
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

						if spec.Name.Name != c.interfaceName {
							continue
						}

						p := &serviceParser{
							serviceName: c.interfaceName,
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

func (c *genCmd) getFilePath(dir, templateName string) string {
	fileName := "z_" + strings.ToLower(c.interfaceName) + "_" + templateName + ".go"
	return path.Join(c.path, internalFolder, dir, fileName)
}
