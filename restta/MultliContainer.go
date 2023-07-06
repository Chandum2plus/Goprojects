package main

import (
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net/http"
)

// this example shows how to have a program with 2 webservices containers
// each having a http server listening on its own port
//
// the first "hello" is added to the restful. DefaultContainer (and uses DefaultServerMux)
// for the second "hello", a new container and serveMux is created
// and requires a new http.Server with hte container being hte handler.
//this first server is spawn in its own go-routine such that the program proceeds to create the second.
//
// GET http://localhost:8080/hello
// GET http://localhost:8081/hello

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	container2 := restful.NewContainer()
	ws2 := new(restful.WebService)
	ws2.Route(ws2.GET("/hello2").To(hello2))
	container2.Add(ws2)
	server := &http.Server{Addr: ":8081", Handler: container2}
	log.Fatal(server.ListenAndServe())
}
func hello(req *restful.Request, res *restful.Response) {
	io.WriteString(res, "default world war")
}
func hello2(req *restful.Request, res *restful.Response) {
	io.WriteString(res, "second world war")
}
