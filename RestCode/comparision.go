package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

// accept the two integer number and check whether they are eqaual or not
func main() {
	ws := &restful.WebService{}
	route := &restful.RouteBuilder{}
	route.Doc("Testing  the two number both are equal or not ")
	route.To(comparision)
	route.Consumes(restful.MIME_JSON, restful.MIME_XML)
	route.Produces(restful.MIME_JSON, restful.MIME_XML)
	route.Method("POST")
	route.Path("/comp")
	ws.Route(route)
	cs := restful.NewContainer()
	cs.Add(ws)
	log.Println("Server is starting on port number 8089")
	http.ListenAndServe(":8089", cs)
}
func comparision(req *restful.Request, res *restful.Response) {

	//var (
	//	num   int
	//	num2  int
	//	Equal string
	//)
	//req.ReadEntity(&num)
	//req.ReadEntity(&num2)
	//fmt.Print("numbers- ", num)
	//fmt.Print("number 2 ", num2)
	//if num == num2 {
	//	Equal = strconv.Itoa(num+num2) + "  both are equal"
	//	fmt.Println("both are equal")
	//} else {
	//	Equal = strconv.Itoa(num+num2) + "  not equal"
	//	fmt.Println("not equal")
	//}
	//res.WriteAsJson(&Equal)

	type Compa struct {
		Num  interface{} `json:"num"`
		Num2 interface{} `json:"num2"`
	}
	Res := ""
	n := Compa{}

	req.ReadEntity(&n)
	fmt.Print("Number - ", n)
	if n.Num == n.Num2 {
		//Res = strconv.Atoi(n.Num) + " Both are equal"
		res.WriteAsJson(&Res)
		fmt.Println("Both are Equal")
	} else {
		//Res = strconv.Itoa(n.Num2) + " Not Equal"
		fmt.Println("Not Equal")
		res.WriteAsJson(&Res)
	}

}
