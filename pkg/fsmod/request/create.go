package request

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
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
	Content string `json:"content"`
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
}

func NewCreateFileParams() *CreateFileParams {
	return &CreateFileParams{
		Permissions: 438, // 0666
	}
}
