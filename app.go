package main

import (
	"context"
	"encoding/json"
	"fmt"
)

// settings struct
type settings struct {
	location string `json:"stringValue"`
	darkMode bool   `json:"boolValue"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func settingsToJson(set settings) string {
	json, err := json.Marshal(set)
	if err != nil {
		fmt.Println(err)
	}
	return string(json)
}

// Greet returns a greeting for the given name
func (a *App) Greet() string {
	return fmt.Sprintf("hello")
}

func (a *App) GetWeather() string {
	return "Sunny"
}
