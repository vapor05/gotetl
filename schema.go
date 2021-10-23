package goetl

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	intRegexp   = regexp.MustCompile(`\d+`)
	floatRegexp = regexp.MustCompile(`\d+.\d+`)
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

func DetectValueType(value string) (interface{}, error) {
	switch {
	case floatRegexp.MatchString(value):
		fmt.Println("Match float!")
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return UnknownType{}, err
		}
		return FloatType{floatValue}, nil
	case intRegexp.MatchString(value):
		fmt.Println("Match int!")
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return UnknownType{}, err
		}
		return IntegerType{intValue}, nil
	case boolRegexp.MatchString(value):
		fmt.Println("Match bool!")
		boolValue, err := strconv.ParseBool(strings.ToLower(value))
		if err != nil {
			return UnknownType{}, err
		}
		return BoolType{boolValue}, nil
	}

	fmt.Println("Match nothing!")
	return StringType{value}, nil
}
