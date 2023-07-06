package main

import (
	"fmt"
	"log"
	"net/http"
)

// Here, this home function .................
func home(w http.ResponseWriter, r *http.Request) {
	// this is what the function will print ............
	fmt.Fprintf(w, "hello this chandu here !")
}

// function to handle the request  ....
func handel() {
	// will call home function by default ....
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func main() {
	// Staring the API ..........
	handel()
}
