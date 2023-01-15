package coreutils

import (
	"context"
	"fmt"

	"github.com/mvdan/u-root-coreutils/pkg/cp"

	"mvdan.cc/sh/v3/interp"
)

type ErrUnsupported struct {
	Name string
}

func (e *ErrUnsupported) Error() string {
	return fmt.Sprintf("unsupported coreutil: %q", e.Name)
}

func Handle(next interp.ExecHandlerFunc) interp.ExecHandlerFunc {
	return func(ctx context.Context, args []string) error {
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
			// TODO: return ErrUnsupported for the coreutils which we know about but
			// don't yet support
			return next(ctx, args)
		}
	}
}
