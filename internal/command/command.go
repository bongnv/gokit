package command

import (
	"context"
	"flag"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/writer"
	"github.com/google/subcommands"
)

// Execute ...
func Execute(ctx context.Context) int {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&scaffoldCmd{}, "")
	subcommands.Register(&genCmd{
		parser: &parser.DefaultParser{},
		writer: &writer.FileWriter{},
	}, "")
	subcommands.Register(&crudCmd{
		parser: &parser.DefaultParser{},
		writer: &writer.FileWriter{},
	}, "")

	flag.Parse()
	return int(subcommands.Execute(ctx))
}
