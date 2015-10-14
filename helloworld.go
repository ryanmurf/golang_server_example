package main

import (
	"net/http"
	"io"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/first/path", helloworld)
	mux.HandleFunc("/second/path", anotherpage)
	mux.Handle("/elmo/", http.StripPrefix("/elmo/", http.FileServer(http.Dir("/go/src/helloworld"))))
	mux.HandleFunc("/", help)
	http.ListenAndServe(":8097", mux)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func anotherpage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You requested page: " + r.URL.String())
}

func help(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the help page. You can call: \n\t/first/path\n\t/second/path\n\t/elmo")
}
