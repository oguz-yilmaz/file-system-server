package response

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol/resources"
)

type CreateFileResponse struct {
	protocol.Response

	Result resources.TextFile `json:"result"`
}
