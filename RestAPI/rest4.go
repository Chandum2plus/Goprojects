package main

import (
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("server is starting ")
}
func handleRequests() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":1000", nil))
}
func main() {
	handleRequests()
}
