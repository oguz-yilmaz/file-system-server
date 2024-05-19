package main

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/app"
)

func main() {
	defaultConfig := app.NewDefaultConfig()
	startArgs := []string{}

	app.Run(&defaultConfig, startArgs)
}
