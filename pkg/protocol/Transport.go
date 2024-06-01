package protocol

import (
	jsoniter "github.com/json-iterator/go"
)

type TransportChannel interface {
	Read(req *Request) (*jsoniter.Decoder, error)
	Write(res any) error
	Close() error
}
