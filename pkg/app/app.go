package app

import (
	"errors"
	"fmt"
	"net"
	"os"

	Conf "github.com/oguz-yilmaz/file-system-server/pkg/config"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol/channels"
)

type App struct {
}

func (app *App) StartServer() {
	fmt.Println("server started.")
}

func makeChannel(config *Conf.Config) (protocol.TransportChannel, error) {
	switch config.ProtocolConfig.TransferProtocol {
	case Conf.TCP_RPC:
		// Create a TCP connection
		conn, err := net.Dial(config.ProtocolConfig.Network, config.ProtocolConfig.Address)
		if err != nil {
			return nil, err
		}

		return channels.NewTCPChannel(conn), nil
	case Conf.HTTP_RPC:
		// Handle HTTP-RPC if needed
		// This can be implemented similarly to TCPChannel
		return nil, errors.New("HTTP-RPC is not implemented yet")
	default:
		// Default to StdinChannel
		return channels.NewStdinTransport(os.Stdin, os.Stdout), nil
	}
}

func Run(config *Conf.Config, startArgs []string) {
	app, err := NewApp(config, startArgs)
	if err == nil {
		app.StartServer()
	}

	if err != nil {
		panic(err)
	}
}

func NewApp(config *Conf.Config, startArgs []string) (App, error) {
	app := App{}

	return app, nil
}
