package command

import (
	"context"
	"flag"
	"fmt"

	"github.com/bongnv/gokit/internal/gokit"
	"github.com/google/subcommands"
)

type genCmd struct {
	dir string
}

func (*genCmd) Name() string     { return "gen" }
func (*genCmd) Synopsis() string { return "Generate go-kit codes." }
func (*genCmd) Usage() string {
	return `print [-directory rootDir]:
  Generate go-kit codes.
`
}

func (c *genCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dir, "directory", ".", "root dir of a go-kit project")
}

func (c *genCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	opts := gokit.Options{Path: c.dir}
	err := gokit.Generate(opts)
	if err != nil {
		fmt.Println("Executed with err:", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
