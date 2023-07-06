package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/helio").To(helio))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8083", nil))

}
func helio(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte("Hello this is chandu kumar and i am in m2-plus\n"))
	resp.Write([]byte("Hello i am learning Restful api in Go language "))
}
