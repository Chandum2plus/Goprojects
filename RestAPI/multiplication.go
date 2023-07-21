package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	ws := &restful.WebService{}
	route := &restful.RouteBuilder{}
	route.Path("/multi")
	route.Doc("Testing the multiplication of the two number ")
	route.Method("POST")
	route.To(multi)
	route.Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(route)
	cs := restful.NewContainer()
	cs.Add(ws)

	log.Println("==== Starting the server on 8081 =====")
	http.ListenAndServe(":8081", cs)
}

// :8081/multi
func multi(req *restful.Request, res *restful.Response) {
	type Number struct {
		Num []int `json:"num"`
	}
	nu := Number{}
	req.ReadEntity(&nu)
	fmt.Println(nu)
	mult := 1

	for _, n := range nu.Num {
		mult = mult * n
	}
	res.WriteAsJson(mult)
}
