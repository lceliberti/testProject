package main

import "net/http"

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
