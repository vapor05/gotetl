package main

import (
    "fmt"
    "log"

    "goetl/internal/pkg/cmd/goetl"
)



func main() {
    cmd, err := goetl.BuildCommand()
    if err != nil {
        log.Fatalf("could not build command: %v", err)
    }
    rc, err := cmd.Action()
    if err != nil {
        log.Fatalf("failed to run command, %v. error: %v", cmd.GetName(), err)
    }

    fmt.Println("Complete with code: ", rc)
}
