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
	route.Doc("Testing the Age eligible for vote or not !")
	route.Path("/vote")
	route.To(vote)
	route.Method("POST")
	route.Consumes(restful.MIME_JSON, restful.MIME_XML)
	route.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(route)
	cs := restful.NewContainer()
	cs.Add(ws)
	log.Println("Starting the sever on port number - :8089")
	http.ListenAndServe(":8089", cs)

}

// http://localhost/:8089/vote
func vote(req *restful.Request, res *restful.Response) {
	//type VOTE struct {
	//	Age int `json:"age"`
	//}
	//a := VOTE{}
	a := 0
	req.ReadEntity(&a)
	fmt.Println("Your Age -", a)
	Res := ""

	if a >= 18 && a <= 70 {
		fmt.Println("You are Eligible for vote")             // this is written for displaying on console
		Res = strconv.Itoa(a) + " You are Eligible for Vote" //this is written for displaying on postman
	} else if a > 70 {
		fmt.Println("You can't Vote , because you has crossed the 70 years")             // this is written for displaying on console
		Res = strconv.Itoa(a) + " You can't vote, because you has crossed the age limit" // this is written for displaying on postman
	} else if a < 18 {
		Res = strconv.Itoa(a) + " You are not Eligible for vote" // this is written for displaying on postman

		fmt.Println("You are not Eligible for vote ") // this is written for displaying on console
	}
	//res.WriteAsJson(&a)
	res.WriteAsJson(Res) //this is responding the Result
}
