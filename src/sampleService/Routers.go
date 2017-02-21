package main

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		fmt.Println("getting headers")

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

/*
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login/", login)
	http.HandleFunc("/delete/", DeleteAllTasksFunc)
	http.HandleFunc("/homeh/", homeHandler)
	http.HandleFunc("/webserviceget/", callWebServiceGet)
	http.HandleFunc("/submitlogin/", submitLogin)
	http.HandleFunc("/submittodos/", submitTodos)
	//http.HandleFunc("/webservicepostparam/", callWebServicePostParam)
*/

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		sayhelloName,
	},
	Route{
		"login",
		"GET",
		"/login",
		login,
	},
	Route{
		"submittodos",
		"POST",
		"/submittodos",
		submitTodos,
	},
	Route{
		"submittodos",
		"GET",
		"/submittodos",
		submitTodos,
	},
	Route{
		"webserviceget",
		"GET",
		"/webserviceget",
		callWebServiceGet,
	},
}
