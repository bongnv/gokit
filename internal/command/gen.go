package command

import (
	"context"
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/task"
	"github.com/bongnv/gokit/internal/writer"
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

type genCmd struct {
	path          string
	interfaceName string
	parser        parser.Parser
	writer        writer.Writer
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

func (c *genCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.Do(); err != nil {
		fmt.Println("Executed with err:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *genCmd) Do() error {
	if c.parser == nil {
		c.parser = &parser.DefaultParser{
			Path:        c.path,
			ServiceName: c.interfaceName,
		}
	}

	if c.writer == nil {
		c.writer = &writer.FileWriter{}
	}

	s, err := c.parser.Parse()
	if err != nil {
		return err
	}

	tasks := task.Group{
		&fileGenerator{
			filePath:     c.getFilePath("endpoint", endpointsTemplateName),
			templateName: endpointsTemplateName,
			service:      s,
			writer:       c.writer,
		},
		&fileGenerator{
			filePath:     c.getFilePath("server", serverTemplateName),
			templateName: serverTemplateName,
			service:      s,
			writer:       c.writer,
		},
	}

	return tasks.Do()
}

func (c *genCmd) getFilePath(dir, templateName string) string {
	fileName := "z_" + strings.ToLower(c.interfaceName) + "_" + templateName + ".go"
	return path.Join(c.path, internalFolder, dir, fileName)
}
