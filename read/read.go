package read

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func ReadFileLines(filename string, c chan string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		c <- s.Text()
	}

	close(c)
}

func ReadFileLine(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	hasLine := s.Scan()
	if !hasLine {
		return "", errors.New("File is empty")
	}
	line := s.Text()

	return line, nil
}
