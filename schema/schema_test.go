package schema

import (
	"testing"
)

func TestParseRow(t *testing.T) {
	want := Row{[]Column{
		{"ID"},
		{"var1"},
		{"date_created"},
		{"column4"},
	}}
	line := "ID,var1,date_created,column4"
	row, err := ParseRow(line)

	for i, w := range want.Columns {
		c := row.Columns[i]
		if c.Name != w.Name || err != nil {
			t.Fatalf(`ParseRow(line) = %q, want: %v`, row, want)
		}
	}

	row, err = ParseRow("")
	if len(row.Columns) != 0 || err == nil {
		t.Fatalf(`ParseRow("") = %q, %v want: Row{[]Column{}}, error`, row, err)
	}
}
