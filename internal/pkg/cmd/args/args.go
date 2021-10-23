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
        opts = options.CommandOptions{
            Command: cmd,
            Filename: filename,
            Format: format,
        }
    default:
        return opts, errors.New("unrecognized subcommand: " + cmd)
    }

    return opts, nil
}
