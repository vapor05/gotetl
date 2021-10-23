package schema

import (
    "fmt"

    "goetl/internal/pkg/cmd/options"
)


type Command struct {
    Name string
    Opts options.CommandOptions
}

func (c Command) Action() (int, error) {
    fmt.Println("Run schema command for file, "+c.Opts.Filename+" with format, "+c.Opts.Format)

    return 0, nil
}

func (c Command) GetName() string {
    return c.Name
}
