package data

import (
	"errors"

	"github.com/vapor05/go-etl/read"
	"github.com/vapor05/go-etl/schema"
)

type CsvFile struct {
	File   string
	Name   string
	Schema schema.Schema
}

func (csv *CsvFile) ReadHeader() error {
	header, err := read.ReadFileLine(csv.File)
	if err != nil {
		return errors.New("Couldn't read file header")
	}

	sch, err := schema.ParseSchema(header)
	if err != nil {
		return errors.New("Couldn't parse schema")
	}

	csv.Schema = sch
	return nil
}
