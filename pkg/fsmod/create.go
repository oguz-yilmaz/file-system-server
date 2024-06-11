package fsmod

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CreateFile(params *CreateFileParams) (*File, error) {
	cwd := params.Dir

	filePath := filepath.Join(cwd, params.Name)

	if !params.Overwrite {
		if _, err := os.Stat(filePath); err == nil {
			return nil, fmt.Errorf("file already exists")
		}
	}

	err := os.WriteFile(filePath, []byte(params.Content), fs.FileMode(params.Permissions))
	if err != nil {
		return nil, err
	}

	file := NewFile(params)

	// TODO: omit this as we have another method to get the file size
	// or just count the bytes of the content to make it more efficient
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	size := int(info.Size())
	file.Size = &size

	return file, nil
}

func WriteFile(file *TextFile) error {
	return nil
}
