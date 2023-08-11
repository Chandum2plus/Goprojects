package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func strArray(req *restful.Request, res *restful.Response) {
	type StringArray struct {
		Str []string `json:"str"`
	}
	//	var strar = [...]string{}  // if you want to Display the given array then you can print like this
	s := StringArray{}
	req.ReadEntity(&s)
	for i := 0; i < len(s.Str); i++ {
		fmt.Printf("%s\n", s.Str[i])
	}
	res.WriteAsJson(&s.Str)
}
func main() {
	ws := &restful.WebService{}
	routebldr := &restful.RouteBuilder{}
	routebldr.Doc("Testing the String Array using the Restful Api")
	routebldr.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_JSON, restful.MIME_XML)
	routebldr.To(strArray)
	routebldr.Path("/stringArray")
	routebldr.Method("POST")
	ws.Route(routebldr)
	container := restful.NewContainer()
	container.Add(ws)
	log.Println(" ======= Starting the Server on port number 8081 ========")
	http.ListenAndServe(":8081", container)
}
