package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {
	ws := &restful.WebService{}
	routeBuilder := &restful.RouteBuilder{}
	routeBuilder.Path("/simple")
	routeBuilder.Doc("testing the simple interest api")
	routeBuilder.To(Simple)
	routeBuilder.Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	routeBuilder.Method("POST")

	ws.Route(routeBuilder)
	cs := restful.NewContainer()
	cs.Add(ws)

	log.Println("====== Server is starting on 8081 =========")
	log.Fatal(http.ListenAndServe(":8081", cs))
}
func Simple(req *restful.Request, res *restful.Response) {
	//type SIMPLE struct {
	//	Principle float64 `json:"principle"`
	//	Rate      float64 `json:"rate"`
	//	Time      float64 `json:"time"`
	//}
	//simpl := SIMPLE{}
	simpl := make(map[string]float64)
	req.ReadEntity(&simpl)
	fmt.Println(simpl)
	si := 0.0
	P := 0.0
	R := 0.0
	T := 0.0

	for v, er := range simpl {
		if v == "principle" {
			P = er
		} else if v == "rate" {
			R = er
		} else {
			T = er
		}
		fmt.Println(er)
	}

	si = (P * R * T) / 100
	res.WriteAsJson(si)
}
