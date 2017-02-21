package main

import (
	"time"
)

type Todoname struct {
	Name string `json:"name"`
}

type Todosnames []Todoname

type MyMux struct {
}

type WebData struct {
	Title string
}

type WebTasks struct {
	Name string
}
type Todo struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
