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
	route.Doc("Testing the Positive or Negative Number -")
	route.To(positiveOrNegative)
	route.Method("POST")
	route.Consumes(restful.MIME_JSON, restful.MIME_XML)
	route.Produces(restful.MIME_JSON, restful.MIME_XML)
	route.Path("/positive")
	ws.Route(route)
	cs := restful.NewContainer()
	cs.Add(ws)
	log.Println("Server is starting on port 8081")
	http.ListenAndServe(":8081", cs)

}
func positiveOrNegative(req *restful.Request, res *restful.Response) {
	type POSTV struct {
		Num float64 `json:"num"`
	}
	n := POSTV{}
	fmt.Print("Number - ", n.Num)
	req.ReadEntity(&n)
	if n.Num > 0 {
		fmt.Fprintln(res, "Positive Number ", n.Num)
		fmt.Println("\npositive number ", n.Num)
	} else {
		fmt.Fprintln(res, "Negative Number ", n.Num)
		fmt.Println("\nnegative number ", n.Num)
	}
}
