package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strconv"
)

func main() {
	ws := &restful.WebService{}
	route := &restful.RouteBuilder{}
	route.Path("/even")
	route.Doc("testing the even or odd number api")
	route.Method("POST")
	route.To(EVENODD)
	route.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(route)
	container := restful.NewContainer()
	container.Add(ws)

	log.Println("Starting the server on 8082 port")
	http.ListenAndServe(":8082", container)
}

// :8082/even
func EVENODD(req *restful.Request, res *restful.Response) {
	type NUM struct {
		Num int `json:"num"`
	}
	num := NUM{}
	req.ReadEntity(&num)
	fmt.Println(num)
	Even := ""

	if num.Num%2 == 0 {
		Even = strconv.Itoa(num.Num) + " is a Even Number"
		//	res.WriteAsJson(Even)
		//fmt.Println("even number is ", Even)
	} else {
		Even = strconv.Itoa(num.Num) + " is a Odd Number"
		// res.WriteAsJson(Even)
		//fmt.Println("Odd number is ", Even)
	}
	res.WriteAsJson(Even)
}
