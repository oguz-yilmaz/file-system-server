package channels

import (
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

type StdinChannel struct {
	in  *os.File
	out *os.File
}

func NewStdinTransport(in, out *os.File) *StdinChannel {
	return &StdinChannel{in: in, out: out}
}

func (s *StdinChannel) Read(req *protocol.Request) (*jsoniter.Decoder, error) {
	// read from the stdin stream and write it to req struct
	decoder := json.NewDecoder(s.in)
	if err := decoder.Decode(req); err != nil {
		return nil, err
	}

	return decoder, nil
}

func (s *StdinChannel) Write(res interface{}) error {
	encoder := json.NewEncoder(s.out)

	// Write res struct into the stdout stream
	return encoder.Encode(res)
}

func (s *StdinChannel) Close() error {
	return nil
}
