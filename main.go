package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	filename string
	format   string
)

func init() {
	flag.StringVar(&filename, "file", "", "File to parse")
	flag.StringVar(&filename, "f", "", "File to parse")
	flag.StringVar(&format, "format", "", "Format of input file. Formats {csv | json}")
}

func checkFlags() {
	if filename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if format != "csv" && format != "json" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {

	flag.Parse()
	checkFlags()
	fmt.Println(filename)
}
