package manipulation

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/oguz-yilmaz/file-system-server/pkg/protocol/methods"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol/resources"
)

func CreateFile(params *methods.CreateFileParams) (*resources.TextFile, error) {
	cwd := params.Dir
	if cwd == nil {
		programCwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		cwd = &programCwd
	}

	ext := filepath.Ext(params.Name)
	mimeType := getMIMEType(ext)

	filePath := filepath.Join(*cwd, params.Name)

	// Create the file
	err := os.WriteFile(filePath, []byte(params.Content), fs.FileMode(params.Permissions))
	if err != nil {
		return nil, err
	}

	file := &resources.File{
		Name:      params.Name,
		Dir:       cwd,
		Version:   1,       // Start with version 1
		Extension: ext[1:], // Remove the leading dot from the extension
		Content:   params.Content,
		MetaData:  params.MetaData,
		MIMEType:  mimeType,
	}

	// Get the file size, TODO: omit this as we have another method to get the file size
	// or just count the bytes of the content to make it more efficient
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	file.Size = info.Size()

	return file, nil
}

// TODO: implement
func getMIMEType(ext string) resources.MIMEType {
	return resources.MIMETextPlain
}
