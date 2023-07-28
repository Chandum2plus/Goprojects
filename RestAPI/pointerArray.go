package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

/*
we can declare an array that contains pointers as its elements. every element of this array is a pointer variable that can hold
address of any variable of appropriates type. the syntax of declaring an array of pointers is similar to that of declaring
arrays except that an asterisk of before the datatype name
*/
func main() {
	ws := &restful.WebService{}
	route := &restful.RouteBuilder{}
	route.Path("/pointerArray")
	route.Doc("Testing the pointer Array using the Rest Api")
	route.Consumes(restful.MIME_XML, restful.MIME_JSON).Produces(restful.MIME_JSON, restful.MIME_XML)
	route.Method("POST")
	route.To(pointerArray)
	ws.Route(route)

	cs := restful.NewContainer()
	cs.Add(ws)

	log.Println("================= Service is starting on port number 8080 ==============")
	http.ListenAndServe(":8080", cs)
}
func pointerArray(req *restful.Request, res *restful.Response) {
	// Here we are declaring the variable with value
	a := 5
	b := 10
	c := 15

	// Here we are declaring the array of pointer. which is empty array here is no element only defining the size
	pa := [3]*int{}
	pa[0] = &a
	pa[1] = &b
	pa[2] = &c
	req.ReadEntity(&pa)
	// Here use the for loop for iterate the array lengths
	fmt.Println("========== This is the Output =======")
	for i := 0; i < 3; i++ {
		fmt.Printf("pa [%d] = %p\t ", i, pa[i])
		fmt.Printf("*pa [%d] = %d \n", i, *pa[i])
	}
	res.WriteAsJson(&pa)
	/*
	   Here pa is declared as an array of pointers. Every element of this array is a pointer to an integer.
	   pa[i] gibes the value of the ith element of pa which is an address of any int variable and *pa[i] gives the value
	   of that int variable.the array of pointers can also contain addresses of elements of another array.
	*/
}
