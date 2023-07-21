package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

func main() {

	ws := &restful.WebService{}
	//creating rout path using restful.Routebuilder
	helloRouteBuilder := &restful.RouteBuilder{}
	ws.Produces(restful.MIME_JSON)
	helloRouteBuilder.Path("/sum")                // host ke bad me aata hai es se aapka api heat hoga kyoki ye main path
	helloRouteBuilder.To(sum)                     // ye uske liye h jb api heat hogs user path ke name se to Hello function call karega esko ahandler v bolt hai
	helloRouteBuilder.Doc("testing sum api")      // ye kawal doc reader ka kaam krta hai ye kwal api ke liye use hota hai
	helloRouteBuilder.Produces(restful.MIME_JSON) // ye kaisa data server ko send karega
	helloRouteBuilder.Method("POST")              // kaisa method call hoga
	//registering api/rout
	ws.Route(helloRouteBuilder)
	cs := restful.NewContainer()
	cs.Add(ws)
	log.Println("************************starting http server on 8085 **********************************")
	http.ListenAndServe(":8085", cs)

}

func sum(req *restful.Request, res *restful.Response) {

	type No struct {
		Nos []int `json:"no"`
	}
	no := No{}
	req.ReadEntity(&no)
	fmt.Println(no)
	sum := 0.0

	for _, n := range no.Nos {
		sum = sum + n
	}
	/*Nos := make(map[string]interface{})
	req.ReadEntity(&Nos)
	nos := Nos["no"]
	fmt.Println(Nos)
	//fmt.Printf("%T\n", nos)
	//nosinterfaceArray := nos.([]interface{})
	//sum := 0
	//for _, noInterface := range nosinterfaceArray {
	//	fmt.Printf("%T\n", noInterface)
	//	intno, err := noInterface.(json.Number).Int64()
	//	if err != nil {
	//		fmt.Print(err)
	//	} else {
	//		sum = sum + int(intno)
	//	}
	//	fmt.Println(noInterface)
	}*/
	res.WriteAsJson(sum)

}
