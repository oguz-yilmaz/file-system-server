package fsmod

import (
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func ReadFile(params *ReadFileParams) (*File, error) {
	if params.CheckExists {
		if _, err := os.Stat(params.Path); err != nil {
			if os.IsNotExist(err) {
				return nil, errors.New("file does not exist")
			}
			return nil, err
		}
	}

	file, err := os.Open(params.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if params.Offset > 0 {
		_, err = file.Seek(params.Offset, 0)
		if err != nil {
			return nil, err
		}
	}

	readBytes := params.MaxBytes
	if readBytes == 0 {
		readBytes = -1
	}

	reader := io.LimitReader(file, readBytes)
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	if params.Binary {
		encoded := base64.StdEncoding.EncodeToString(content)

		return &File{Content: encoded}, nil
	}

	return &File{Content: string(content), Name: params.Path}, nil
}
