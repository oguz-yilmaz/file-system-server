package main

import (
	"github.com/oguz-yilmaz/file-system-server/pkg/app"
	"github.com/oguz-yilmaz/file-system-server/pkg/config"
)

func main() {
	defaultConfig := config.NewDefaultConfig()
	startArgs := []string{}

	app.Run(&defaultConfig, startArgs)
}
