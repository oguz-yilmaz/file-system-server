package response

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol/resources"
)

type CreateFileSuccessResponse struct {
	protocol.Response

	Result resources.File `json:"result"`
}

type CreateFileErrorResponse struct {
	protocol.Response

	Error protocol.Error `json:"error"`
}

func NewCreateFileSuccessResponse(req protocol.Request, file *resources.File) *CreateFileSuccessResponse {
	return &CreateFileSuccessResponse{
		Response: protocol.Response{
			ID: req.ID,
		},
		Result: *file,
	}
}
