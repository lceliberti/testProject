package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

func main() {

	var vmethod string

	vmethod = "post"
	if vmethod == "get" {
		resp, err := http.Get("http://localhost:8080/todos")

		body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}

		var todos []Todo

		fmt.Println(os.Stdout, string(body)) //<-- here !
		if err = json.Unmarshal([]byte(body), &todos); err != nil {
			log.Println("Error still unmarshelling")
			log.Println(err)
		}

		for _, tasks := range todos {
			fmt.Println("ID: ", tasks.Id, "Name: ", tasks.Name, "Complete: ", tasks.Completed)

		}

	} else {
		url := "http://localhost:8080/todos"
		fmt.Println("URL:>", url)

		var jsonStart = "["
		var jsonBody = `{"name":"my breakfast appt"},{"name":"my lunch appt"}`
		jsonBody = jsonBody + `,{"name":"my dinner appt"}`
		var jsonEnd = "]"
		var jsonRaw = jsonStart + jsonBody + jsonEnd
		//var jsonStr = []byte(`[{"name":"my breakfast appt"},{"name":"my lunch appt"}]`)
		//Convert JSON string into byte slice
		var jsonStr = []byte(jsonRaw)

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
	/*
		fmt.Println("Header")
		fmt.Println(resp.Header)
		fmt.Println("Status")
		fmt.Println(string(resp.StatusCode))
		fmt.Println(string(resp.Status))
		fmt.Println("Trailer")
		fmt.Println(resp.Trailer)
		fmt.Println("Proto")
		fmt.Println(resp.Proto)
	*/
}
