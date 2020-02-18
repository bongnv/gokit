package command

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

// Execute ...
func Execute(ctx context.Context) int {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&scaffoldCmd{}, "")
	subcommands.Register(&genCmd{}, "")

	flag.Parse()
	return int(subcommands.Execute(ctx))
}
