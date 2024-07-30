package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Checks if the current request URL path exactly matches '/'. if it doesnt, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. if we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from SenseiSnippets"))
}

// add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

// Add a snippetCreate hanlder function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check wheter the request is using POST or not.
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Register the two new handler functionsand corresponding URL patterns
	// the servermux, in exactly the same way we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// http.ListenAndServe func starts the server params: Tcp network address
	// to listen on this case port 4000 and the servemux that was just created.
	// Log.Fatal() func logs the error message and exit. Any error returned by
	// ListenAndServer() is always non-nil.
	log.Print("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// need to restructure the project
