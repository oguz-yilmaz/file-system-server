package channels

import (
	"encoding/json"
	"os"

	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

type StdinChannel struct {
	in  *os.File
	out *os.File
}

func NewStdinTransport(in, out *os.File) *StdinChannel {
	return &StdinChannel{in: in, out: out}
}

func (s *StdinChannel) Read() (*protocol.Request, error) {
	var req protocol.Request

	// read from the stdin stream and write it to req struct
	decoder := json.NewDecoder(s.in)
	if err := decoder.Decode(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (s *StdinChannel) Write(res *protocol.Response) error {
	encoder := json.NewEncoder(s.out)

	// Encode and write res struct into the stdout stream
	return encoder.Encode(res)
}

func (s *StdinChannel) Close() error {
	return nil
}
