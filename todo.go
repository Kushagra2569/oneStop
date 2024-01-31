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

const (
	High priority = iota + 1
	Medium
	Low
)

func todoToJson(todo Todo) []byte {
	json, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func JsonToTodo(todoStr []byte) Todo {
	var todo Todo
	err := json.Unmarshal(todoStr, &todo)
	if err != nil {
		fmt.Println(err)
	}
	return todo
}

func getTodo() Todo {
	todoStr, err := os.ReadFile(todoFile)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
	}
	todo := JsonToTodo(todoStr)
	return todo
}

func saveTodoToFile(todo Todo) {
	json := todoToJson(todo)
	err := os.WriteFile(todoFile, json, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
