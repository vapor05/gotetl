package args

import (
    "errors"
    "flag"
    "fmt"
    "os"

    "goetl/internal/pkg/cmd/options"
)

var (
    filename string
    format   string
)

func verifyFlags(fs flag.FlagSet) error {
    if filename == "" {
        fs.PrintDefaults()
        return errors.New("file value is required")
    }
    if format != "csv" && format != "json" {
        fs.PrintDefaults()
        return errors.New("format must be 'csv' or 'json'")
    }

    return nil
}

func GetOptions() (options.CommandOptions, error) {
    var opts options.CommandOptions

    schemaCmd := flag.NewFlagSet("schema", flag.ContinueOnError)
    schemaCmd.StringVar(&filename, "file", "", "File to parse")
    schemaCmd.StringVar(&format, "format", "", "Format of input file. Formats {csv | json}")

    if len(os.Args) < 2 {
        return opts, errors.New("subcommand is expected")
    }

    cmd := os.Args[1]

    switch cmd {
    case "schema":
        schemaCmd.Parse(os.Args[2:])
        err := verifyFlags(*schemaCmd)
        if err != nil {
            return opts, fmt.Errorf("Bad arguments passed: %w", err)
        }
        opts = CommandOptions{
            name: cmd,
            commandOptions: map[string]string{
                "Filename": filename,
                "Format": format,
            },
        }
    default:
        return opts, errors.New("unrecognized subcommand: " + cmd)
    }

    return opts, nil
}

type CommandOptions struct {
    name string
    commandOptions map[string]string
}

func (co CommandOptions) CommandName() string {
    return co.name
}

func (co CommandOptions) Options() map[string]string {
    return co.commandOptions
}
