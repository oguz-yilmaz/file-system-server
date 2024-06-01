package channels

import (
	"net"

	jsoniter "github.com/json-iterator/go"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type TCPChannel struct {
	conn net.Conn
}

func NewTCPChannel(conn net.Conn) *TCPChannel {
	return &TCPChannel{conn: conn}
}

func (t *TCPChannel) Read(req *protocol.Request) (*jsoniter.Decoder, error) {
	decoder := json.NewDecoder(t.conn)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return decoder, nil
}

func (t *TCPChannel) Write(res any) error {
	encoder := json.NewEncoder(t.conn)

	// Encode writes the JSON encoding of v to the stream(t.conn),
	// followed by a newline character.
	return encoder.Encode(res)
}

func (t *TCPChannel) Close() error {
	return t.conn.Close()
}
