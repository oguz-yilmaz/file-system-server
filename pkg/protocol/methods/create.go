package methods

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol/resources"
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
	 * The directory where the file should be created
	 */
	Dir string `json:"dir"`
	/**
	 * The content of the file
	 */
	Text string `json:"text"`
}

type CreateFileResponse struct {
	protocol.Response

	Result resources.TextFile `json:"result"`
}
