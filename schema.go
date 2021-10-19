package goetl

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	intRegexp   = regexp.MustCompile(`\d+`)
	floatRegexp = regexp.MustCompile(`\d+.\d+?`)
	boolRegexp  = regexp.MustCompile(`(?i)true|(?i)false`)
)

type Column struct {
	Name string
}

type Schema struct {
	Columns []Column
}

func ParseSchema(line string) (Schema, error) {
	names := strings.Split(line, ",")
	if names[0] == "" {
		return Schema{}, errors.New("Empty line provided")
	}
	columns := make([]Column, 0, len(names))

	for _, c := range names {
		columns = append(columns, Column{c})
	}

	return Schema{columns}, nil
}

type StringType struct {
	Value string
}

type IntegerType struct {
	Value int
}

type FloatType struct {
	Value float64
}

type BoolType struct {
	Value bool
}

type DateType struct {
	Value time.Time
}

type UnknownType struct{}

func DetectValueType(value string) interface{} {
	switch {
	case intRegexp.MatchString(value):
		intValue, _ := strconv.Atoi(value)
		return IntegerType{intValue}
	case floatRegexp.MatchString(value):
		floatValue, _ := strconv.ParseFloat(value, 64)
		return FloatType{floatValue}
	case boolRegexp.MatchString(value):
		boolValue, _ := strconv.ParseBool(strings.ToLower(value))
		return BoolType{boolValue}
	}

	return UnknownType{}
}
