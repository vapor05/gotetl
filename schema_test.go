package goetl

import (
	"testing"
)

func TestParseSchema(t *testing.T) {
	want := Schema{[]Column{
		{"ID"},
		{"var1"},
		{"date_created"},
		{"column4"},
	}}
	line := "ID,var1,date_created,column4"
	sch, err := ParseSchema(line)

	for i, w := range want.Columns {
		c := sch.Columns[i]
		if c.Name != w.Name || err != nil {
			t.Fatalf(`ParseSchema(line) = %q, want: %v`, sch, want)
		}
	}

	sch, err = ParseSchema("")
	if len(sch.Columns) != 0 || err == nil {
		t.Fatalf(`ParseSchema("") = %q, %v want: Row{[]Column{}}, error`, sch, err)
	}
}

func TestDetectValueType(t *testing.T) {
	wantInt := IntegerType{15}
	actual, err := DetectValueType("15")
	if actual != wantInt || err != nil {
		t.Fatalf(`DetectValueType("15") = %q, want: %v`, actual, wantInt)
	}

	wantF := FloatType{5.5}
	actual, err = DetectValueType("5.5")
	if actual != wantF || err != nil {
		t.Fatalf(`DetectValueType("5.5") = %q, want: %v`, actual, wantF)
	}

	wantB := BoolType{false}
	actual, err = DetectValueType("False")
	if actual != wantB || err != nil {
		t.Fatalf(`DetectValueType("False") = %q, want: %v`, actual, wantB)
	}

	wantS := StringType{"some string value 234"}
	actual, err = DetectValueType("some string value 234")
	if actual != wantS || err != nil {
		t.Fatalf(`DetectValueType("some string value 234") = %q, want: %v, error: %v`, actual, wantS, err)
	}
}
