package main

import (
	"fmt"
	"log"
)

var currentId int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write 2nd presentation"})
	RepoCreateTodo(Todo{Name: "Schedule 2nd meeting"})

}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	fmt.Println("Creating Task: ")
	fmt.Println(t.Name)
	t.Name = t.Name //+ stringutil.Reverse("!oG ,olleH\n")
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	log.Println("Could not find Todo with id of %d to delete", id)
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)

}
