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
		return nil, errors.New("HTTP-RPC is not implemented yet")
	case Conf.STD_IN:
		return channels.NewStdinTransport(os.Stdin, os.Stdout), nil
	default:
		return nil, errors.New("Unknown protocol")
	}
}

func (app *App) StartServer() {
	conf := Conf.NewDefaultConfig()

	channel, err := makeChannel(&conf)
	if err != nil {
		panic(err)
	}

	data, err := channel.Read()
	if err != nil {
		panic(err)
	}

	switch data.Method {
	case protocol.METHOD_CREATE_FILE:
		fmt.Println("creating file")
	case protocol.METHOD_READ_FILE:
		fmt.Println("reading file")
	case protocol.METHOD_DELETE_FILE:
		fmt.Println("deleting file")
	case protocol.METHOD_LIST_FILES:
		fmt.Println("listing files")
	case protocol.METHOD_CREATE_DIR:
		fmt.Println("creating directory")
	case protocol.METHOD_DELETE_DIR:
		fmt.Println("deleting directory")
	case protocol.METHOD_LIST_DIRS:
		fmt.Println("listing directories")
	case protocol.METHOD_MOVE:
		fmt.Println("moving file")
	case protocol.METHOD_COPY:
		fmt.Println("copying file")
	case protocol.METHOD_RENAME:
		fmt.Println("renaming file")
	case protocol.METHOD_SEARCH:
		fmt.Println("searching for file")
	case protocol.METHOD_GET_INFO:
		fmt.Println("getting file info")
	default:
		fmt.Println("unknown method")
	}

	fmt.Println("Received :", data)
}

func Run(config *Conf.Config, startArgs []string) {
	app, err := NewApp(config, startArgs)
	if err == nil {
		for {
			app.StartServer()
		}
	}

	if err != nil {
		panic(err)
	}
}

func NewApp(config *Conf.Config, startArgs []string) (App, error) {
	app := App{}

	return app, nil
}
