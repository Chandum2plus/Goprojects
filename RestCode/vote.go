package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

/*
Write a  program to read the age of a candidate and determine whether he is eligible to cast his/her own vote.
Test Data : 21
Expected Output :
Congratulation! You are eligible for casting your vote.
*/
func main() {
	ws := &restful.WebService{}
	route := &restful.RouteBuilder{}
	route.Doc("Testing the Age Eligible for vote or not")
	route.Path("/vote")
	route.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_XML, restful.MIME_JSON)
	route.To(vote)
	route.Method("POST")
	ws.Route(route)
	cs := restful.NewContainer()
	cs.Add(ws)
	log.Println("Server is starting on port number 8080")
	http.ListenAndServe(":8080", cs)
}
func vote(req *restful.Request, res *restful.Response) {
	type AGE struct {
		Age float64 `json:"age"`
	}
	a := AGE{}
	fmt.Println(a)
	req.ReadEntity(&a)
	if a.Age >= 18 && a.Age <= 70 {
		fmt.Fprintln(res, "You are Eligible for Vote - ", a.Age)
	} else if a.Age > 70 {
		fmt.Fprintln(res, "You can't Vote, because your age has crossed 70 - ", a.Age)
	} else if a.Age < 18 {
		fmt.Fprintln(res, "You are not Eligible for Vote - ", a.Age)
	}
}
