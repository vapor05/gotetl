package goetl

import (
	"errors"

	goetl "goetl/read"
)

type CsvFile struct {
	File   string
	Name   string
	Schema Schema
}

func (csv *CsvFile) ReadHeader() error {
	header, err := goetl.ReadFileLine(csv.File)
	if err != nil {
		return errors.New("Couldn't read file header")
	}

	sch, err := ParseSchema(header)
	if err != nil {
		return errors.New("Couldn't parse schema")
	}

	csv.Schema = sch
	return nil
}
