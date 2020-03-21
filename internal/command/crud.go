package command

import (
	"context"
	"flag"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/bongnv/gokit/internal/generator"
	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/task"
	"github.com/bongnv/gokit/internal/writer"
	"github.com/google/subcommands"
)

const (
	crudTemplateName = "crud"
)

type crudCmd struct {
	path     string
	resource string
	parser   parser.Parser
	writer   writer.Writer
}

func (*crudCmd) Name() string     { return "crud" }
func (*crudCmd) Synopsis() string { return "Generate go-kit codes." }
func (*crudCmd) Usage() string {
	return `print [-directory rootDir]:
  Generate go-kit codes.
`
}

func (c *crudCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "directory", ".", "root path of a go-kit project")
	f.StringVar(&c.resource, "resource", "", "service interface")
}

func (c *crudCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.Do(); err != nil {
		fmt.Println("Executed with err:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *crudCmd) Do() error {
	absPath, err := filepath.Abs(c.path)
	if err != nil {
		return err
	}

	fileName := "z_" + strings.ToLower(c.resource) + ".go"
	tasks := task.Group{
		&generator.Generator{
			FilePath:     path.Join(c.path, fileName),
			TemplateName: crudTemplateName,
			Service: &parser.Service{
				Name:        c.resource,
				PackageName: path.Base(absPath),
			},
			Writer: c.writer,
		},
		&genCmd{
			path:          c.path,
			interfaceName: c.resource + "Service",
			parser:        c.parser,
			writer:        c.writer,
		},
	}

	return tasks.Do()
}
