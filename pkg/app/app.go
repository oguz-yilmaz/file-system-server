package app

import (
	"errors"
	"fmt"
	"net"
	"os"

	jsoniter "github.com/json-iterator/go"
	Conf "github.com/oguz-yilmaz/file-system-server/pkg/config"
	"github.com/oguz-yilmaz/file-system-server/pkg/fsmod"
	"github.com/oguz-yilmaz/file-system-server/pkg/fsmod/request"
	"github.com/oguz-yilmaz/file-system-server/pkg/fsmod/response"
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

func (app *App) StartServer(channel protocol.TransportChannel, conf Conf.Config) {
	var req = protocol.Request{}
	_, err := channel.Read(&req)
	if err != nil {
		panic(err)
	}

	switch req.Method {
	case protocol.METHOD_CREATE_FILE:
		var createFileParams = request.NewCreateFileParams()
		if err := jsoniter.Unmarshal([]byte(req.Params), &createFileParams); err != nil {
			fmt.Println("Error decoding CreateFileRequest:", err)

			return
		}

        PrintStruct(createFileParams, "")

		// -- Validate the file name
		if createFileParams.Name == "" {
			fmt.Println("Error: File name is required")

			return
		}

		// -- Valitade the directory
		if createFileParams.Dir == "" {
			createFileParams.Dir = conf.FileSystemConfig.RootPath
		}

		file, err := fsmod.CreateFile(createFileParams)
		if err != nil {
			fmt.Println("Error creating file:", err)

			return
		}

		fmt.Println("Created file:", file)

		successResponse := response.NewCreateFileSuccessResponse(req, file)
		err = channel.Write(successResponse)

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
}

func Run(config *Conf.Config, startArgs []string) {
	app, err := NewApp(config, startArgs)
	if err == nil {
		conf := Conf.NewDefaultConfig()
		// Not a go channel, it is communication channel, e.g TCP, HTTP, RPC, etc.
		channel, err := makeChannel(&conf)
		if err != nil {
			panic(err)
		}

		for {
			app.StartServer(channel, conf)
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

func PrintStruct(v interface{}, prefix string) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()  // Dereference pointers
	}

	if val.Kind() != reflect.Struct {
		fmt.Printf("Expected a struct, got %s\n", val.Kind())
		return
	}

	// Loop over the fields of the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldName := fieldType.Name

		// Construct a prefix for nested fields
		if prefix != "" {
			fieldName = prefix + "." + fieldName
		}

		// Handle nested structs, except for time.Time or standard library types
		if field.Kind() == reflect.Struct && !field.Type().PkgPath().StartsWith("time") {
			PrintStruct(field.Interface(), fieldName)
		} else {
			fmt.Printf("%s: Type(%T) Value(%v)\n", fieldName, field.Interface(), field.Interface())
		}
	}
}
