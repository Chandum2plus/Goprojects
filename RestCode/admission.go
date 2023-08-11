package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

/*
Write a C program to determine eligibility for admission to a professional course based on the following criteria: Go to the editor
Eligibility Criteria : Marks in Maths >=65 and Marks in Phy >=55 and Marks in Chem>=50 and Total in all three subject >=190 or Total in Maths and Physics >=140 -------------------------------------
Input the marks obtained in Physics :65
Input the marks obtained in Chemistry :51
Input the marks obtained in Mathematics :72
Total marks of Maths, Physics and Chemistry : 188
Total marks of Maths and Physics : 137 The candidate is not eligible.
Expected Output :
The candidate is not eligible for admission.
*/
func main() {
	webservice := &restful.WebService{}
	RoutBuilder := &restful.RouteBuilder{}
	RoutBuilder.Doc("Testing the Student is Eligible for Admission or Not according to their Marks ")
	RoutBuilder.Path("admission")
	RoutBuilder.Method("POST")
	RoutBuilder.To(admission)
	RoutBuilder.Consumes(restful.MIME_JSON, restful.MIME_XML)
	RoutBuilder.Produces(restful.MIME_JSON, restful.MIME_XML)
	webservice.Route(RoutBuilder)
	container := restful.NewContainer()
	container.Add(webservice)
	log.Println("Server is Starting on port number 8989")
	port := ":8989"
	http.ListenAndServe(port, container)
}

// http://localhost:/8989/admission
func admission(req *restful.Request, res *restful.Response) {
	type MARKS struct {
		Phy  float64 `json:"phy"`
		Math float64 `json:"math"`
		Chem float64 `json:"chem"`
	}
	m := MARKS{}
	fmt.Println(m)
	req.ReadEntity(&m)
	if m.Math >= 65 && m.Phy >= 55 && m.Chem >= 50 {
		// This is for printing on Editor console
		fmt.Println("You are Eligible for Admission")
		fmt.Printf("Your Total Marks is - %f\n", m.Math+m.Phy+m.Chem)
		// this is for printing on postman
		fmt.Fprintln(res, "You are Eligible for Admission")
		fmt.Fprintf(res, "Your Marks is = %f\n", m.Chem+m.Phy+m.Math)

	} else if (m.Phy+m.Chem+m.Math) >= 190 || (m.Phy+m.Math) >= 140 {
		// This is for Printing on Editor Console
		fmt.Println("You are Eligible for Admission")
		fmt.Printf("Your Total Marks is - %f\n", m.Math+m.Phy+m.Chem)
		fmt.Printf("Physics and Math's Marks is = %f\n", m.Math+m.Phy)
		// This is for printing on Postman
		fmt.Fprintln(res, "You are Eligible for Admission")
		fmt.Fprintf(res, "Your Total Marks is = %f\n", m.Math+m.Phy+m.Chem)
		fmt.Fprintf(res, "Physic and Math's Marks is = %f\n", m.Math+m.Phy)
	} else {
		// this for Printing on Editor console
		fmt.Println("You are Not Eligible for Admission")
		fmt.Printf("Your Total Marks is - %f\n", m.Math+m.Phy+m.Chem)
		// This is for Printing on Postman
		fmt.Fprintln(res, "You are Not Eligible for Admission")
		fmt.Fprintf(res, "Your Total marks is = %f\n", m.Phy+m.Math+m.Chem)
	}
}
