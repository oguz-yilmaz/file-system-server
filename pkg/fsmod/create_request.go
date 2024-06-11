package fsmod

import (
	jsoniter "github.com/json-iterator/go"
	Conf "github.com/oguz-yilmaz/file-system-server/pkg/config"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
	"github.com/oguz-yilmaz/file-system-server/pkg/util"
)

type CreateFileRequest struct {
	protocol.Request

	Params CreateFileParams `json:"params"`
}

type CreateFileParams struct {
	/**
	 * The name of the file to create
	 */
	Name string `json:"name"`
	/**
	 * The content of the file
	 */
	Content string `json:"content"` // TODO use byte array instead?
	/**
	 * The directory where the file should be created
	 */
	Dir string `json:"dir"`
	/**
	 * The type of the file, e.g. text, image, etc.
	 * Not always possible to infer from the extension, it might not be present
	 */
	FileType string `json:"file-type"`
	/**
	 * The permissions of the file, should be sent as decimal number like 438 (octal 0666)
	 */
	Permissions int `json:"permissions"`
	/**
	 * If the file should be overwritten if it already exists, default is false
	 */
	Overwrite bool `json:"overwrite"`
}

func NewCreateFileParams(params map[string]any, conf Conf.Config) *CreateFileParams {
	defaultPermissions := 438 // 0666
	defaultOverwrite := true
	defaultDir := conf.FileSystemConfig.RootPath

	permissions, ok := params["permissions"].(int)
	if !ok {
		permissions = defaultPermissions
	}

	overwrite, ok := params["overwrite"].(bool)
	if !ok {
		overwrite = defaultOverwrite
	}

	dir, ok := params["dir"].(string)
	if !ok {
		dir = defaultDir
	}

	return &CreateFileParams{
		Permissions: permissions,
		Overwrite:   overwrite,
		Dir:         dir,
	}
}

func ValidateCreateRequest(c CreateFileParams, ch protocol.TransportChannel, conf Conf.Config) (bool, CreateFileParams) {
	if c.Name == "" {
		msg := "Error: File name is required"

		errDecoding := ch.Write(protocol.NewError(123, msg))
		ch.Write(errDecoding)

		return false, c
	}

	return true, c
}

func HandleCreateFile(req protocol.Request, channel protocol.TransportChannel, conf Conf.Config) {
	var createFileParams = NewCreateFileParams(map[string]any{}, conf)
	if err := jsoniter.Unmarshal([]byte(req.Params), &createFileParams); err != nil {
		protocol.HandleError(channel, 123, "Error decoding CreateFileRequest")
	}
	util.PrintStruct(createFileParams, "CreateFileParams@")

	validated, params := ValidateCreateRequest(*createFileParams, channel, conf)
	if !validated {
		return
	}

	file, err := CreateFile(&params)
	if err != nil {
		protocol.HandleError(channel, 123, "Error creating file")
	}

	successResponse := NewCreateFileSuccessResponse(req, file)
	channel.Write(successResponse)
}