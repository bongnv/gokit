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
	"github.com/bongnv/gokit/internal/iohelper"
	"github.com/google/subcommands"
)

const (
	endpointsTemplateName = "endpoints"
	serverTemplateName    = "server"
	defaultInterfaceName  = "Service"
)

var (
	internalFolder = "internal"
)

type serviceCmd struct {
	path          string
	interfaceName string
	parser        parser.Parser
	writer        iohelper.Writer
}

func (*serviceCmd) Name() string     { return "service" }
func (*serviceCmd) Synopsis() string { return "Generate go-kit codes for a service." }
func (*serviceCmd) Usage() string {
	return `print [-directory rootDir]:
  Generate go-kit codes for a service.
`
}

func (c *serviceCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "directory", ".", "root path of a go-kit project")
	f.StringVar(&c.interfaceName, "interface", defaultInterfaceName, "service interface")
}

func (c *serviceCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.Do(); err != nil {
		fmt.Println("Executed with err:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *serviceCmd) Do() error {
	s, err := c.parser.Parse(c.path, c.interfaceName)
	if err != nil {
		return err
	}

	tasks := task.Group{
		&generator.Generator{
			FilePath:               c.getFilePath(endpointsTemplateName),
			TemplateName:           endpointsTemplateName,
			Data:                   s,
			Writer:                 c.writer,
			WithAutogeneratedNotes: true,
		},
		&generator.Generator{
			FilePath:               c.getFilePath(serverTemplateName),
			TemplateName:           serverTemplateName,
			Data:                   s,
			Writer:                 c.writer,
			WithAutogeneratedNotes: true,
		},
	}

	return tasks.Do()
}

func (c *serviceCmd) getFilePath(templateName string) string {
	fileName := "z_" + templateName + ".go"
	return path.Join(c.path, internalFolder, strings.ToLower(c.interfaceName), fileName)
}
