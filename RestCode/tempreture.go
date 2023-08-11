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
	RoutBuilder.To(tempreture)
	RoutBuilder.Path("/temp")
	RoutBuilder.Method("POST")

	webservice.Route(RoutBuilder)
	container := restful.NewContainer()
	container.Add(webservice)
	log.Println("Server is starting on port number - 7887")
	port := ":7887"
	http.ListenAndServe(port, container)
}

/*
13. Write a  program to read temperature in centigrade and display a suitable message according to the temperature state below:
Temp < 0 then Freezing weather
Temp 0-10 then Very Cold weather
Temp 10-20 then Cold weather
Temp 20-30 then Normal in Temp
Temp 30-40 then Its Hot
Temp >=40 then Its Very Hot
Test Data :
42
Expected Output :
Its very hot.
*/
func tempreture(req *restful.Request, res *restful.Response) {
	type TEMPRETURE struct {
		Temp float64 `json:"temp"`
	}
	t := TEMPRETURE{}
	fmt.Print(t)
	req.ReadEntity(&t)
	if t.Temp <= 0 {
		fmt.Fprintln(res, "Freezing Weather ")
	} else if t.Temp > 0 && t.Temp <= 10 {
		fmt.Fprintln(res, "Very cold Weather ")
	} else if t.Temp > 10 && t.Temp <= 20 {
		fmt.Fprintln(res, "Cold Weather ")
	} else if t.Temp > 20 && t.Temp <= 30 {
		fmt.Fprintln(res, "Normal Temperature")
	} else if t.Temp > 30 && t.Temp >= 40 {
		fmt.Fprintln(res, "It is Hot")
	} else if t.Temp >= 45 {
		fmt.Fprintln(res, "IT is too Hot")
	}
}
