package protocol

// @see https://www.jsonrpc.org/specification
type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// We will just specify the type of the params in all the Request types separately
	// Params
}

type Response struct {
	RPC   string `json:"jsonrpc"`
	ID    *int   `json:"id,omitempty"`
	Error *Error `json:"error,omitempty"`

	// We will just specify the type of the result in all the Response types separately
	// Result
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
	Data    string `json:"data,omitempty"` // Optional: detailed error information etc.
}
