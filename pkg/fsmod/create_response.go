package fsmod

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

type CreateFileSuccessResponse struct {
	protocol.Response

	Result File `json:"result"`
}

type CreateFileErrorResponse struct {
	protocol.Response

	Error protocol.Error `json:"error"`
}

func NewCreateFileSuccessResponse(req protocol.Request, file *File) *CreateFileSuccessResponse {
	return &CreateFileSuccessResponse{
		Response: protocol.Response{
			ID: req.ID,
		},
		Result: *file,
	}
}
