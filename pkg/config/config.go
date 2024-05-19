package config

type TransferProtocol string

const (
	TCP_RPC  TransferProtocol = "tcp-rpc"
	HTTP_RPC TransferProtocol = "http-rpc"
	STD_IN   TransferProtocol = "std-in"
)

type FileSystemConfig struct {
	RootPath string `key:"root-path"`
}

type ProtocolConfig struct {
	TransferProtocol TransferProtocol `key:"transfer-protocol"`
	Address          string           `key:"address"`
	Network          string           `key:"network"`
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
			TransferProtocol: STD_IN,
		},
	}
}
