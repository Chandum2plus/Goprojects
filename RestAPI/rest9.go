// How to create simple API

package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"strings"
)

type Student struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Email   string `json:"email"`
	Class   string `json:"class"`
	Address string `json:"address"`
}
type studentResources struct {
	students map[string]Student
}

func getRoute(path, doc, method string, handler func(*restful.Request, *restful.Response)) *restful.RouteBuilder {
	route := &restful.RouteBuilder{}
	route.Path(path)
	route.Doc(doc)
	route.Method(method)
	route.Produces(restful.MIME_JSON)
	route.Method(strings.ToUpper(method))

	return route
}
func (s studentResources) createstuent(request *restful.Request, response *restful.Response) {
	ws := new(restful.WebService)
	ws.Doc("student")
	setRootPath(ws, ":/stdnt")
	ws.Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

}

func Hello(req *restful.Request, res *restful.Response) {
	res.WriteAsJson("Hello world")
}
func HelloName(req *restful.Request, res *restful.Response) {
	res.WriteAsJson("Hello this is Chandu kumar")
}
func main() {
	// Creating we service object/struct.
	ws := &restful.WebService{}

	// This is the first method Creating routePath
	// Creating routepath using restful.RouteBuilder object/struct

	helloRouteBuilder := getRoute("/hello", "testing hello API", "GET", Hello)

	// These are the second method of creating route path ->

	//helloRouteBuilder = &restful.RouteBuilder{} // this is the route builder
	//helloRouteBuilder.Path("/hello")				//this is the routePath
	//helloRouteBuilder.To(Hello)     					// this is the handler function
	//helloRouteBuilder.Doc("testing hello api")     // it use the only for documentation
	//helloRouteBuilder.Produces(restful.MIME_XML)   // returning data format json or xml
	//helloRouteBuilder.Method("GET")        // this is the method
	//param:= &restful.Parameter{}    // this is the routeParameter
	//helloRouteBuilder.Param(param) */
	//

	// Registering API/route
	ws.Route(helloRouteBuilder)

	// Registering new HelloName route/api
	// this is the first method of the registering route builder
	heloNameRouteBuilder := getRoute("/helloname", "getting name", "GET", HelloName)

	// this is the second method of the route

	//heloNameRouteBuilder:=&restful.RouteBuilder{}
	//heloNameRouteBuilder.Path("/hello/name")
	//heloNameRouteBuilder.Doc("Testing Hello name api")
	//heloNameRouteBuilder.Method("GET")
	//heloNameRouteBuilder.Produces(restful.MIME_XML)

	ws.Route(heloNameRouteBuilder)
	cs := restful.NewContainer()
	cs.Add(ws)
	address := "localhost:8070"
	log.Println("Starting http server on :- ", address)
	http.ListenAndServe(address, cs)

}
