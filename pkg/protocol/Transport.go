package protocol

type TransportChannel interface {
	Read() (*Request, error)
	Write(*Response) error
	Close() error
}
