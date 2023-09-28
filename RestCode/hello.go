package main

// important libraries
import (
	"fmt"
	"log"
	"net/http"
)

// Home function
func Home() (w http.ResponseWriter, r *http.Request) {
	// This is what the function will print.
	fmt.Fprintf(w, "Welcome to Educative Home!")

	return
}

// function to handle requests
func handleReq() {
	// will call Home function by default.
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8200", nil))
}

func main() {
	// starting the API
	handleReq()
}
