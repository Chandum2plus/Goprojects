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
	route.Doc("Testing the Year leap or not ")
	route.To(leap)
	route.Method("POST")
	route.Consumes(restful.MIME_JSON, restful.MIME_XML)
	route.Produces(restful.MIME_JSON, restful.MIME_XML)
	route.Path("/leap")
	ws.Route(route)
	cs := restful.NewContainer()
	cs.Add(ws)
	log.Println("Starting the Server on port number 8099")
	http.ListenAndServe(":8089", cs)
}
func leap(req *restful.Request, res *restful.Response) {
	type YEAR struct {
		Year int `json:"year"`
	}
	y := YEAR{}
	fmt.Println(y)
	req.ReadEntity(&y)
	if y.Year%4 == 0 {
		fmt.Fprintln(res, "Leap year ", y.Year)
	} else {
		fmt.Fprintln(res, "Not Leap year ", y.Year)
	}
	res.WriteAsJson(&y)
}
