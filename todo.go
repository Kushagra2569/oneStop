package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type priority int

type Todo struct {
	Id          int      `json:"Id"`
	Description string   `json:"Description"`
	Status      bool     `json:"Status"`
	Priority    priority `json:"priorityValue"`
}

type Todos struct {
	Todos      []Todo `json:"todos"`
	fileLoaded bool
}

const (
	High priority = iota + 1
	Medium
	Low
)

func todosToJson(todos Todos) []byte {
	json, err := json.Marshal(todos)
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func JsonToTodos(todoStr []byte) Todos {
	var todos Todos
	err := json.Unmarshal(todoStr, &todos)
	if err != nil {
		fmt.Println(err)
	}
	return todos
}

func getTodosFromFile() Todos {
	todoStr, err := os.ReadFile(todoFile)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
	}
	todos := JsonToTodos(todoStr)
	return todos
}

func saveTodoToFile(todos Todos) {
	json := todosToJson(todos)
	err := os.WriteFile(todoFile, json, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
