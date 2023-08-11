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
	RoutBuilder.To(greatest)
	RoutBuilder.Path("/greatest")
	RoutBuilder.Method("POST")

	webservice.Route(RoutBuilder)
	container := restful.NewContainer()
	container.Add(webservice)
	log.Println("Server is starting on port number - 7887")
	port := ":7887"
	http.ListenAndServe(port, container)
}
func greatest(req *restful.Request, res *restful.Response) {
	type GREATEST struct {
		Num1 float64 `json:"num1"`
		Num2 float64 `json:"num2"`
		Num3 float64 `json:"num3"`
	}
	g := GREATEST{}
	fmt.Println(g)
	req.ReadEntity(&g)
	fmt.Fprintln(res, "Find the Largest Value among the Three number")
	if g.Num1 > g.Num2 && g.Num1 > g.Num3 {
		fmt.Fprintf(res, " %f Number is Greatest ", g.Num1)
	} else if g.Num2 > g.Num1 && g.Num2 > g.Num3 {
		fmt.Fprintf(res, "%f Number is Greates ", g.Num2)
	} else {
		fmt.Fprintf(res, "%f Number is Greatest ", g.Num3)
	}
}
