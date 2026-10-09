package main

import (
	"log"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox a lfessak"))
}

func main() {
	// initialize a servemux then register the home function
	// as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// starting a new web server
	// TCP addr is :4000
	// logging any errors
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
