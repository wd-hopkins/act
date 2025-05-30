package main

import (
	_ "embed"

	"github.com/wd-hopkins/act/cmd"
	"github.com/wd-hopkins/act/pkg/common"
)

//go:embed VERSION
var version string

func main() {
	ctx, cancel := common.CreateGracefulJobCancellationContext()
	defer cancel()

	// run the command
	cmd.Execute(ctx, version)
}
