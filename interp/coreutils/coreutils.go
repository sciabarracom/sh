package coreutils

import (
	"context"
	"fmt"
	"io"

	"github.com/mvdan/u-root-coreutils/pkg/cp"
)

func Run(ctx context.Context, args []string, stdin io.Reader, stdout, stderr io.Writer) (exit int) {
	name, args := args[0], args[1:]
	switch name {
	case "cp":
		return cp.RunMain(args, stdin, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "interp/coreutils: unsupported builtin: %s", name)
		return 1
	}
}
