package command

import (
	"context"
	"flag"
	"fmt"
	"path"

	"github.com/google/subcommands"
)

const (
	mainTemplateName     = "main"
	handlersTemplateName = "handlers"
	serviceTemplateName  = "service"
)

var (
	mainFileName     = path.Join("cmd", "server", "main.go")
	handlersFileName = path.Join("internal", "handlers", "service.go")
	serviceFileName  = "service.go"
)

type scaffoldCmd struct {
	dir string
	pkg string
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
	f.StringVar(&c.pkg, "pkg", "github.com/hello", "Package name")
}

func (c *scaffoldCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := c.do(); err != nil {
		fmt.Println("Executed with err:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func (c *scaffoldCmd) do() error {
	writer := &fileWriter{}
	serviceInfo := &Service{
		PackageName: path.Base(c.pkg),
		Package:     c.pkg,
	}

	tasks := taskGroup{
		&fileGenerator{
			filePath:     path.Join(c.dir, serviceFileName),
			templateName: serviceTemplateName,
			service:      serviceInfo,
			writer:       writer,
		},
		&genCmd{
			path: c.dir,
		},
		&fileGenerator{
			filePath:     path.Join(c.dir, handlersFileName),
			templateName: handlersTemplateName,
			service:      serviceInfo,
			writer:       writer,
		},
		&fileGenerator{
			filePath:     path.Join(c.dir, mainFileName),
			templateName: mainTemplateName,
			service:      serviceInfo,
			writer:       writer,
		},
	}

	return tasks.do()
}
