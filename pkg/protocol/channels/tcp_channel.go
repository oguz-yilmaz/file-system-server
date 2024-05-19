package channels

import (
	"encoding/json"
	"net"

	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

type TCPChannel struct {
	conn net.Conn
}

func NewTCPChannel(conn net.Conn) *TCPChannel {
	return &TCPChannel{conn: conn}
}

func (t *TCPChannel) Read() (*protocol.Request, error) {
	var req protocol.Request

	decoder := json.NewDecoder(t.conn)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (t *TCPChannel) Write(res *protocol.Response) error {
	encoder := json.NewEncoder(t.conn)

	// Encode writes the JSON encoding of v to the stream(t.conn),
	// followed by a newline character.
	return encoder.Encode(res)
}

func (t *TCPChannel) Close() error {
	return t.conn.Close()
}
