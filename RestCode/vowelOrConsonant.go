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
	RoutBuilder.Doc("Testing the Character Vowel or Consonant ")
	RoutBuilder.Consumes(restful.MIME_JSON, restful.MIME_XML)
	RoutBuilder.Produces(restful.MIME_JSON, restful.MIME_XML)
	RoutBuilder.To(vowelOrConsonant)
	RoutBuilder.Path("/TestChar")
	RoutBuilder.Method("POST")

	webservice.Route(RoutBuilder)
	container := restful.NewContainer()
	container.Add(webservice)
	log.Println("Server is starting on port number - 7887")
	port := ":7887"
	http.ListenAndServe(port, container)
}
func vowelOrConsonant(req *restful.Request, res *restful.Response) {
	type TESTCHAR struct {
		Char rune `json:"char"`
	}
	c := TESTCHAR{}
	//m:=make(map[string]interface{})
	fmt.Println(c)
	req.ReadEntity(&c.Char)
	if c.Char == 'a' || c.Char == 'e' || c.Char == 'i' || c.Char == 'o' || c.Char == 'u' || c.Char == 'A' || c.Char == 'E' || c.Char == 'I' || c.Char == 'O' || c.Char == 'U' {
		fmt.Fprintf(res, "%c This is Vowel", c.Char)
	} else {
		fmt.Fprintf(res, "%c This is Consonant", c.Char)
	}
}
