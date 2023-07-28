package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

/*
print your name, date of birth, and mobile number.
Expected Output:

Name   : Alexandra Abramov
DOB    : July 14, 1975
Mobile : 99-9999999999

*/

func main() {
	ws := &restful.WebService{}
	route := &restful.RouteBuilder{}
	route.Path("/details")
	route.Doc("Testing the Details printing using the Rest api")
	route.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_XML, restful.MIME_JSON)
	route.Method("POST")
	route.To(printDetails)
	ws.Route(route)

	cs := restful.NewContainer()
	cs.Add(ws)

	log.Println("======= Server is starting on port number 8081 =======")
	http.ListenAndServe(":8081", cs)
}

// port:= http//localhost/:8081/details
func printDetails(req *restful.Request, res *restful.Response) {
	type Details struct {
		Name   string `json:"name"`
		Dob    string `json:"dob"`
		Mobile int64  `json:"mobile"`
	}
	details := Details{}
	req.ReadEntity(&details)
	log.Println(details)
	res.WriteAsJson(details)
}
