package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		ws := &restful.WebService{}
		route := &restful.RouteBuilder{}
		route.Path("/alpha")
		route.Method("POST")
		route.Doc("Printing the alphabet using Rest api")
		route.Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
		route.To(alpabete)
		ws.Route(route)

		cs := restful.NewContainer()
		cs.Add(ws)
		log.Println("======= Server is starting on port number 8082 ======")
		http.ListenAndServe(":8082", cs)

	*/
	handlereq()

}

/*
	func alpabete(req *restful.Request, res *restful.Response) {
		res.WriteAsJson("Hello this is Chandu here !")
		res.WriteAsJson("hello india")
		log.Println("Hello world")
		var a string
		req.ReadEntity(&a)
		res.ResponseWriter.Write([]byte("hello this is india"))
		res.WriteAsJson(a)

}
*/
func handlereq() {
	http.HandleFunc("/hello", homepages)
	log.Println("====== server is starting on 1000 ======")
	log.Fatal(http.ListenAndServe(":1000", nil))
}
func homepages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to homepage !\n")
	fmt.Fprintln(w, " ########")
	fmt.Fprintln(w, " #         #")
	fmt.Fprintln(w, "#")
	fmt.Fprintln(w, "#")
	fmt.Fprintln(w, " #           #")
	fmt.Fprintln(w, "  #########")

	fmt.Println("EndPoint hit : homepages")
}
