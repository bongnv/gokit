package command

import (
	"context"
	"flag"

	"github.com/bongnv/gokit/internal/parser"
	"github.com/bongnv/gokit/internal/iohelper"
	"github.com/google/subcommands"
)

// Execute ...
func Execute(ctx context.Context) int {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&scaffoldCmd{
		writer: &iohelper.FileWriter{},
	}, "")

	subcommands.Register(&serviceCmd{
		parser: &parser.DefaultParser{},
		writer: &iohelper.FileWriter{},
	}, "")

	subcommands.Register(&crudCmd{
		serviceParser: &parser.DefaultParser{},
		crudParser:    parser.CRUDParser{},
		writer:        &iohelper.FileWriter{},
		reader:        &iohelper.FileReader{},
		entityParser:  parser.EntityParser{},
	}, "")

	subcommands.Register(&entityCmd{
		writer:       &iohelper.FileWriter{},
		entityParser: parser.EntityParser{},
	}, "")

	flag.Parse()
	return int(subcommands.Execute(ctx))
}
