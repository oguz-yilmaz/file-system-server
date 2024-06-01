package fsmod

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/oguz-yilmaz/file-system-server/pkg/fsmod/request"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol/resources"
)

func CreateFile(params *request.CreateFileParams) (*resources.File, error) {
	cwd := params.Dir

	filePath := filepath.Join(cwd, params.Name)

	// Create the file, use the below function and pass the below
	// file struct to it to create the file
	fmt.Println("--- fsmod@CreateFile ---")
	fmt.Println("filePath: ", filePath)
	fmt.Println("Content: ", params.Content)
	fmt.Println("Permissions: ", fs.FileMode(params.Permissions))
	fmt.Println("Overwrite: ", params.Overwrite)
	fmt.Println("--- fsmod@CreateFile ---")

	if !params.Overwrite {
		if _, err := os.Stat(filePath); err == nil {
			return nil, fmt.Errorf("file already exists")
		}
	}

	err := os.WriteFile(filePath, []byte(params.Content), fs.FileMode(params.Permissions))
	if err != nil {
		return nil, err
	}

	file := resources.NewFile(params)

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

func WriteFile(file *resources.TextFile) error {
	return nil
}
