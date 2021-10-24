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
        return nil, fmt.Errorf("failed to parse command args: %w", err)
    }

    switch opts.CommandName() {
    case "schema":
        schemaOpts, err := schema.BuildSchemaOptions(opts)
        if err != nil {
            return nil, fmt.Errorf("could not build schema command options: %w", err)
        }
        cmd = schema.Command{Name: opts.CommandName(), Opts: schemaOpts}
    default:
        return nil, fmt.Errorf("unknown scommand")
    }

    return cmd, nil
}
