// Request represents a JSON-RPC request message.
type Request struct {
	// jsonrpc is the JSON-RPC protocol version.
	// This field is required and must be set to "2.0".
	JsonRPC string `json:"jsonrpc"`

	// Method is the name of the method to be invoked.
	// This field is required.
	Method string `json:"method"`

	// Params is an array of parameters to be passed to the method.
	// The number and type of parameters depends on the specific method being called.
	// This field is optional, and may be omitted if the method does not require any parameters.
	Params []any `json:"params,omitempty"`

	// ID is a unique identifier for the request.
	// This field is optional, but recommended for tracking request/response pairs.
	ID *int `json:"id,omitempty"`
}
