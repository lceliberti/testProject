package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todoId := vars["todoId"]
	inttodoId, err := strconv.Atoi(todoId)
	fmt.Fprintln(w, "Todo show:", todoId)
	if err = json.NewEncoder(w).Encode(RepoFindTodo(inttodoId)); err != nil {
		panic(err)
	}

}

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
	log.Println("Error unmarshelling")
	log.Println(err)

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
	/*call funtion which will loop through slice*/
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
