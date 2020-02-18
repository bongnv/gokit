package main

import (
	"context"
	"os"

	"github.com/bongnv/gokit/internal/command"
)

func main() {
	ctx := context.Background()
	os.Exit(command.Execute(ctx))
}
