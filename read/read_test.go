package read

import (
	"testing"
)

func TestReadFileLines(t *testing.T) {
	want := []string{
		"here is line one",
		"some more text",
		"hello from test file!",
		"98 32 123.05",
		"2003",
		"the final line!?",
	}
	c := make(chan string)
	go ReadFileLines("testfile.txt", c)

	i := 0
	for l := range c {
		if want[i] != l {
			t.Fatalf(`Unexpected line from file read, want: %v but got: %v`, want[i], l)
		}
		i += 1
	}
}

func TestReadFileLine(t *testing.T) {
	want := "here is line one"
	line, err := ReadFileLine("testfile.txt")
	if want != line || err != nil {
		t.Fatalf(`Wrong line returned. want: %v but got %v`, want, line)
	}

	line, err = ReadFileLine("empty.txt")
	if line != "" || err == nil {
		t.Fatalf(`Empty line and error expected, want "" but got %v. want error but got %v`, line, err)
	}
}
