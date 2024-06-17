package fsmod

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

type ReadFileSuccessResponse struct {
	protocol.Response

	Result File `json:"result"`
}

type ReadFileErrorResponse struct {
	protocol.Response

	Error error `json:"error"`
}

func NewReadFileSuccessResponse(req protocol.Request, file *File) *ReadFileSuccessResponse {
	return &ReadFileSuccessResponse{
		Response: protocol.Response{
			ID: req.ID,
		},
		Result: *file,
	}
}

func NewReadFileErrorResponse(req protocol.Request, error error) *ReadFileErrorResponse {
	return &ReadFileErrorResponse{
		Response: protocol.Response{
			ID: req.ID,
		},
		Error: error,
	}
}
