package main

import (
	"context"
	"encoding/json"
	"fmt"
)

// settings struct
type settings struct {
	Location string `json:"stringValue"`
	DarkMode bool   `json:"boolValue"`
}

const (
	settingsFile = "settings.json"
	todoFile     = "todo.json"
)

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
