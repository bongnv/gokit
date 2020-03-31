package command

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"path"

	"github.com/bongnv/gokit/internal/generator"
	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/task"
	"github.com/bongnv/gokit/internal/iohelper"
	"github.com/google/subcommands"
)

const (
	mainTemplateName     = "main"
	handlersTemplateName = "handlers"
	serviceTemplateName  = "service"
)

var (
	mainFileName     = path.Join("cmd", "server", "main.go")
	handlersFileName = path.Join("internal", "handlers", "handlers.go")
	serviceFileName  = "service.go"
)

type scaffoldCmd struct {
	dir    string
	pkg    string
	writer iohelper.Writer
}

func (*scaffoldCmd) Name() string     { return "scaffold" }
func (*scaffoldCmd) Synopsis() string { return "Scaffold a go-kit project." }
func (*scaffoldCmd) Usage() string {
	return `scaffold [-directory someDir]:
  Scaffold a go-kit project.
`
}

func (c *scaffoldCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dir, "directory", ".", "Director to scaffold")
	f.StringVar(&c.pkg, "package", "", "project package")
}

func (c *scaffoldCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.Do(); err != nil {
		fmt.Println(err)
		f.Usage()
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *scaffoldCmd) validate() error {
	if c.pkg == "" {
		return errors.New("project package shouldn't be empty")
	}

	return nil
}

func (c *scaffoldCmd) Do() error {
	if err := c.validate(); err != nil {
		return err
	}

	serviceInfo := &parser.Data{
		PackageName: path.Base(c.pkg),
		Package:     c.pkg,
	}

	tasks := task.Group{
		&generator.Generator{
			FilePath:     path.Join(c.dir, serviceFileName),
			TemplateName: serviceTemplateName,
			Data:         serviceInfo,
			Writer:       c.writer,
		},
		&serviceCmd{
			path:          c.dir,
			interfaceName: "Service",
			parser:        &parser.DefaultParser{},
			writer:        c.writer,
		},
		&generator.Generator{
			FilePath:     path.Join(c.dir, handlersFileName),
			TemplateName: handlersTemplateName,
			Data:         serviceInfo,
			Writer:       c.writer,
		},
		&generator.Generator{
			FilePath:     path.Join(c.dir, mainFileName),
			TemplateName: mainTemplateName,
			Data:         serviceInfo,
			Writer:       c.writer,
		},
	}

	return tasks.Do()
}
