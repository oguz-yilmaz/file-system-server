package fsmod

import (
	"path/filepath"

	Conf "github.com/oguz-yilmaz/file-system-server/pkg/config"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

type ReadFileRequest struct {
	protocol.Request

	Params ReadFileParams `json:"params"`
}

type ReadFileParams struct {
	/**
	 * The path of the file to read
	 */
	Path string `json:"path"`
	/**
	 * The path of the parent directory of the file
	 */
	Dir string `json:"dir"`
	/**
	 * If set to true, reads the file as binary (base64 string), otherwise as text.
	 * Default is false (read as text)
	 */
	Binary bool `json:"binary,omitempty"`
	/**
	 * Specifies the maximum number of bytes to read from the file.
	 * If not set or zero, it is read until the end of the file.
	 */
	MaxBytes int64 `json:"max-bytes,omitempty"`
	/**
	 * Specifies the offset from where to start reading the file.
	 * If not set or zero, the file is read from the beginning.
	 */
	Offset int64 `json:"offset,omitempty"`
	/**
	 * If set to true, checks if the file exists before reading.
	 * Default is false.
	 */
	CheckExists bool `json:"check-exists,omitempty"`
}

func NewReadFileParams(params map[string]any, conf Conf.Config) *ReadFileParams {
	readFileParams := &ReadFileParams{
		Path:        conf.RootPath,
		Binary:      false,
		MaxBytes:    0,
		Offset:      0,
		CheckExists: false,
	}

	if path, ok := params["path"].(string); ok {
		readFileParams.Path = path
	}
	if dir, ok := params["dir"].(string); ok {
		if !filepath.IsAbs(dir) {
			readFileParams.Dir = filepath.Join(conf.RootPath, dir)
		} else {
			readFileParams.Dir = dir
		}
	}
	if maxBytes, ok := params["max-bytes"].(int64); ok {
		readFileParams.MaxBytes = maxBytes
	}
	if offset, ok := params["offset"].(int64); ok {
		readFileParams.Offset = offset
	}
	if checkExists, ok := params["check-exists"].(bool); ok {
		readFileParams.CheckExists = checkExists
	}
	if binary, ok := params["binary"].(bool); ok {
		readFileParams.Binary = binary
	}

	return readFileParams
}

func ValidateReadFileParams(params ReadFileParams) error {
	if params.Path == "" {
		return NewInvalidParamsError()
	}

	return nil
}

func HandleReadFile(req protocol.Request, channel protocol.TransportChannel, conf Conf.Config) {
	params := NewReadFileParams(map[string]any{}, conf)
	if err := ValidateReadFileParams(*params); err != nil {
		channel.Write(NewReadFileErrorResponse(req, NewGenericError()))

		return
	}

	file, err := ReadFile(params)
	if err != nil {
		channel.Write(NewReadFileErrorResponse(req, NewGenericError()))

		return
	}

	channel.Write(NewReadFileSuccessResponse(req, file))
}
