package config

import "os"

type TransferProtocol string

const (
	TCP_RPC  TransferProtocol = "tcp-rpc"
	HTTP_RPC TransferProtocol = "http-rpc"
	STD_IN   TransferProtocol = "std-in"
)

type GeneralConfig struct {
	WorkerPool int `key:"worker-pool"`
}

type FileSystemConfig struct {
	RootPath string `key:"root-path"`
}

type ProtocolConfig struct {
	TransferProtocol TransferProtocol `key:"transfer-protocol"`
	Address          string           `key:"address"`
	Network          string           `key:"network"`
}

type Config struct {
	GeneralConfig    `key:"general"`
	FileSystemConfig `key:"file-system"`
	ProtocolConfig   `key:"protocol"`
}

func NewDefaultConfig() Config {
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir, err = os.UserHomeDir()

		if err != nil {
			panic(err)
		}
	}

	return Config{
		GeneralConfig: GeneralConfig{
			WorkerPool: 3,
		},
		FileSystemConfig: FileSystemConfig{
			RootPath: currentDir,
		},
		ProtocolConfig: ProtocolConfig{
			TransferProtocol: STD_IN,
		},
	}
}
