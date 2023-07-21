package main

import (
	"fmt"
	"log"
	"net/http"
)

//	func main() {
//		ws := new(restful.WebService)
//		ws.Route(ws.GET("/helio").To(helio))
//		restful.Add(ws)
//		log.Fatal(http.ListenAndServe(":8083", nil))
//
// }
//
//	func helio(req *restful.Request, resp *restful.Response) {
//		resp.Write([]byte("Hello this is chandu kumar and i am in m2-plus\n"))
//		resp.Write([]byte("Hello i am learning Restful api in Go language "))
//	}
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello this is chandu")
	fmt.Println("HomePage is hit :")
}
func handleRequest() {
	http.HandleFunc("/", Homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Fprint()
}
func main() {
	handleRequest()
}
