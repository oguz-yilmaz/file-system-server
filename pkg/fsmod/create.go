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

	var err error
	if params.CreateDirs {
		err = createFileWithDirs(filePath, []byte(params.Content), fs.FileMode(params.Permissions), fs.FileMode(params.CreateDirPermissions))
		if err != nil {
			return nil, err
		}
	} else {
		err = os.WriteFile(filePath, []byte(params.Content), fs.FileMode(params.Permissions))
		if err != nil {
			return nil, err
		}
	}

	file := NewFile(params)
	file.Size = len(params.Content)

	return file, nil
}

func createFileWithDirs(filePath string, content []byte, permissions os.FileMode, dirPermissions os.FileMode) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		return err
	}

	err := os.WriteFile(filePath, content, permissions)
	if err != nil {
		return err
	}

	return nil
}
