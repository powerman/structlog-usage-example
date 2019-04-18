package pkg

import (
	"io"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
)

// Log is a synonym for convenience.
type Log = *structlog.Logger

var log = structlog.New()

func Something(arg int) error {
	err := io.EOF
	// - Err а не PrintErr.
	return log.Err("failed to do something", "arg", arg, "err", err)
}

func Something2(arg int) error {
	err := io.EOF
	// no way to log arg
	return errors.Wrap(err, "failed to do something")
}
