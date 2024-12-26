package system

import (
	"fmt"
	"os"
	"strings"
)

type FileHandler struct {
	path string
	data []string
}

func NewFileHandler(path string) *FileHandler {
	return &FileHandler{path, nil}
}

func (f *FileHandler) ParseFile() {
	data, err := os.ReadFile(f.path)
	if err != nil {
		fmt.Printf("Error occured: %s", err)
	}

	f.data = strings.Split(string(data), "|")
}
