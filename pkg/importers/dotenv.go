package importers

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type DotenvImporter struct {
	Path string
}

func NewDotenvImporter(path string) *DotenvImporter {
	return &DotenvImporter{
		Path: path,
	}
}

func (instance *DotenvImporter) Import() ([]*Variable, error) {
	return nil, nil
}

func (instance *DotenvImporter) Parse(acc []*Variable, data string) ([]*Variable, error) {
	return nil, nil
}

func (instance *DotenvImporter) Read() ([]*Variable, error) {
	f, err := os.Open(instance.Path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// variables := make([]*Variable, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil, nil
}
