package schema

import (
    "fmt"
    "errors"

    "goetl/internal/pkg/cmd/options"
)

type CommandOptions struct {
  Command string
  Filename string
  Format string
}

func (co CommandOptions) CommandName() string {
    return co.Command
}

type Command struct {
    Name string
    Opts CommandOptions
}

func (c Command) Action() (int, error) {
    fmt.Println("Run schema command for file, "+c.Opts.Filename+" with format, "+c.Opts.Format)

    return 0, nil
}

func (c Command) GetName() string {
    return c.Name
}

func BuildSchemaOptions(opts options.CommandOptions) (CommandOptions, error) {
    var schemaOpts CommandOptions

    cmd := opts.CommandName()
    filename, ok := opts.Options()["Filename"]
    if !ok {
        return schemaOpts, errors.New("incorrect options.CommandOptions passed to BuildSchemaOptions func, Filename is missing")
    }
    format, ok := opts.Options()["Format"]
    if !ok {
        return schemaOpts, errors.New("incorrect options.CommandOptions passed to BuildSchemaOptions func, Format is missing")
    }
    schemaOpts = CommandOptions{
        Command: cmd,
        Filename: filename,
        Format: format,
    }

    return schemaOpts, nil
}
