package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	//mux := &MyMux{}
	log.Println("Capturing Request")

	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login/", login)
	http.HandleFunc("/delete/", DeleteAllTasksFunc)
	http.HandleFunc("/homeh/", homeHandler)
	http.HandleFunc("/webserviceget/", callWebServiceGet)
	http.HandleFunc("/submitlogin/", submitLogin)
	http.HandleFunc("/submittodos/", submitTodos)
	//http.HandleFunc("/webservicepostparam/", callWebServicePostParam)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	//http.HandleFunc("/complete/", views.CompleteTaskFunc)
	//http.HandleFunc("/delete/", views.DeleteTaskFunc)

}

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Showing My Tasks")
}
func DeleteAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleteing all of my tasks")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	log.Println("method:", r.Method)
	//templateh := "<html> <head><title> Hello World!</title></head><body><form action=\"/login\" method=\"post\">Username:<input type=\"text\" name=\"username\">Password:<input type=\"password\" name=\"password\"><input type=\"submit\" value=\"Login\"></form></body></html>"
	if r.Method == "GET" {
		log.Println("Parsing the login file")
		//var t = template.Must(template.New("name").Parse(templateh))
		var t = template.Must(template.New("login.html").ParseFiles("templates/login.html"))
		//var t = template.Must(template.New("test.html").ParseFiles("templates/test.html"))
		log.Println("Parsed the login file")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func submitLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		log.Println("username:", r.Form["username"])
		log.Println("password:", r.Form["password"])
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("test.html").ParseFiles("templates/test.html"))
	wd := WebData{
		Title: "Home Sweet Home",
	}
	tmpl.Execute(w, &wd)
}

func submitTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	log.Println("method:", r.Method)

	if r.Method == "GET" {
		log.Println("Parsing the submit file")
		var t = template.Must(template.New("submittodos.html").ParseFiles("templates/submittodos.html"))
		log.Println("Parsed the submit file")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("taskname:", r.Form["taskname"])
		{
			url := "http://localhost:8080/todos"
			fmt.Println("URL:>", url)
			var taskname string
			var tdns Todosnames
			for _, tasks := range r.Form["taskname"] {
				taskname = tasks
				fmt.Println("task name: ", taskname)

				//Populate a struct with an appointment
				tdn := Todoname{Name: taskname}
				//Declare a variable of slice type

				//populate slice by appending with struct
				tdns = append(tdns, tdn)

				jsonStr, err := json.Marshal(tdns)
				//Post Request
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				req.Header.Set("X-Custom-Header", "myvalue")
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()

				fmt.Println("response Status:", resp.Status)
				fmt.Println("response Headers:", resp.Header)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("response Body:", string(body))
			}
		}
	}
}
func callWebServiceGet(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/todos")

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}

	/*
		resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

		resp, err := http.PostForm("http://example.com/form",
			url.Values{"key": {"Value"}, "id": {"123"}})
	*/
	var todos []Todo

	fmt.Println(os.Stdout, string(body)) //<-- here !
	if err = json.Unmarshal([]byte(body), &todos); err != nil {
		log.Println("Error still unmarshelling")
		log.Println(err)
	}
	var taskName string
	for _, tasks := range todos {
		taskName = tasks.Name
		fmt.Println("ID: ", tasks.Id, "Name: ", tasks.Name, "Complete: ", tasks.Completed)
		tmpl := template.Must(template.New("webresults.html").ParseFiles("templates/webresults.html"))
		wt := WebTasks{
			Name: taskName,
		}
		tmpl.Execute(w, &wt)

	}

	log.Println("Header")
	log.Println(resp.Header)
	log.Println("Status")
	log.Println(string(resp.StatusCode))
	log.Println(string(resp.Status))
	log.Println("Trailer")
	log.Println(resp.Trailer)
	log.Println("Proto")
	log.Println(resp.Proto)
	fmt.Println("method:", r.Method) //get request method
	log.Println("method:", r.Method)
	//templateh := "<html> <head><title> Hello World!</title></head><body><form action=\"/login\" method=\"post\">Username:<input type=\"text\" name=\"username\">Password:<input type=\"password\" name=\"password\"><input type=\"submit\" value=\"Login\"></form></body></html>"
}
