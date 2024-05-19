package app

import (
	"fmt"

	"github.com/oguz-yilmaz/file-system-server/pkg/config"
)

type App struct {
}

func (app *App) StartServer() {
	fmt.Println("server started.")
}

func Run(config *config.Config, startArgs []string) {
	app, err := NewApp(config, startArgs)
	if err == nil {
		app.StartServer()
	}

	if err != nil {
		panic(err)
	}
}

func NewApp(config *config.Config, startArgs []string) (App, error) {
	app := App{}

	return app, nil
}
