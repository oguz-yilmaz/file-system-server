package protocol

type Transport interface {
	Read() (*Request, error)
	Write(*Response) error
	Close() error
}
