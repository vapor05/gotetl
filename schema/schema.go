package schema

import (
	"errors"
	"strings"
)

type Column struct {
	Name string
}

type Row struct {
	Columns []Column
}

func ParseRow(line string) (Row, error) {
	names := strings.Split(line, ",")
	if names[0] == "" {
		return Row{}, errors.New("Empty line provided")
	}
	columns := make([]Column, 0, len(names))

	for _, c := range names {
		columns = append(columns, Column{c})
	}

	return Row{columns}, nil
}
