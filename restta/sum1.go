package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	ws := &restful.WebService{}
	sumRouteBuilder := &restful.RouteBuilder{}
	sumRouteBuilder.Path("/sum")
	sumRouteBuilder.Doc("testing the sum api")
	sumRouteBuilder.Produces(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_JSON, restful.MIME_XML)
	sumRouteBuilder.To(sum1)
	sumRouteBuilder.Method("POST")
	ws.Route(sumRouteBuilder)
	container := restful.NewContainer()
	container.Add(ws)
	// starting the server on :8081
	log.Println("======== Starting the server on 8081 ==========")
	port := ":8081"
	http.ListenAndServe(port, container)
}

// http//:8081/sum
func sum1(req *restful.Request, res *restful.Response) {
	type SUM struct {
		Num1 int `json:"num1"`
		Num2 int `json:"num2"`
	}
	num := SUM{}
	req.ReadEntity(&num)
	fmt.Println(num)
	sum := 0
	sum = num.Num1 + num.Num2
	res.WriteAsJson(sum)

}
