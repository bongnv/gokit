package command

import (
	"context"
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/bongnv/gokit/internal/generator"
	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/task"
	"github.com/bongnv/gokit/internal/writer"
	"github.com/google/subcommands"
)

const (
	crudTemplateName        = "crud"
	crudHandlerTemplateName = "crud_handler"
)

var (
	handlersFolder = path.Join(internalFolder, "handlers")
)

type crudCmd struct {
	path          string
	resource      string
	serviceParser parser.Parser
	crudParser    parser.Parser
	writer        writer.Writer
}

func (*crudCmd) Name() string     { return "crud" }
func (*crudCmd) Synopsis() string { return "Generate go-kit codes for a crud service." }
func (*crudCmd) Usage() string {
	return `print [-directory rootDir]:
  Generate go-kit codes for a crud service..
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
	d, err := c.crudParser.Parse(c.path, c.resource)
	if err != nil {
		return err
	}

	crudFilePath := path.Join(c.path, strings.ToLower(c.resource)+"_service.go")
	handlerFilePath := path.Join(c.path, handlersFolder, strings.ToLower(c.resource)+"_handler.go")
	tasks := task.Group{
		&generator.Generator{
			FilePath:     crudFilePath,
			TemplateName: crudTemplateName,
			Data:         d,
			Writer:       c.writer,
		},
		&serviceCmd{
			path:          c.path,
			interfaceName: c.resource + "Service",
			parser:        c.serviceParser,
			writer:        c.writer,
		},
		&generator.Generator{
			FilePath:     handlerFilePath,
			TemplateName: crudHandlerTemplateName,
			Data:         d,
			Writer:       c.writer,
		},
	}

	return tasks.Do()
}
