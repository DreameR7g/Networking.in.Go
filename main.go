package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // Accesses the static directory
	http.Handle("/", fileServer)                        // Accepts the path to Static and fileServer variable
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 1338\n")
	if err := http.ListenAndServe(":1338", nil); err != nil {
		log.Fatal(err)
	}
}
