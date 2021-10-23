package goetl

import (
    "fmt"

    "goetl/internal/pkg/cmd/schema"
    "goetl/internal/pkg/cmd/args"
)

type Command interface {
    Action() (int, error)
    GetName() string
}

func BuildCommand() (Command, error) {
    var cmd Command

    opts, err := args.GetOptions()
    if err != nil {
        return cmd, fmt.Errorf("failed to parse command args: %w", err)
    }

    switch opts.Command {
    case "schema":
        cmd = schema.Command{Name: opts.Command, Opts: opts}
    }

    return cmd, nil
}
