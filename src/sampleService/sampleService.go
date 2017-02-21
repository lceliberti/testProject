package main

import (
	"log"
	"net/http"
)

func main() {
	//mux := &MyMux{}
	log.Println("Capturing Request")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":9090", router))
	log.Println("setting URL")
	setURL()

	//http.HandleFunc("/complete/", views.CompleteTaskFunc)
	//http.HandleFunc("/delete/", views.DeleteTaskFunc)
}
