package config

type TransferProtocol string

const (
	TCP_RPC  TransferProtocol = "tcp-rpc"
	HTTP_RPC TransferProtocol = "http-rpc"
)

type FileSystemConfig struct {
	RootPath string `key:"root-path"`
}

type ProtocolConfig struct {
	TransferProtocol TransferProtocol `key:"transfer-protocol"`
}

type Config struct {
	FileSystemConfig `key:"file-system"`
	ProtocolConfig   `key:"protocol"`
}

func NewDefaultConfig() Config {
	return Config{
		FileSystemConfig: FileSystemConfig{
			RootPath: "/default/path",
		},
		ProtocolConfig: ProtocolConfig{
			TransferProtocol: TCP_RPC,
		},
	}
}
