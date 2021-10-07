package data

import (
	"testing"

	"github.com/vapor05/go-etl/schema"
)

func TestReadHeader(t *testing.T) {
	want := CsvFile{
		File: "test_simple.csv",
		Name: "test_simple",
		Schema: schema.Schema{
			Columns: []schema.Column{
				{Name: "id"},
				{Name: "col1"},
				{Name: "date"},
				{Name: "number"},
				{Name: "last"},
			}},
	}
	csv := CsvFile{
		File: "test_simple.csv",
		Name: "test_simple",
	}
	err := csv.ReadHeader()

	if err != nil {
		t.Fatalf(`ReadHeader returned an unexpected error, %v`, err)
	}

	for i, w := range want.Schema.Columns {
		c := csv.Schema.Columns[i]
		if c.Name != w.Name {
			t.Fatalf(`ReadHeader parsed incorrect column for csv file schema, want: %v, got %v`, w, c)
		}
	}
}
