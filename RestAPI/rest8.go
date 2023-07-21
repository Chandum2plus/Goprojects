package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func homes(r *restful.Request, w *restful.Response) {

	w.Write([]byte("hello world\n"))
	w.Write([]byte("Hello india"))
}
func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/helo").To(homes))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
