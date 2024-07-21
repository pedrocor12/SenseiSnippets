package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SenseiSnippets"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// http.ListenAndServe func starts the server params: Tcp network address
	// to listen on this case port 4000 and the servemux that was just created.
	// Log.Fatal() func logs the error message and exit. Any error returned by
	// ListenAndServer() is always non-nil.
	log.Print("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
