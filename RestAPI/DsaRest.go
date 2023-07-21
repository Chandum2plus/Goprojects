package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
)

func BinarySearch() *restful.WebService {
	log.Println("")
}
func SetPath() *restful.WebService {
	ws := &restful.WebService{}
	route := restful.RouteBuilder{}
	route.Path("/Binary")
	route.Doc("Binary Search")
	route.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_XML, restful.MIME_JSON)
	route.Method("POST")
	ws.Route(&route)
	cs := restful.NewContainer()
	cs.Add(ws)
	return ws
}
func main() {

}
