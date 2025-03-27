package service

import (
	"os"
	"strings"
)

type FilePresenter struct {
	path string
}

func NewFilePresenter(path string) *FilePresenter {
	return &FilePresenter{path: path}
}

func (fpres *FilePresenter) Present(lines []string) error {
	content := strings.Join(lines, "\n")

	return os.WriteFile(fpres.path, []byte(content), 0644) //6 - чтение и запись для владельца, 4 и 4 - чтении для группы и остальных
}
