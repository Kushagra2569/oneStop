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

// Greet returns a greeting for the given name
func (a *App) Greet() string {
	return fmt.Sprintf("hello")
}

func (a *App) GetWeather() string {
	return "Sunny"
}

func (a *App) NewTodo(desc string, priorityJS int) string {
	var Pr priority

	if priorityJS == 0 {
		Pr = High
	} else if priorityJS == 1 {
		Pr = Medium
	} else {
		Pr = Low
	}

	todo := Todo{
		Id:          1,
		Description: desc,
		Status:      true,
		Priority:    Pr,
	}
	todoStr := todoToJson(todo)
	return string(todoStr)
}

func (t *Todo) GetTodo() string {
	todo := getTodo()
	todoStr := todoToJson(todo)
	return string(todoStr)
}

func (t *Todo) print() {
	fmt.Println(t.Description)
}
