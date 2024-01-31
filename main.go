package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//Proposed features for now
//Home page with weather, time, todos, and calendar as well as a small music widget
//Todo List

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	todo := &Todo{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "oneStop",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			todo,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
