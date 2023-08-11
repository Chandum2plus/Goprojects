package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

/*
	Write a C program to read the roll no, name and marks of three subjects and calculate the total, percentage and division.

Test Data :
Input the Roll Number of the student :784
Input the Name of the Student :James
Input the marks of Physics, Chemistry and Computer Application : 70 80 90
Expected Output :
Roll No : 784
Name of Student : James
Marks in Physics : 70
Marks in Chemistry : 80
Marks in Computer Application : 90
Total Marks = 240
Percentage = 80.00
Division = First
*/
func main() {
	webservice := &restful.WebService{}
	RoutBuilder := &restful.RouteBuilder{}
	RoutBuilder.Doc("Test the Roll and Name and marks of the five subject and calculate the total percentage and division")
	RoutBuilder.Consumes(restful.MIME_JSON, restful.MIME_XML)
	RoutBuilder.Produces(restful.MIME_JSON, restful.MIME_XML)
	RoutBuilder.To(percentage)
	RoutBuilder.Path("/percentage")
	RoutBuilder.Method("POST")

	webservice.Route(RoutBuilder)
	container := restful.NewContainer()
	container.Add(webservice)
	log.Println("Server is starting on port number - 6767")
	port := ":6767"
	http.ListenAndServe(port, container)
}
func percentage(req *restful.Request, res *restful.Response) {
	type PERCENTAGE struct {
		Name            string  `json:"name"`
		Roll            int     `json:"roll"`
		Fundamental     float64 `json:"fundamental"`
		DSA             float64 `json:"dsa"`
		Go              float64 `json:"go"`
		OperatingSystem float64 `json:"operating_system"`
		Github          float64 `json:"github"`
	}
	var Perc float64
	var Res float64
	p := PERCENTAGE{}
	fmt.Print(p)
	req.ReadEntity(&p)
	fmt.Fprintln(res, "PLEASE FILL THESE PARAMETERS -")
	fmt.Fprintln(res, "Name\tRoll\tFundamental\tDSA\tGo\tOperating_system\tGithub")

	fmt.Fprintln(res, "")

	if (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) >= 150 && (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) <= 224 {

		fmt.Fprintln(res, "------OUTPUT -------")
		fmt.Fprintln(res, "Name - ", p.Name)
		fmt.Fprintln(res, "Roll - ", p.Roll)

		fmt.Fprintln(res, "Division - Third")
		Res = p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA
		Perc = (p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA) / 500 * 100
		fmt.Fprintln(res, "Total Marks - ", Res)
		fmt.Fprintf(res, "Percentage = %f", Perc)

	} else if (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) >= 225 && (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) <= 299 {

		fmt.Fprintln(res, "------OUTPUT -------")
		fmt.Fprintln(res, "Name - ", p.Name)
		fmt.Fprintln(res, "Roll - ", p.Roll)

		fmt.Fprintln(res, "Division - Second")
		Res = p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA
		Perc = (p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA) / 500 * 100
		fmt.Fprintln(res, "Total Marks - ", Res)
		fmt.Fprintf(res, "Percentage = %f", Perc)

	} else if (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) >= 300 && (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) <= 400 {

		fmt.Fprintln(res, "------OUTPUT -------")
		fmt.Fprintln(res, "Name - ", p.Name)
		fmt.Fprintln(res, "Roll - ", p.Roll)

		fmt.Fprintln(res, "Division - First")
		Res = p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA
		Perc = (p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA) / 500 * 100
		fmt.Fprintln(res, "Total Marks - ", Res)
		fmt.Fprintf(res, "Percentage = %f", Perc)
	} else if (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) > 400 && (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) <= 450 {

		fmt.Fprintln(res, "------OUTPUT -------")
		fmt.Fprintln(res, "Name - ", p.Name)
		fmt.Fprintln(res, "Roll - ", p.Roll)

		fmt.Fprintln(res, "Division - District Top")
		Res = p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA
		Perc = (p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA) / 500 * 100
		fmt.Fprintln(res, "Total Marks - ", Res)
		fmt.Fprintf(res, "Percentage = %f", Perc)
	} else if (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) > 450 && (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) <= 500 {

		fmt.Fprintln(res, "------OUTPUT -------")
		fmt.Fprintln(res, "Name - ", p.Name)
		fmt.Fprintln(res, "Roll - ", p.Roll)

		fmt.Fprintln(res, "Division - State Top")
		Res = p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA
		Perc = (p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA) / 500 * 100
		fmt.Fprintln(res, "Total Marks - ", Res)
		fmt.Fprintf(res, "Percentage = %f", Perc)
	} else if (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) >= 100 && (p.Go+p.Github+p.Fundamental+p.OperatingSystem+p.DSA) <= 149 {
		fmt.Fprintln(res, "------OUTPUT -------")
		fmt.Fprintln(res, "Name - ", p.Name)
		fmt.Fprintln(res, "Roll - ", p.Roll)

		fmt.Fprintln(res, "Division - Failed")
		Res = p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA
		Perc = (p.Go + p.Github + p.Fundamental + p.OperatingSystem + p.DSA) / 500 * 100
		fmt.Fprintln(res, "Total Marks - ", Res)
		fmt.Fprintf(res, "Percentage = %f", Perc)
	}
	// TEST THE CODE LIKE THIS ON POSTMAN
	// {
	// "name":"Chandu_kumar",
	// "roll":1223,
	// "fundamental":68,
	// "dsa":98,
	// "go":78,
	// "operating_system":87,
	// "github":98

	// }
}
