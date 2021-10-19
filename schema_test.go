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
