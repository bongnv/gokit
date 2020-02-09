package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bongnv/gokit/internal/gokit"
)

func main() {
	opts := parseOptionsFromCmd()
	if err := gokit.Generate(opts); err != nil {
		exitWithErr(err)
	}
}

func parseOptionsFromCmd() gokit.Options {
	opts := gokit.Options{}
	flag.StringVar(&opts.Path, "d", ".", "Path to source directory")

	flag.Parse()

	return opts
}

func exitWithErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
