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
	route.Path("/student")
	route.Doc("Testing the Student details api")
	route.Method("POST")
	route.To(stdnt)
	route.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(route)
	cs := restful.NewContainer()
	cs.Add(ws)

	log.Println("===== Starting the server on 8083 ======")
	http.ListenAndServe(":8083", cs)
}
func stdnt(req *restful.Request, res *restful.Response) {
	type STUDENT struct {
		Name string `json:"name"`
		Add  string `json:"add"`
		Id   int    `json:"id"`
		Roll int    `json:"roll"`
	}
	//student := make(map[string]STUDENT)
	student := STUDENT{}
	req.ReadEntity(&student)
	fmt.Println(student)

	res.WriteAsJson(student)
}
