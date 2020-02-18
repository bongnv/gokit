package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type scaffoldCmd struct {
	dir string
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
}

func (c *scaffoldCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("Running scaffold cmd in ", c.dir)
	return subcommands.ExitSuccess
}
