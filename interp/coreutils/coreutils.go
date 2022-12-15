package coreutils

import (
	"context"
	"fmt"

	"github.com/mvdan/u-root-coreutils/pkg/cp"
	"mvdan.cc/sh/v3/interp"
)

// func Handle(ctx context.Context, args []string, stdin io.Reader, stdout, stderr io.Writer) (exit int) {
func Handle(ctx context.Context, args []string) error {
	hc := interp.HandlerCtx(ctx)
	// TODO: hc.Dir, hc.Env
	name, args := args[0], args[1:]
	switch name {
	case "cp":
		exit := cp.RunMain(args, hc.Stdin, hc.Stdout, hc.Stderr)
		if exit != 0 {
			return interp.NewExitStatus(uint8(exit))
		}
		return nil
	default:
		return fmt.Errorf("unsupported builtin: %s", name)
	}
}
