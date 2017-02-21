package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type App struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type test_struct struct {
	Test string
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

/*
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todoId := vars["todoId"]
	inttodoId, err := strconv.Atoi(int(todoId))
	fmt.Fprintln(w, "Todo show:", todoId)
	if err = json.NewEncoder(w).Encode(RepoFindTodo(inttodoId)); err != nil {
		panic(err)
	}

}
*/
func TodoDestroy(w http.ResponseWriter, r *http.Request) {
	var todo []Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	err = errors.New("an error")
	/*Unmarshal body parameters into a Todo slice*/
	todo = unmarshallToDo(body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity

	/*Loop through slice to execute delete function*/
	for _, td := range todo {
		err = RepoDestroyTodo(td.Id)
		log.Println("Deleting: ", td.Id)

	}

}

func TodoFind(w http.ResponseWriter, r *http.Request) {
	var todo Todos
	var retTodo Todos
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	err = errors.New("an error")
	/*Unmarshal body parameters into a Todo slice*/
	todo = unmarshallToDo(body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity

	/*Loop through slice to execute delete function*/
	for _, td := range todo {
		retTodo = RepoFindTodo(td.Id)
		log.Println("Found: ", td.Id)
	}

	if err := json.NewEncoder(w).Encode(retTodo); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	err = errors.New("an error")
	var todo []Todo
	/*unmarshal body into todo slice*/
	todo = unmarshallToDo(body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422) // unprocessable entity

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
	/*call funtion which will loop through slice and create a task*/
	t := callRepoFunction(todo)
	log.Println(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}

func unmarshallToDo(ibody []byte) []Todo {
	var todo []Todo
	if err := json.Unmarshal(ibody, &todo); err != nil {
		log.Println("Error unmarshelling the body")
		log.Println(err)

		log.Println(todo)
	}
	return todo
}

func callRepoFunction(todo []Todo) Todo {
	var t Todo
	for _, td := range todo {
		t = RepoCreateTodo(td)
		log.Println(t.Name)
	}
	return t
}
