package main

import (
	"fmt"
	"log"
	"net/http"
)

// FUNCTION OF THE ROUTES

// ======================== formHandler Functions

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succcessful\n")
	// getting the value from the form  and put it in the variable
	name := r.FormValue("name")
	address := r.FormValue("address")
	// printing out the value of the form "%s It is used to format string values"
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// =================== hellohandler Function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// if url is in the wrong path 404 server will be displayed
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// if method is not GET since we dont want the user to put something in the Hello Page 404 server will be displayed
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello Go")
}

func main() {
	// getting the index.html
	fileServer := http.FileServer(http.Dir("./static"))

	//routes for the server to go
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	//checking if the server is starting
	fmt.Printf("Starting  server at port 8080\n ")
	//catching error upon runing the server
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		// to check the log if the server goes down or error.
		log.Fatal(err)
	}

}
