package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	WebServices := &restful.WebService{}
	RouteBuilder := &restful.RouteBuilder{}
	RouteBuilder.Doc("Testing the Height")
	RouteBuilder.To(menHeight)
	RouteBuilder.Method("POST")
	RouteBuilder.Path("/height")
	RouteBuilder.Consumes(restful.MIME_JSON, restful.MIME_XML)
	RouteBuilder.Produces(restful.MIME_JSON, restful.MIME_XML)
	WebServices.Route(RouteBuilder)
	container := restful.NewContainer()
	container.Add(WebServices)
	port := ":9090"
	log.Println("Server is Starting on port number 9090")
	http.ListenAndServe(port, container)
}

/*
Write a C program to accept the height of a person in centimeters and categorize the person according to their height.
Test Data : 135
Expected Output :
The person is Dwarf.
*/
func menHeight(req *restful.Request, res *restful.Response) {
	type HEIGHT struct {
		Height float64 `json:"height"`
	}
	fmt.Println("Please Enter the Height in Centimeter -")
	h := HEIGHT{}
	fmt.Println(h)
	req.ReadEntity(&h)
	if h.Height < 150 {
		fmt.Fprintln(res, "Dwarf Height - ", h.Height)
	} else if h.Height >= 150 && h.Height < 165 {
		fmt.Fprintln(res, "Average Height - ", h.Height)
	} else if h.Height >= 165 && h.Height <= 195 {
		fmt.Fprintln(res, "Tall Height - ", h.Height)
	} else {
		fmt.Fprintln(res, "Abnormal Height - ", h.Height)
	}

}
