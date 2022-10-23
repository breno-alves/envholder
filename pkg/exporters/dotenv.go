package exporters

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

type DotenvOutput = string

type DotenvExporter[D DotenvOutput] struct {
	Path      string
}

func NewDotenv(path string, _ bool) Exporter[DotenvOutput] {
	return &DotenvExporter[DotenvOutput]{
		Path: path,
	}
}

func (instance *DotenvExporter[D]) Parse(acc []*Variable, line DotenvOutput) ([]*Variable, error) {
	splitedLine := strings.Split(line, "=")

	if len(splitedLine) != 2 {
		return nil, errors.New("invalid line")
	} 

	key := splitedLine[0]
	value := splitedLine[1]
	acc = append(acc, &Variable{Name: key, Value: value})

	return acc, nil
}

func (instance *DotenvExporter[D]) Export() ([]*Variable, error) {
	f, err := os.Open(instance.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	variables := make([]*Variable, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		variables, err = instance.Parse(variables, line)
		if err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return variables, nil
}
