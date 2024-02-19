package main

import (
	"net/http"
)

func formHandler() {

}

func main() {
	fileServer := http.FileServer(http.Dir("/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

}
