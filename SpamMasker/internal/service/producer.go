package service

import (
	"os"
	"strings"
)

type FileProducer struct {
	path string
}

func NewFileProducer(path string) *FileProducer {
	return &FileProducer{path: path}
}

func (fprod *FileProducer) Produce() ([]string, error) {
	content, err := os.ReadFile(fprod.path)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), "\n"), nil
}
