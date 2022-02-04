package igloo

import (
	"fmt"
	"io/fs"
)

func ReadFSFile(fsys fs.FS, path string) ([]byte, error) {
	fileBytes, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("reading font file %v: %w", path, err)
	}

	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("font file empty %v", path)
	}

	return fileBytes, nil
}
