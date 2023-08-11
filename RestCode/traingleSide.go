package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	webservice := &restful.WebService{}
	RoutBuilder := &restful.RouteBuilder{}
	RoutBuilder.Doc("Testing the Number is Greatest among the Three number ")
	RoutBuilder.Consumes(restful.MIME_JSON, restful.MIME_XML)
	RoutBuilder.Produces(restful.MIME_JSON, restful.MIME_XML)
	RoutBuilder.To(tRaingle)
	RoutBuilder.Path("/side")
	RoutBuilder.Method("POST")

	webservice.Route(RoutBuilder)
	container := restful.NewContainer()
	container.Add(webservice)
	log.Println("Server is starting on port number - 7887")
	port := ":7887"
	http.ListenAndServe(port, container)
}
func tRaingle(req *restful.Request, res *restful.Response) {
	type TRAINGLE struct {
		SideA float64 `json:"sideA"`
		SideB float64 `json:"sideB"`
		SideC float64 `json:"sideC"`
	}
	t := TRAINGLE{}
	fmt.Print(t)
	req.ReadEntity(&t)
	if t.SideA == t.SideB && t.SideA == t.SideC { // check whether all sides are equal
		fmt.Fprintln(res, "This is Equilateral Triangle")
	} else if t.SideA == t.SideB || t.SideA == t.SideC || t.SideB == t.SideC { // check whether two sides are equal
		fmt.Fprintln(res, "This is Isosceles Triangle")
	} else { // check whether no sides are equal
		fmt.Fprintln(res, "This is Scalene Triangle")
	}
}
