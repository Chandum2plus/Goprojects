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
	RoutBuilder.To(test)
	RoutBuilder.Path("/Test")
	RoutBuilder.Method("POST")

	webservice.Route(RoutBuilder)
	container := restful.NewContainer()
	container.Add(webservice)
	log.Println("Server is starting on port number - 7887")
	port := ":7887"
	http.ListenAndServe(port, container)
}

/*
16. Write a C program to check whether a character is an alphabet, digit or special character.
Test Data :
@
Expected Output :
This is a special character.
*/
func test(req *restful.Request, res *restful.Response) {
	type TESTALPHABET struct {
		Input interface{} `json:"input"`
	}
	t := TESTALPHABET{}
	fmt.Print(t.Input)
	req.ReadEntity(&t)
	if t.Input == 'a' && t.Input == 'z' || t.Input == 'A' && t.Input == 'Z' {
		fmt.Fprintf(res, " %c This is the Alphabet", t.Input)
	} else if t.Input == '0' || t.Input == '9' {
		fmt.Fprintf(res, "%c This is the Number ", t.Input)
	} else {
		fmt.Fprintf(res, "%c This is the Special Charecter", t.Input)
	}
}
