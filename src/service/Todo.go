package main

import "time"

type Todo struct {
	Id        uint64    `json:"id"`
	Name      tname     `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

type tname string
