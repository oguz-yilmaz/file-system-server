package protocol

import (
	jsoniter "github.com/json-iterator/go"
)

// @see https://www.jsonrpc.org/specification
type Request struct {
	RPC    string              `json:"jsonrpc"`
	ID     int                 `json:"id"`
	Method string              `json:"method"`
	Params jsoniter.RawMessage `json:"params"`
}

type Response struct {
	RPC   string `json:"jsonrpc"`
	ID    int    `json:"id,omitempty"`
	Error *Error `json:"error,omitempty"`

	// Each specific Response type will define its own Result field
	Result jsoniter.RawMessage `json:"result,omitempty"`
}

// Notification is a Request without an id
type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}

// Error represents an error in the Response
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	METHOD_CREATE_FILE = "createFile"
	METHOD_READ_FILE   = "readFile"
	METHOD_DELETE_FILE = "deleteFile"
	METHOD_LIST_FILES  = "listFiles"
	METHOD_CREATE_DIR  = "createDir"
	METHOD_DELETE_DIR  = "deleteDir"
	METHOD_LIST_DIRS   = "listDirs"
	METHOD_MOVE        = "move"
	METHOD_COPY        = "copy"
	METHOD_RENAME      = "rename"
	METHOD_SEARCH      = "search"
	METHOD_GET_INFO    = "getInfo"
)

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func HandleError(channel TransportChannel, code int, errMsg string) {
	e := NewError(code, errMsg)

	channel.Write(e)
}

// Handle Batch Requests @see https://www.jsonrpc.org/specification#batch
// TODO: To send several Request objects at the same time, the Client MAY send an Array filled with Request objects.
