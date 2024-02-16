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
	IdNum      int    `json:"idNum"`
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
	todoStr, err := LoadFile(todoFile)
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		}
	}
	todos := JsonToTodos(todoStr)
	return todos
}

func saveTodosToFile(todos Todos) {
	json := todosToJson(todos)
	err := WriteFile(todoFile, json)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *Todos) NewTodo(desc string, priorityJS int) string {
	var Pr priority

	if priorityJS == 0 {
		Pr = High
	} else if priorityJS == 1 {
		Pr = Medium
	} else {
		Pr = Low
	}

	todo := Todo{
		Id:          t.IdNum + 1,
		Description: desc,
		Status:      true,
		Priority:    Pr,
	}
	t.Todos = append(t.Todos, todo)
	t.IdNum = t.IdNum + 1
	return string(todosToJson(*t))
}

func (t *Todos) GetTodos() string {
	if !t.fileLoaded {
		*t = getTodosFromFile()
		t.fileLoaded = true
	}
	todoStr := todosToJson(*t) //Kush: fix unnecessary conversion from json to struct and back to json
	return string(todoStr)
}

func (t *Todos) UpdateTodo(TodoId int, status bool, priorityJS int) string {
	var Pr priority

	if priorityJS == 1 {
		Pr = High
	} else if priorityJS == 2 {
		Pr = Medium
	} else {
		Pr = Low
	}

	for i := 0; i < len(t.Todos); i++ {
		if t.Todos[i].Id == TodoId {
			t.Todos[i].Status = status
			t.Todos[i].Priority = Pr
			break
		}
	}
	fmt.Println(string(todosToJson(*t)))
	return string(todosToJson(*t))
}

func (t *Todos) SaveTodos() {
	saveTodosToFile(*t)
}

func (t *Todos) DeleteTodo(TodoId int) string {

	for i := 0; i < len(t.Todos); i++ {
		if t.Todos[i].Id == TodoId {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
			break
		}
	}
	return string(todosToJson(*t))
}
